package service

import (
	"crypto/ecdsa"
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log"
	"math/big"
	"os"
	"path/filepath"
	"strings"
	"time"

	utils "github.com/meta-node-blockchain/verify_server/utils"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"

	// "github.com/ethereum/go-ethereum/ethclient"
	// "github.com/ethereum/go-ethereum/rpc"
	"github.com/meta-node-blockchain/meta-node/cmd/client"
	"github.com/meta-node-blockchain/meta-node/pkg/logger"
	pb "github.com/meta-node-blockchain/meta-node/pkg/proto"

	"github.com/meta-node-blockchain/meta-node/pkg/transaction"

)

type LogData struct {
	UserWalletAddress      string `json:"userWalletAddress"`
	PhoneNumber            string `json:"phoneNumber"`
	EncryptedMessageBase64 string `json:"encryptedMessage"`
	EphemeralPublicKeyHex  string `json:"ephemeralPublicKey"`
	IVHex                  string `json:"iv"`
	Timestamp              string `json:"timestamp"`
}

type ValidateOTPResult struct {
	PublicKey string
	Wallet    common.Address
}

func CheckOTP(fromAddress common.Address ,client *client.Client,contractAddress string, contractABI string, RPC_HTTP_URL string, phoneNumber string, OTP string, botID string) {
	parsedABI, err := abi.JSON(strings.NewReader(contractABI))
	if err != nil {
		fmt.Printf("❌ Failed to parse ABI: %v\n", err)
		return
	}

	uintOtp, err := utils.StringToUint256(OTP)
	if err != nil {
		fmt.Printf("❌ Error converting OTP to uint256: %v\n", err)
		return
	}

	verifyOTPData, err := parsedABI.Pack("validateOTP", uintOtp, phoneNumber)
	if err != nil {
		fmt.Printf("❌ Failed to pack ABI: %v\n", err)
		return
	}

	toAddress := common.HexToAddress(contractAddress)

	// fromAddress := common.HexToAddress("0xa620249dc17f23887226506b3eb260f4802a7efc") // Replace with actual address
	relatedAddress := []common.Address{}
	maxGas := uint64(5_000_000_000)
	maxGasPrice := uint64(1_000_000_000)
	timeUse := uint64(0)

	// Step 15: Send transaction
		callData := transaction.NewCallData(verifyOTPData)

		bData, err := callData.Marshal()
		if err != nil {
			logger.Error(fmt.Sprintf("Marshal calldata for %s failed", "verifyOTP"), err)
		}


		receipt, err := client.SendTransactionWithDeviceKey(
			fromAddress,
			toAddress,
			big.NewInt(0),
			bData,
			relatedAddress,
			maxGas,
			maxGasPrice,
			timeUse,
		)
		if err !=nil {
			logger.Error("Error:",err)
		}
		if receipt.Status() == pb.RECEIPT_STATUS_RETURNED {
				log.Printf("✅ Step 15: Transaction sent successfully! ", )
			var decodedResult ValidateOTPResult
			err = parsedABI.UnpackIntoInterface(&decodedResult, "validateOTP", receipt.Return())
			if err != nil {
				fmt.Printf("❌ Failed to unpack result: %v\n", err)
			} else {
				fmt.Printf("✅ Decoded PublicKey: %s\n", decodedResult.PublicKey)
				fmt.Printf("✅ Decoded Wallet Address: %s\n", decodedResult.Wallet.Hex())
				if(botID == "email"){
					CallCompleteAuthentication(fromAddress,client, parsedABI, toAddress, phoneNumber, decodedResult.PublicKey, "72a147b91248b0396f34d2cebf5d9817336163f944d87bf40e66cddd06bddf0e", decodedResult.Wallet)

				}
			}

		}

}

