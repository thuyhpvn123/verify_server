// SPDX-License-Identifier: GPL-3.0
pragma solidity ^0.8.19;

contract AuthOTP {

    // --- ENUMS & STRUCTS ---
    enum TypeMethod { WhatsApp, Telegram, Email }

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
        uint256 botId;
        uint256 timeRequest;
    }

    struct HashRecord {
        bytes32 dataHash;
        uint256 timestamp;
    }
    mapping(string => HashRecord) public publicKeyHashes;

    // --- STATE VARIABLES ---
    mapping(string => OTP) public OTPs;
    mapping(uint => DetailBot) public detailBots;
    uint public detailBotsCount;

    // --- LOGIC ĐÃ SỬA ĐỔI ---
    
    // Giữ lại: "Sổ vàng" xác nhận một địa chỉ ví đã được xác thực.
    mapping(address => bool) public authenticatedWallets;

    // Giữ lại: Mapping để "nhớ" tạm thời ví nào đang được xác thực bởi SĐT nào.
    // Logic khóa vĩnh viễn sẽ được loại bỏ.
    mapping(string => address) public phoneToWallet;

    // ĐÃ LOẠI BỎ: Không cần liên kết 2 chiều nữa.
    // mapping(address => string) public walletToPhone;

    // Giữ lại: Cooldown chống spam.
    mapping(address => uint256) public walletCooldown;

    mapping(address => bytes32) public authenticationHashes;


    // --- EVENTS ---
    event AuthenticationHashStored(address indexed wallet, bytes32 dataHash);
    event AuthenticationCompleted(address indexed wallet, string indexed phoneNumber);
    event AuthenticationRequested(address indexed wallet, uint256 otp, string chatbotPhone, TypeMethod typeMethod);

    // --- FUNCTIONS ---
    constructor() {}

    function addBot(string memory _phoneNumber, TypeMethod _typeMethod) public {
        detailBots[detailBotsCount] = DetailBot({
            phoneNumber: _phoneNumber,
            typeMethod: _typeMethod,
            busy: false,
            timeOccupied: 0,
            status: true
        });
        detailBotsCount++;
    }

    function updateBot(uint _botId, string memory _phoneNumber, TypeMethod _typeMethod, bool _status) public {
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

    function requestAuthentication(string memory _userPhoneNumber, address _walletAddress, string memory _publicKey, TypeMethod _typeMethod) public {
        require(OTPs[_userPhoneNumber].timeRequest + 60 < block.timestamp, "OTP for this phone was sent recently.");
        require(walletCooldown[_walletAddress] + 60 < block.timestamp, "OTP for this wallet was sent recently.");

        // ĐÃ LOẠI BỎ: Dòng code kiểm tra khóa vĩnh viễn đã được xóa bỏ.
        // Giờ đây một SĐT có thể bắt đầu phiên xác thực cho bất kỳ ví nào.

        uint availableBotId = findAvailableBot(_typeMethod);
        require(availableBotId < detailBotsCount, "No available bots.");

        uint256 otp = uint256(keccak256(abi.encodePacked(block.timestamp, _walletAddress))) % 1000000;

        detailBots[availableBotId].busy = true;
        detailBots[availableBotId].timeOccupied = block.timestamp + 300;

        OTPs[_userPhoneNumber] = OTP({
            OTP: otp,
            publicKey: _publicKey,
            verified: false,
            botId: availableBotId,
            timeRequest: block.timestamp
        });

        // THAY ĐỔI: Chỉ tạo liên kết tạm thời cho phiên làm việc này.
        // Lần sau nếu dùng SĐT này cho ví khác, nó sẽ ghi đè lên.
        phoneToWallet[_userPhoneNumber] = _walletAddress;

        walletCooldown[_walletAddress] = block.timestamp;

        emit AuthenticationRequested(_walletAddress, otp, detailBots[availableBotId].phoneNumber, detailBots[availableBotId].typeMethod);
    }

    function validateOTP(uint256 _otp, string memory userPhoneNumber) public returns (string memory publicKey, address wallet) {
        OTP storage request = OTPs[userPhoneNumber];
        require(request.OTP == _otp, "Invalid OTP");
        require(!request.verified, "Already verified");
        require(block.timestamp <= request.timeRequest + 300, "OTP expired");

        request.verified = true;
        detailBots[request.botId].busy = false;

        // Giữ nguyên: Vẫn trả về ví đã được liên kết tạm thời cho phiên này.
        return (request.publicKey, phoneToWallet[userPhoneNumber]);
    }

    function completeAuthentication(
        string memory _userPhoneNumber,
        bytes memory _encryptedMessage,   // Dữ liệu đã mã hóa AES
        bytes memory _encryptedSecretKey  // Khóa AES đã mã hóa RSA
    ) public {
        OTP storage request = OTPs[_userPhoneNumber];
        require(request.verified, "OTP not verified yet.");

        address userWallet = phoneToWallet[_userPhoneNumber];
        require(userWallet != address(0), "Phone not linked to any wallet for this session.");

        // 1. Đánh dấu địa chỉ ví này đã được xác thực
        authenticatedWallets[userWallet] = true;

        bytes32 dataHash = keccak256(abi.encodePacked(_encryptedMessage, _encryptedSecretKey));

        authenticationHashes[userWallet] = dataHash;

        emit AuthenticationCompleted(userWallet, _userPhoneNumber);

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