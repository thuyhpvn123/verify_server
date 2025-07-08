// SPDX-License-Identifier: GPL-3.0
pragma solidity ^0.8.19;

contract AuthOTP {

    enum TypeMethod { WhatsApp, Telegram }

    struct DetailBot {
        string phoneNumber;
        TypeMethod typeMethod;
        bool busy; 
        uint256 timeOccupied; //thời gian giới hạn
        bool status;
    }

    
    struct OTP {
        uint256 OTP;
        string  publicKey;
        bool verified;
        uint256 botId;
        uint256 timeRequest; // Thời gian yêu cầu OTP
    }

    struct HashRecord {
        bytes32 dataHash;
        uint256 timestamp;
    }

    mapping(string => HashRecord)  publicKeyHashes;
    mapping(string => address) public publicKeyUsers;

    //userPhoneNumber for OTP
    mapping(string  => OTP)  OTPs;
    mapping (uint => DetailBot) public detailBots;

    uint detailBotsCount;


    event AuthenticationRequested(address indexed user, uint256 otp, string chatbotPhone, TypeMethod typeMethod);
    
    constructor() {

    }
    
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

    
    function requestAuthentication(string memory _userPhoneNumber, string memory _publicKey , TypeMethod _typeMethod) public  returns (uint256)  {
        require(OTPs[_userPhoneNumber].timeRequest+300 < block.timestamp, "OTP just sent, try again later");

        // Tìm bot khả dụng theo loại TypeMethod
        uint availableBotId = findAvailableBot(_typeMethod);
        require(availableBotId < detailBotsCount, "No available bots, please wait");

        // Sinh OTP ngẫu nhiên
        uint256 otp = uint256(keccak256(abi.encodePacked(block.timestamp, msg.sender))) % 1000000;

        // Gán bot cho người dùng
        detailBots[availableBotId].busy = true;
        detailBots[availableBotId].timeOccupied = block.timestamp + 300;

        // Lưu OTP
        OTPs[_userPhoneNumber] = OTP({
            OTP: otp,
            publicKey: _publicKey,
            verified: false,
            botId: availableBotId,
            timeRequest: block.timestamp
        });

        // ✅ Lưu publicKey => address (user)
        publicKeyUsers[_publicKey] = msg.sender;

        emit AuthenticationRequested(msg.sender, otp, detailBots[   ].phoneNumber, detailBots[availableBotId].typeMethod);

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
        return detailBotsCount; // Trả về giá trị ngoài phạm vi nếu không có bot rảnh
    }

    
    function validateOTP(uint256 _otp, string memory userPhoneNumber) public returns (string memory) {
        OTP storage request = OTPs[userPhoneNumber];

        require(request.OTP == _otp, "Invalid OTP");
        require(!request.verified, "Already verified");
        require(block.timestamp <= request.timeRequest + 300, "OTP expired"); // OTP có hiệu lực trong 5 phút

        // Đánh dấu là đã xác thực
        request.verified = true;
        
        // Giải phóng bot để tiếp tục nhận yêu cầu mới
        detailBots[request.botId].busy = false;
        return request.publicKey;
    }

    
function completeAuthentication(string memory _data, string memory _publicKey) public {
    // Kiểm tra xem publicKey có tồn tại không
    require(publicKeyUsers[_publicKey] != address(0), "Public key not registered");

    // Kiểm tra xem publicKey đã có hash chưa
    require(publicKeyHashes[_publicKey].timestamp == 0, "Public key has no previous hash");

    // Tạo hash mới từ dữ liệu
    bytes32 hashedData = keccak256(abi.encodePacked(_data));

    // Lưu vào publicKeyHashes
    publicKeyHashes[_publicKey] = HashRecord({
        dataHash: hashedData,
        timestamp: block.timestamp
    });
}

    
    function verifyHash(string memory _publicKey, bytes32 _dataHash) public view returns (bool) {
        HashRecord memory record = publicKeyHashes[_publicKey];

        // Kiểm tra xem hash có tồn tại và còn trong thời gian hợp lệ không (3 ngày)
        return record.dataHash == _dataHash && (block.timestamp <= record.timestamp + 259200);
    }
function getOTPInfo(string memory userPhoneNumber) public view returns (
    uint256 otp, 
    string memory publicKey, 
    bool verified, 
    uint256 botId, 
    uint256 timeRequest
) {
    OTP storage request = OTPs[userPhoneNumber];
    return (
        request.OTP,
        request.publicKey,
        request.verified,
        request.botId,
        request.timeRequest
    );
}
function getPublicKeyHash(string memory _publicKey) public view returns (bytes32, uint256) {
    require(publicKeyHashes[_publicKey].timestamp != 0, "No record found");
    return (publicKeyHashes[_publicKey].dataHash, publicKeyHashes[_publicKey].timestamp);
}
}
