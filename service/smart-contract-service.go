package service

import (
	"WhatsappVerifyOTP/model"
	"WhatsappVerifyOTP/untils"
	"context"
	"crypto/ecdsa"
	"encoding/json"
	"fmt"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/rpc"
)

func CheckOTP(contractAddress string, contractABI string, INFURA_WS_URL string, phoneNumber string, OTP string, botID string) {
	// Kết nối với Ethereum qua WebSocket
	client, err := rpc.DialWebsocket(context.Background(), INFURA_WS_URL, "")
	if err != nil {
		fmt.Printf("❌ Failed to connect to Ethereum WebSocket: %v\n", err)
	}
	defer client.Close()

	// Parse ABI của contract
	parsedABI, err := abi.JSON(strings.NewReader(contractABI))
	if err != nil {
		fmt.Printf("❌ Failed to parse ABI: %v", err)
	}

	// Chuyển OTP thành uint256
	uintOtp, err := untils.StringToUint256(OTP)
	if err != nil {
		fmt.Printf("❌ Error converting OTP to uint256: %v", err)
	}

	// Pack ABI để tạo dữ liệu gọi hàm
	verifyOTPData, err := parsedABI.Pack("validateOTP", uintOtp, phoneNumber)
	if err != nil {
		fmt.Printf("❌ Failed to pack ABI: %v", err)
	}

	// Địa chỉ contract
	toAddress := common.HexToAddress(contractAddress)

	// Tạo struct CallMsg để gửi yêu cầu `eth_call`
	msgVerifyOTP := map[string]interface{}{
		"to":   toAddress.Hex(),               // Địa chỉ contract
		"data": hexutil.Encode(verifyOTPData), // Dữ liệu đã encode đúng chuẩn "0x"
		//
	}

	// Kết quả nhận được từ contract
	var result hexutil.Bytes
	err = client.CallContext(context.Background(), &result, "eth_call", msgVerifyOTP, "latest")
	if err != nil {
		fmt.Printf("❌ Failed to call contract: %v", err)
	}

	// Nếu result rỗng hoặc không hợp lệ, log lỗi ngay
	if len(result) == 0 {
		fmt.Printf("❌ Error: Contract returned empty result")
	}

	// Khai báo biến để nhận giá trị giải mã
	var publicKey string

	// Giải mã kết quả từ contract
	err = parsedABI.UnpackIntoInterface(&publicKey, "validateOTP", result)
	if err != nil {
		fmt.Printf("❌ Failed to unpack result: %v", err)
	} else {
		fmt.Printf("✅ Decoded PublicKey: %s\n", publicKey)
		CallCompleteAuthentication(client, parsedABI, toAddress, phoneNumber, publicKey, "0xa65f97f69e75e627c59f99bad2abd5096bfc5964dd8e66e28951aa9c984e7939", model.WhatsApp.Int())
	}
	// Kiểm tra kết quả giải mã

	// Gọi tiếp quá trình xác thực
}

func CallCompleteAuthentication(client *rpc.Client, parsedABI abi.ABI, contractAddress common.Address, phoneNumber, publicKey, privateKeyHex string, messageType int) {
	// 🛠️ 1. Tạo VerifyInfo JSON
	verifyData := model.VerifyInfo{
		PhoneNumber: phoneNumber,
		PublicKey:   publicKey,
	}
	verifyDataJson, err := json.Marshal(verifyData)
	if err != nil {
		fmt.Printf("Lỗi khi mã hóa JSON: %v", err)
	}

	// 🛠️ 2. Mã hóa dữ liệu
	encryptedData, _, err := untils.EncryptData(publicKey, verifyDataJson)

	// 🛠️ 3. Nạp Private Key
	privateKeyHex = strings.TrimPrefix(privateKeyHex, "0x") // Loại bỏ "0x" nếu có
	privateKey, err := crypto.HexToECDSA(privateKeyHex)
	if err != nil {
		fmt.Printf("Lỗi nạp private key: %v", err)
	}

	// 🛠️ 4. Lấy địa chỉ từ private key
	publicKeyECDSA := privateKey.Public().(*ecdsa.PublicKey)
	senderAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	// 🛠️ 5. Lấy nonce từ RPC (sửa lỗi `json: cannot unmarshal string into Go value of type uint64`)
	var nonceHex string
	err = client.CallContext(context.Background(), &nonceHex, "eth_getTransactionCount", senderAddress.Hex(), "pending")
	if err != nil {
		fmt.Printf("Lỗi lấy nonce: %v", err)
	}

	// Chuyển nonce từ hex về uint64
	nonce, err := hexutil.DecodeUint64(nonceHex)
	if err != nil {
		fmt.Printf("Lỗi chuyển đổi nonce từ hex: %v", err)
	}

	// 🛠️ 6. Đóng gói dữ liệu theo ABI
	completeAuthenticationData, err := parsedABI.Pack("completeAuthentication", encryptedData, publicKey)
	if err != nil {
		fmt.Printf("Lỗi đóng gói dữ liệu: %v", err)
	}

	// 🛠️ 7. Lấy gas price qua RPC
	var gasPriceHex string
	err = client.CallContext(context.Background(), &gasPriceHex, "eth_gasPrice")
	if err != nil {
		fmt.Printf("Lỗi lấy gas price: %v", err)
	}

	// Chuyển gasPrice từ hex về *big.Int
	gasPrice := new(big.Int)
	gasPrice.SetString(strings.TrimPrefix(gasPriceHex, "0x"), 16)

	// 🛠️ 8. Lấy Chain ID qua RPC
	var chainIDHex string
	err = client.CallContext(context.Background(), &chainIDHex, "eth_chainId")
	if err != nil {
		fmt.Printf("Lỗi lấy chain ID: %v", err)
	}

	// Chuyển chainID từ hex về *big.Int
	chainID := new(big.Int)
	chainID.SetString(strings.TrimPrefix(chainIDHex, "0x"), 16)

	// 🛠️ 9. Tạo giao dịch
	gasLimit := uint64(300000)
	tx := types.NewTransaction(nonce, contractAddress, big.NewInt(0), gasLimit, gasPrice, completeAuthenticationData)

	// 🛠️ 10. Ký giao dịch
	signer := types.NewEIP155Signer(chainID)
	signedTx, err := types.SignTx(tx, signer, privateKey)
	if err != nil {
		fmt.Printf("Lỗi ký giao dịch: %v", err)
	}

	// 🛠️ 11. Gửi giao dịch qua WebSocket
	rawTxBytes, err := signedTx.MarshalBinary()
	if err != nil {
		fmt.Printf("Lỗi mã hóa giao dịch: %v", err)
	}

	rawTxHex := hexutil.Encode(rawTxBytes) // Chuyển sang hex
	var txHash common.Hash
	err = client.CallContext(context.Background(), &txHash, "eth_sendRawTransaction", rawTxHex)
	if err != nil {
		fmt.Printf("Lỗi gửi giao dịch: %v", err)
	}

	fmt.Printf("✅ Giao dịch gửi thành công! TxHash: %s\n", txHash.Hex())
}
