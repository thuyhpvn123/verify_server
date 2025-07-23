package service

import (
	"context"
	"crypto/ecdsa"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"math/big"
	"os"
	"path/filepath"
	"strings"
	"time"
	"verify_server/utils" // Ensure your utils package is correctly named 'utils'

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	// No longer need "github.com/ethereum/go-ethereum/rpc"
)

type LogData struct {
	UserWalletAddress      string `json:"userWalletAddress"`
	PhoneNumber            string `json:"phoneNumber"`
	EncryptedMessageBase64 string `json:"encryptedMessage"`
	EncryptedAESKeyBase64  string `json:"encryptedSecretKey"`
	Timestamp              string `json:"timestamp"`
}
type ValidateOTPResult struct {
	PublicKey string
	Wallet    common.Address
}

func CheckOTP(contractAddress string, contractABI string, INFURA_WS_URL string, phoneNumber string, OTP string, botID string) {
	// 1. FIX: Use ethclient.Dial to get the correct client type
	client, err := ethclient.Dial(INFURA_WS_URL)
	if err != nil {
		fmt.Printf("❌ Failed to connect to Ethereum: %v\n", err)
		return
	}
	defer client.Close()

	// Parse ABI of the contract
	parsedABI, err := abi.JSON(strings.NewReader(contractABI))
	if err != nil {
		fmt.Printf("❌ Failed to parse ABI: %v\n", err)
		return
	}

	// Convert OTP to uint256
	uintOtp, err := utils.StringToUint256(OTP)
	if err != nil {
		fmt.Printf("❌ Error converting OTP to uint256: %v\n", err)
		return
	}

	// Pack ABI to create call data
	verifyOTPData, err := parsedABI.Pack("validateOTP", uintOtp, phoneNumber)
	if err != nil {
		fmt.Printf("❌ Failed to pack ABI: %v\n", err)
		return
	}

	toAddress := common.HexToAddress(contractAddress)

	// Create CallMsg struct for eth_call
	msgVerifyOTP := map[string]interface{}{
		"to":   toAddress.Hex(),
		"data": hexutil.Encode(verifyOTPData),
	}

	var result hexutil.Bytes
	// Use the rpc client from ethclient for the call
	err = client.Client().CallContext(context.Background(), &result, "eth_call", msgVerifyOTP, "latest")
	if err != nil {
		fmt.Printf("❌ Failed to call contract: %v\n", err)
		return
	}

	if len(result) == 0 {
		fmt.Println("❌ Error: Contract returned empty result")
		return
	}

	var decodedResult ValidateOTPResult
	err = parsedABI.UnpackIntoInterface(&decodedResult, "validateOTP", result)
	if err != nil {
		fmt.Printf("❌ Failed to unpack result: %v\n", err)
	} else {
		fmt.Printf("✅ Decoded PublicKey: %s\n", decodedResult.PublicKey)
		fmt.Printf("✅ Decoded Wallet Address: %s\n", decodedResult.Wallet.Hex())

		// 2. FIX: Removed the extra 'model.WhatsApp.Int()' argument from the call
		CallCompleteAuthentication(client, parsedABI, toAddress, phoneNumber, decodedResult.PublicKey, "a65f97f69e75e627c59f99bad2abd5096bfc5964dd8e66e28951aa9c984e7939", decodedResult.Wallet)
	}
}

// Function signature is correct and matches the call now
func CallCompleteAuthentication(client *ethclient.Client, parsedABI abi.ABI, contractAddress common.Address, phoneNumber, publicKey, privateKeyHex string, userWalletAddress common.Address) {
	// ... (rest of the function remains the same)
	message := fmt.Sprintf("Wallet address: %s is authorized", userWalletAddress.Hex())

	encryptedMessageBase64, encryptedAESKeyBase64, err := utils.EncryptData(publicKey, []byte(message))
	if err != nil {
		log.Printf("❌ Step 2: Error encrypting data: %v\n", err)
		return
	}

	encryptedMessageBytes, err := base64.StdEncoding.DecodeString(encryptedMessageBase64)
	if err != nil {
		log.Printf("❌ Step 3: Error decoding encrypted message from base64: %v\n", err)
		return
	}
	encryptedAESKeyBytes, err := base64.StdEncoding.DecodeString(encryptedAESKeyBase64)
	if err != nil {
		log.Printf("❌ Step 3: Error decoding encrypted AES key from base64: %v\n", err)
		return
	}

	err = saveEncryptedDataToLog(userWalletAddress.Hex(), phoneNumber, encryptedMessageBase64, encryptedAESKeyBase64)
	if err != nil {
		log.Printf("Can not write file into folder log")
	}

	privateKey, err := crypto.HexToECDSA(privateKeyHex)
	if err != nil {
		log.Printf("❌ Step 4: Error loading private key: %v", err)
		return
	}
	senderAddress := crypto.PubkeyToAddress(*privateKey.Public().(*ecdsa.PublicKey))

	nonce, err := client.PendingNonceAt(context.Background(), senderAddress)
	if err != nil {
		log.Printf("❌ Step 5: Error getting nonce: %v", err)
		return
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Printf("❌ Step 6: Error getting gas price: %v", err)
		return
	}

	chainID, err := client.ChainID(context.Background())
	if err != nil {
		log.Printf("❌ Step 7: Error getting chain ID: %v", err)
		return
	}

	completeAuthData := packData(parsedABI, "completeAuthentication", phoneNumber, encryptedMessageBytes, encryptedAESKeyBytes)

	gasLimit := uint64(500000)
	tx := types.NewTransaction(nonce, contractAddress, big.NewInt(0), gasLimit, gasPrice, completeAuthData)

	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		log.Printf("❌ Step 10: Error signing transaction: %v", err)
		return
	}

	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		log.Printf("❌ Step 11: Error sending transaction: %v", err)
		return
	}

	log.Printf("✅ Step 11: Transaction sent successfully! TxHash: %s", signedTx.Hash().Hex())
}

func packData(parsedABI abi.ABI, method string, args ...interface{}) []byte {
	data, err := parsedABI.Pack(method, args...)
	if err != nil {
		log.Fatalf("Fatal: Failed to pack data for %s: %v", method, err)
	}
	return data
}

func saveEncryptedDataToLog(userWalletAddress, phoneNumber, encryptedMessage, encryptedKey string) error {
	logDir := "log"
	if _, err := os.Stat(logDir); os.IsNotExist(err) {
		err = os.Mkdir(logDir, 0755)
		if err != nil {
			return fmt.Errorf("failed to create log directory: %w", err)
		}
	}

	data := LogData{
		UserWalletAddress:      userWalletAddress,
		PhoneNumber:            phoneNumber,
		EncryptedMessageBase64: encryptedMessage,
		EncryptedAESKeyBase64:  encryptedKey,
		Timestamp:              time.Now().Format(time.RFC3339),
	}

	fileContent, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal log data to JSON: %w", err)
	}

	fileName := fmt.Sprintf("%s_%s.json", userWalletAddress, time.Now().Format("20060102150405"))
	filePath := filepath.Join(logDir, fileName)

	err = os.WriteFile(filePath, fileContent, 0644)
	if err != nil {
		return fmt.Errorf("failed to write to log file: %w", err)
	}

	return nil
}