func CallCompleteAuthentication(fromAddress common.Address ,client *client.Client, parsedABI abi.ABI, contractAddress common.Address, phoneNumber, publicKey, privateKeyHex string, userWalletAddress common.Address) {
	// Step 1: Create message
	message := fmt.Sprintf("Wallet address: %s is authorized", userWalletAddress.Hex())
	log.Printf("📝 Step 1: Created message: %s", message)

	// Step 2: Generate ephemeral key pair
	ephemeralPrivKey, err := crypto.GenerateKey()
	if err != nil {
		log.Printf("❌ Step 2: Error generating ephemeral key: %v\n", err)
		return
	}

	ephemeralPubKey := crypto.FromECDSAPub(&ephemeralPrivKey.PublicKey)
	ephemeralPubKeyHex := hex.EncodeToString(ephemeralPubKey)
	ephemeralPrivKeyBytes := crypto.FromECDSA(ephemeralPrivKey)

	log.Printf("🔑 Step 2: Generated ephemeral key pair")
	log.Printf("   - Ephemeral public key: %s", ephemeralPubKeyHex)

	// Step 3: Decode user's public key
	userPubKeyBytes, err := hex.DecodeString(strings.TrimPrefix(publicKey, "0x"))
	if err != nil {
		log.Printf("❌ Step 3: Error decoding user public key: %v\n", err)
		return
	}
	log.Printf("🔍 Step 3: Decoded user public key (length: %d bytes)", len(userPubKeyBytes))

	// Step 4: Calculate shared secret using ECDH
	log.Printf("🔐 Step 4: Calculating ECDH shared secret...")
	sharedSecretHex, err := utils.ECDHSharedSecretHex(ephemeralPrivKeyBytes, userPubKeyBytes)
	if err != nil {
		log.Printf("❌ Step 4: Error calculating ECDH shared secret: %v\n", err)
		return
	}

	sharedSecretBytes, err := hex.DecodeString(sharedSecretHex)
	if err != nil {
		log.Printf("❌ Step 4: Error decoding shared secret: %v\n", err)
		return
	}
	log.Printf("✅ Step 4: Shared secret calculated (length: %d bytes)", len(sharedSecretBytes))

	// Step 5: Generate random IV for AES-CBC
	iv := make([]byte, 16) // AES block size is 16 bytes
	if _, err := rand.Read(iv); err != nil {
		log.Printf("❌ Step 5: Error generating IV: %v\n", err)
		return
	}
	ivHex := hex.EncodeToString(iv)
	log.Printf("🎲 Step 5: Generated IV: %s", ivHex)

	// Step 6: Encrypt message using AES-CBC
	log.Printf("🔐 Step 6: Encrypting message with AES-CBC...")
	encryptedBytes, err := utils.EncryptAESCBC(sharedSecretBytes, []byte(message), iv)
	if err != nil {
		log.Printf("❌ Step 6: Error encrypting message: %v\n", err)
		return
	}
	encryptedMessageBase64 := base64.StdEncoding.EncodeToString(encryptedBytes)
	log.Printf("✅ Step 6: Message encrypted successfully (length: %d bytes)", len(encryptedBytes))

	// Step 7: Save to log
	err = saveEncryptedDataToLog(userWalletAddress.Hex(), phoneNumber, encryptedMessageBase64, ephemeralPubKeyHex, ivHex)
	if err != nil {
		log.Printf("⚠️ Warning: Could not write file to log folder: %v", err)
	}

	// Step 8: Decode ephemeral public key to bytes for transaction
	ephemeralPubKeyBytes, err := hex.DecodeString(ephemeralPubKeyHex)
	if err != nil {
		log.Printf("❌ Step 8: Error decoding ephemeral public key: %v\n", err)
		return
	}

	// Step 9: Load server private key
	privateKey, err := crypto.HexToECDSA(privateKeyHex)
	if err != nil {
		log.Printf("❌ Step 9: Error loading private key: %v", err)
		return
	}
	senderAddress := crypto.PubkeyToAddress(*privateKey.Public().(*ecdsa.PublicKey))
	log.Printf("📤 Step 9: Sender address: %s", senderAddress.Hex())

	// Step 10: Get nonce

	// Step 13: Pack data for completeAuthentication
	// Smart contract function signature:
	// completeAuthentication(string memory phoneNumber, bytes memory encryptedMessage, bytes memory ephemeralPublicKey, bytes memory iv)
	completeAuthData := packData(parsedABI, "completeAuthentication", phoneNumber, encryptedBytes, ephemeralPubKeyBytes)
	log.Printf("📦 Step 13: Data packed successfully")

	log.Printf("✍️ Step 14: Transaction signed successfully")
	// fromAddress := common.HexToAddress("0xa620249dc17f23887226506b3eb260f4802a7efc") // Replace with actual address
	relatedAddress := []common.Address{}
	maxGas := uint64(5_000_000_000)
	maxGasPrice := uint64(1_000_000_000)
	timeUse := uint64(0)

	// Step 15: Send transaction
		callData := transaction.NewCallData(completeAuthData)

		bData, err := callData.Marshal()
		if err != nil {
			logger.Error(fmt.Sprintf("Marshal calldata for %s failed", "migrateCode"), err)
		}


		receipt, err := client.SendTransactionWithDeviceKey(
			fromAddress,
			contractAddress,
			big.NewInt(0),
			bData,
			relatedAddress,
			maxGas,
			maxGasPrice,
			timeUse,
		)
		if receipt.Status() == pb.RECEIPT_STATUS_RETURNED {
				log.Printf("✅ Step 15: Transaction sent successfully! ", )

		}

}

func packData(parsedABI abi.ABI, method string, args ...interface{}) []byte {
	data, err := parsedABI.Pack(method, args...)
	if err != nil {
		log.Fatalf("Fatal: Failed to pack data for %s: %v", method, err)
	}
	return data
}

func saveEncryptedDataToLog(userWalletAddress, phoneNumber, encryptedMessage, ephemeralPubKey, iv string) error {
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
		EphemeralPublicKeyHex:  ephemeralPubKey,
		IVHex:                  iv,
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

	log.Printf("💾 Saved encrypted data to: %s", filePath)
	return nil
}

func truncateString(s string, maxLen int) string {
	if len(s) <= maxLen {
		return s
	}
	return s[:maxLen]
}
