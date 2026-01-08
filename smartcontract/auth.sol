// SPDX-License-Identifier: GPL-3.0
pragma solidity ^0.8.19;
import {Strings} from "@openzeppelin/contracts/utils/Strings.sol";
contract AuthOTP {

    // --- ENUMS & STRUCTS ---
    enum TypeMethod { WhatsApp, Telegram, Email }
    struct VerificationState {
        bool emailVerified;
        bool phoneVerified;
    }
    struct DetailBot {
        string phoneNumber;
        TypeMethod typeMethod;
        bool busy;
        uint256 timeOccupied;
        bool status;
    }

    struct OTP {
        uint256 OTP;
        string  publicKey;
        bool verified;
        uint256 botId; // Sẽ không được sử dụng cho phương thức Email
        uint256 timeRequest;
        TypeMethod method;
    }

    struct HashRecord {
        bytes32 dataHash;
        uint256 timestamp;
    }
    mapping(string => HashRecord) public publicKeyHashes;

    // --- STATE VARIABLES ---
    
    // MỚI: Thêm một biến để lưu địa chỉ email của mail server
    string public domain; 
    address public owner;

    mapping(string => OTP) public OTPs;
    mapping(uint => DetailBot) public detailBots;
    uint public detailBotsCount;

    mapping(address => bool) public authenticatedWallets;
    mapping(string => address) public identifierToWallet; // Đổi tên để dễ hiểu hơn
    mapping(address => uint256) public walletCooldown;
    mapping(address => bytes32) public authenticationHashes;
    mapping(address => string) public primaryEmails; // ví => email chính (PII)
    mapping(address => string) public subEmails; // ví => sub_email = "0x...@domain"
    // Liên kết định danh ↔ ví (tách riêng)
    mapping(string => address) public emailIdToWallet; // email -> wallet
    mapping(string => address) public phoneIdToWallet; // phone/WA/TG identifier -> wallet
    mapping(address => VerificationState) public verificationStates;
    mapping(address => bool) public isAdmin;
    // --- EVENTS ---
    event AuthenticationHashStored(address indexed wallet, bytes32 dataHash);
    event AuthenticationCompleted(address indexed wallet, string indexed identifier);
    
    // Event cho WhatsApp/Telegram
    event BotAuthenticationRequested(address indexed wallet, uint256 otp, string chatbotPhone, TypeMethod typeMethod);
    
    // Event MỚI dành riêng cho Email
    event EmailAuthenticationRequested(address indexed wallet, uint256 otp, string userEmail, string targetMailServerEmail);
    event EmailSubCreated(address indexed wallet, string primaryEmail, string subEmail);
    event EmailVerified(address indexed wallet, string fromEmail);
    event StepVerified(address indexed wallet, string indexed identifier);
    // --- MODIFIERS ---
    modifier onlyOwner() {
        require(msg.sender == owner, "Only owner can call this function");
        _;
    }
    modifier onlyAdmin() {
        require(isAdmin[msg.sender], "Only admin can call this function");
        _;
    }

    // --- FUNCTIONS ---
    constructor() {
        owner = msg.sender;
    }

    // MỚI: Hàm để chủ sở hữu contract thiết lập . vd : metanode.co
    function setDomainEmail(string memory _domain) public onlyOwner {
        domain = _domain;
    }
    function setAdmin(address _admin, bool _agreed) external onlyOwner {
        isAdmin[_admin] = _agreed;
    }
    function addBot(string memory _phoneNumber, TypeMethod _typeMethod) public onlyOwner {
        detailBots[detailBotsCount] = DetailBot({
            phoneNumber: _phoneNumber,
            typeMethod: _typeMethod,
            busy: false,
            timeOccupied: 0,
            status: true
        });
        detailBotsCount++;
    }

    function updateBot(uint _botId, string memory _phoneNumber, TypeMethod _typeMethod, bool _status) public onlyOwner {
        require(_botId < detailBotsCount, "Invalid bot ID");
        detailBots[_botId].phoneNumber = _phoneNumber;
        detailBots[_botId].typeMethod = _typeMethod;
        detailBots[_botId].status=_status;
    }

    function findAvailableBot(TypeMethod _typeMethod) internal view returns (uint) {
        for (uint i = 0; i < detailBotsCount; i++) {
            if(detailBots[i].status==false)
                continue;
            if (detailBots[i].typeMethod == _typeMethod && 
                (!detailBots[i].busy || block.timestamp > detailBots[i].timeOccupied)) {
                return i;
            }
        }
        return detailBotsCount;
    }
    
    // ĐÃ CẬP NHẬT: Hàm requestAuthentication được điều chỉnh logic
    function requestAuthentication(string memory _identifier, address _walletAddress, string memory _publicKey, TypeMethod _typeMethod) public {

        require(!authenticatedWallets[_walletAddress], "Wallet is already authenticated.");
        require(OTPs[_identifier].timeRequest + 60 < block.timestamp, "OTP for this identifier was sent recently.");
        require(walletCooldown[_walletAddress] + 60 < block.timestamp, "OTP for this wallet was sent recently.");

        uint256 otp = (uint256(keccak256(abi.encodePacked(block.timestamp, _walletAddress))) % 900000) + 100000;
        
        // --- PHÂN LUỒNG LOGIC ---
        if (_typeMethod == TypeMethod.Email) {
            // ---- XỬ LÝ CHO EMAIL ----
            require(bytes(domain).length > 0, "Mail server email not set.");

            // Lưu OTP request mà không cần thông tin bot
            OTPs[_identifier] = OTP({
                OTP: otp,
                publicKey: _publicKey,
                verified: false,
                botId: 0, // Không áp dụng cho email
                timeRequest: block.timestamp,
                method: _typeMethod
            });
            
            string memory targetEmail = registerPrimaryEmail(_identifier);
            // Phát sự kiện dành riêng cho Email
            emit EmailAuthenticationRequested(_walletAddress, otp, _identifier, targetEmail);

        } else {
            // ---- XỬ LÝ CHO WHATSAPP/TELEGRAM (Logic cũ) ----
            uint availableBotId = findAvailableBot(_typeMethod);
            require(availableBotId < detailBotsCount, "No available bots.");

            detailBots[availableBotId].busy = true;
            detailBots[availableBotId].timeOccupied = block.timestamp + 300;

            OTPs[_identifier] = OTP({
                OTP: otp,
                publicKey: _publicKey,
                verified: false,
                botId: availableBotId,
                timeRequest: block.timestamp,
                method: _typeMethod
            });
            address linked = phoneIdToWallet[_identifier];
            require(linked == address(0) || linked == _walletAddress, "identifier already linked");
            phoneIdToWallet[_identifier] = _walletAddress;
            // Phát sự kiện cho Bot
            emit BotAuthenticationRequested(_walletAddress, otp, detailBots[availableBotId].phoneNumber, detailBots[availableBotId].typeMethod);
        }

        // --- LOGIC CHUNG ---
        identifierToWallet[_identifier] = _walletAddress;
        walletCooldown[_walletAddress] = block.timestamp;
    }
    /**
     * @notice FE gọi: gửi email chính -> contract tạo sub_email = [wallet]@emailDomain và lưu on-chain.
     * @param primaryEmail Email chính của user (PII, lưu on-chain theo yêu cầu).
     */
    function registerPrimaryEmail(string memory primaryEmail) internal returns(string memory){
        require(bytes(domain).length > 0, "emailDomain not set");

        primaryEmails[msg.sender] = primaryEmail;

        // ánh xạ email -> wallet (để có thể dùng email làm identifier nếu cần)
        address linked = emailIdToWallet[primaryEmail];
        require(linked == address(0) || linked == msg.sender, "email already linked");
        emailIdToWallet[primaryEmail] = msg.sender;

        // sub_email theo yêu cầu: [địa_chỉ_ví]@domain
        string memory subEmail = string.concat(Strings.toHexString(uint160(msg.sender), 20), "@", domain);
        subEmails[msg.sender] = subEmail;

        emit EmailSubCreated(msg.sender, primaryEmail, subEmail);
        return subEmail;
    }

    function getSubEmail(address wallet) external view returns (string memory) {
        return subEmails[wallet];
    }
    // Đổi tên tham số để dễ hiểu hơn, logic không đổi
    // function validateOTP(uint256 _otp, string memory _identifier) public returns (string memory publicKey, address wallet) onlyAdmin {
    //     OTP storage request = OTPs[_identifier];
    //     require(request.OTP == _otp, "Invalid OTP");
    //     require(!request.verified, "Already verified");
    //     require(block.timestamp <= request.timeRequest + 300, "OTP expired");

        
        
    //     // Chỉ giải phóng bot nếu phương thức không phải là Email
    //     if (detailBots[request.botId].typeMethod != TypeMethod.Email) {
    //         detailBots[request.botId].busy = false;
            
    //     }
    //     wallet = phoneIdToWallet[_identifier];
    //     //phone
    //     if(wallet != address(0) && !verificationStates[wallet]){
    //         verificationStates[wallet].phoneVerified = true;
    //         emit StepVerified(wallet, _identifier);

    //     }
    
    //     if(wallet == address(0)){ //mail
    //         verificationStates[wallet].emailVerified = true;
    //         emit EmailVerified(wallet, _identifier);

    //     }
    //     return (request.publicKey, identifierToWallet[_identifier]);
    // }
    function validateOTP(
        uint256 _otp,
        string memory _identifier
    ) public onlyAdmin returns (string memory publicKey, address wallet) {

        OTP storage request = OTPs[_identifier];

        require(request.OTP == _otp, "Invalid OTP");
        require(!request.verified || authenticationHashes[userWallet] != bytes32(0), "Already verified");
        require(block.timestamp <= request.timeRequest + 300, "OTP expired");

        wallet = identifierToWallet[_identifier];
        require(wallet != address(0), "Wallet not linked");

        if (request.method == TypeMethod.Email) {
            require(verificationStates[userWallet].phoneVerified,"need to verify phone in advance");
            // ✅ EMAIL
            verificationStates[wallet].emailVerified = true;
            emit EmailVerified(wallet, _identifier);
        } else {
            // ✅ PHONE (WA / TG)
            detailBots[request.botId].busy = false;
            verificationStates[wallet].phoneVerified = true;
            emit StepVerified(wallet, _identifier);
        }
        if(verificationStates[wallet].phoneVerified && verificationStates[wallet].emailVerified){
            request.verified = true;
        }
        return (request.publicKey, wallet);
    }


    // Đổi tên tham số để dễ hiểu hơn, logic không đổi
    function completeAuthentication(
        string memory _identifier, //cho email
        bytes memory _encryptedMessage,
        bytes memory _encryptedSecretKey
    ) public {
        // OTP storage request = OTPs[_identifier];
        // require(request.verified, "Already verified");

        address userWallet = identifierToWallet[_identifier];
        require(userWallet != address(0), "Identifier not linked to any wallet for this session.");
        require(verificationStates[userWallet].phoneVerified && verificationStates[userWallet].emailVerified, "need both phone and email verified");
        authenticatedWallets[userWallet] = true;
        bytes32 dataHash = keccak256(abi.encodePacked(_encryptedMessage, _encryptedSecretKey));
        authenticationHashes[userWallet] = dataHash;

        emit AuthenticationCompleted(userWallet, _identifier);
        emit AuthenticationHashStored(userWallet, dataHash);
    }
    
    function verifyAuthenticationHash(
        address _userWallet,
        bytes memory _encryptedMessage,
        bytes memory _encryptedSecretKey
    ) public view returns (bool) {
        // Lấy hash đã được lưu trữ trên blockchain cho ví này
        bytes32 storedHash = authenticationHashes[_userWallet];

        // Nếu không có hash nào được lưu, chắc chắn là không hợp lệ
        if (storedHash == bytes32(0)) {
            return false;
        }

        // Tính toán lại hash từ dữ liệu được cung cấp
        bytes32 providedHash = keccak256(abi.encodePacked(_encryptedMessage, _encryptedSecretKey));

        // So sánh hai hash
        return storedHash == providedHash;
    }
}