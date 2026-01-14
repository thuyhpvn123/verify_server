package service

import (
	"context"
	"crypto/ecdsa"
	"crypto/rand"
	"crypto/tls"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log"
	"math/big"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	
	"github.com/meta-node-blockchain/verify_server/contracts" // ← Your abigen package
	"github.com/meta-node-blockchain/verify_server/utils"
)

// ============================================================================
// 🔧 TYPES
// ============================================================================

type ValidateOTPResult struct {
	PublicKey string
	Wallet    common.Address
}

type LogData struct {
	UserWalletAddress      string `json:"userWalletAddress"`
	PhoneNumber            string `json:"phoneNumber"`
	EncryptedMessageBase64 string `json:"encryptedMessage"`
	EphemeralPublicKeyHex  string `json:"ephemeralPublicKey"`
	IVHex                  string `json:"iv"`
	Timestamp              string `json:"timestamp"`
}

// ============================================================================
// 📞 CHECK OTP - Using abigen
// ============================================================================

func CheckOTP(
	ctx context.Context,
	privateKey *ecdsa.PrivateKey,
	contractAddress common.Address,
	rpcURL string,
	phoneNumber string,
	otp string,
	botID string,
) (*ValidateOTPResult, error) {
	
	// ============================================================
	// ✅ Create immutable copies
	// ============================================================
	phoneNumberCopy := strings.Clone(phoneNumber)
	otpCopy := strings.Clone(otp)
	botIDCopy := strings.Clone(botID)
	
	if botIDCopy == "" {
		return nil, fmt.Errorf("botID is empty")
	}
	
	log.Printf("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	log.Printf("🔍 CheckOTP START")
	log.Printf("   Phone: '%s'", phoneNumberCopy)
	log.Printf("   OTP: '%s'", otpCopy)
	log.Printf("   BotID: '%s'", botIDCopy)
	log.Printf("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")

	// ============================================================
	// 🌐 Connect to blockchain
	// ============================================================
	httpClient := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
		Timeout: 30 * time.Second,
	}
	
	rpcClient, err := rpc.DialHTTPWithClient(rpcURL, httpClient)
	if err != nil {
		return nil, fmt.Errorf("failed to create RPC client: %w", err)
	}
	
	client := ethclient.NewClient(rpcClient)
	defer client.Close()

	// ============================================================
	// 🔑 Setup transactor
	// ============================================================
	publicKey := privateKey.Public().(*ecdsa.PublicKey)
	fromAddress := crypto.PubkeyToAddress(*publicKey)
	
	chainID := big.NewInt(991) // Your chain ID
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		return nil, fmt.Errorf("failed to create transactor: %w", err)
	}

	auth.From = fromAddress
	auth.Value = big.NewInt(0)
	auth.GasLimit = uint64(5_000_000)
	auth.GasPrice = big.NewInt(1_000_000_000) // 1 gwei

	// ============================================================
	// 📜 Initialize contract
	// ============================================================
	instance, err := contract.NewContract(contractAddress, client)
	if err != nil {
		return nil, fmt.Errorf("failed to init contract: %w", err)
	}

	// ============================================================
	// 🔢 Convert OTP to uint256
	// ============================================================
	otpBigInt, ok := new(big.Int).SetString(otpCopy, 10)
	if !ok {
		return nil, fmt.Errorf("invalid OTP format")
	}

	// ============================================================
	// 📤 Send validateOTP transaction
	// ============================================================
	log.Printf("🔄 [%s] Sending validateOTP (botID=%s)...", phoneNumberCopy, botIDCopy)
	startTime := time.Now()
	
	tx, err := instance.ValidateOTP(auth, otpBigInt, phoneNumberCopy)
	if err != nil {
		return nil, fmt.Errorf("transaction failed: %w", err)
	}

	log.Printf("📝 [%s] Tx hash: %s", phoneNumberCopy, tx.Hash().Hex())

	// ============================================================
	// ⏳ Wait for transaction to be mined
	// ============================================================
	ctxTimeout, cancel := context.WithTimeout(context.Background(), 3*time.Minute)
	defer cancel()

	receipt, err := bind.WaitMined(ctxTimeout, client, tx)
	if err != nil {
		return nil, fmt.Errorf("wait mined failed: %w", err)
	}

	elapsed := time.Since(startTime)
	log.Printf("⏱️  [%s] Transaction completed in %v", phoneNumberCopy, elapsed)

	// ============================================================
	// ✅ Check receipt status
	// ============================================================
	if receipt.Status != 1 {
		return nil, fmt.Errorf("transaction failed with status: %d", receipt.Status)
	}

	log.Printf("✅ [%s] Transaction SUCCESS", phoneNumberCopy)
	log.Printf("📦 [%s] Block: %d, Gas used: %d", phoneNumberCopy, receipt.BlockNumber.Uint64(), receipt.GasUsed)

	// ============================================================
	// 📖 Parse return value from logs (if available)
	// ============================================================
	var result ValidateOTPResult
	
	// Parse OTPValidated event
	// for _, vLog := range receipt.Logs {
	// 	// Check if this is OTPValidated event
	// 	// event, err := instance.ParseOTPValidated(*vLog)
	// 	// if err != nil {
	// 	// 	continue // Not the event we're looking for
	// 	// }
		
	// 	result.PublicKey = event.PublicKey
	// 	result.Wallet = event.Wallet
		
	// 	log.Printf("✅ [%s] PublicKey: %s", phoneNumberCopy, result.PublicKey)
	// 	log.Printf("✅ [%s] Wallet: %s", phoneNumberCopy, result.Wallet.Hex())
	// 	break
	// }
	found := false

	for _, vLog := range receipt.Logs {

		// 1️⃣ PHONE (Telegram / WhatsApp)
		if ev, err := instance.ParseStepVerified(*vLog); err == nil {
			result.Wallet = ev.Wallet
			callOpts := &bind.CallOpts{
				From:    fromAddress,
				Context: context.Background(),
			}

			kq, err := instance.OTPs(callOpts,phoneNumberCopy)
			if err != nil {
				log.Printf("error on get OTPs", err)
				return &result, err
			}
			
			result.PublicKey = kq.PublicKey
			found = true

			log.Printf("✅ [%s] StepVerified", phoneNumberCopy)
			log.Printf("   Wallet: %s", ev.Wallet.Hex())
			break
		}

		// 2️⃣ EMAIL
		if ev, err := instance.ParseEmailVerified(*vLog); err == nil {
			result.Wallet = ev.Wallet
			callOpts := &bind.CallOpts{
				From:    fromAddress,
				Context: context.Background(),
			}
			kq, err := instance.OTPs(callOpts,phoneNumberCopy)
			if err != nil {
				log.Println("error on get OTPs", err)
				return &result, err
			}
			
			result.PublicKey = kq.PublicKey
			found = true

			log.Printf("✅ [%s] EmailVerified", phoneNumberCopy)
			log.Printf("   Wallet: %s", ev.Wallet.Hex())
			break
		}
	}

	if !found {
		return nil, fmt.Errorf("no StepVerified / EmailVerified event found")
	}
	if result.Wallet == (common.Address{}) {
		return nil, fmt.Errorf("failed to parse OTPValidated event")
	}

	// ============================================================
	// 🔐 Branch logic based on botID
	// ============================================================
	log.Printf("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	log.Printf("🔍 BRANCHING DECISION")
	log.Printf("   BotID: '%s'", botIDCopy)
	log.Printf("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")

	switch botIDCopy {
	case "telegram":
		log.Printf("ℹ️ [%s] Telegram flow - skip completeAuth", phoneNumberCopy)
		
	case "email":
		log.Printf("ℹ️ [%s] Email flow - calling completeAuth", phoneNumberCopy)
		
		err = CompleteAuthentication(
			privateKey,
			contractAddress,
			client,
			instance,
			phoneNumberCopy,
			result.PublicKey,
			result.Wallet,
		)
		
		if err != nil {
			log.Printf("⚠️ [%s] completeAuth failed: %v", phoneNumberCopy, err)
			return &result, err
		}
		
		log.Printf("✅ [%s] completeAuth successful", phoneNumberCopy)
		
	default:
		log.Printf("⚠️ [%s] Unknown botID: '%s'", phoneNumberCopy, botIDCopy)
	}

	return &result, nil
}

// ============================================================================
// 🔐 COMPLETE AUTHENTICATION
// ============================================================================

func CompleteAuthentication(
	privateKey *ecdsa.PrivateKey,
	contractAddress common.Address,
	client *ethclient.Client,
	instance *contract.Contract,
	phoneNumber string,
	publicKey string,
	userWalletAddress common.Address,
) error {
	fmt.Println("publicKey la:",publicKey)
	log.Printf("═══════════════════════════════════════")
	log.Printf("🔐 CompleteAuthentication START")
	log.Printf("   Phone: '%s'", phoneNumber)
	log.Printf("   Wallet: '%s'", userWalletAddress.Hex())
	log.Printf("═══════════════════════════════════════")

	message := fmt.Sprintf("Wallet address: %s is authorized", userWalletAddress.Hex())

	// ============================================================
	// 🔑 Generate ephemeral key pair
	// ============================================================
	ephemeralPrivKey, err := crypto.GenerateKey()
	if err != nil {
		return fmt.Errorf("error generating key: %w", err)
	}

	ephemeralPubKey := crypto.FromECDSAPub(&ephemeralPrivKey.PublicKey)
	ephemeralPubKeyHex := hex.EncodeToString(ephemeralPubKey)
	ephemeralPrivKeyBytes := crypto.FromECDSA(ephemeralPrivKey)

	log.Printf("🔑 Ephemeral key generated: %s...", ephemeralPubKeyHex[:16])

	// ============================================================
	// 🔐 Compute ECDH shared secret
	// ============================================================
	userPubKeyBytes, err := hex.DecodeString(strings.TrimPrefix(publicKey, "0x"))
	if err != nil {
		return fmt.Errorf("error decoding public key: %w", err)
	}

	sharedSecretHex, err := utils.ECDHSharedSecretHex(ephemeralPrivKeyBytes, userPubKeyBytes)
	if err != nil {
		return fmt.Errorf("error ECDH: %w", err)
	}

	sharedSecretBytes, err := hex.DecodeString(sharedSecretHex)
	if err != nil {
		return fmt.Errorf("error decoding secret: %w", err)
	}

	log.Printf("🔐 Shared secret computed")

	// ============================================================
	// 🔒 Encrypt message
	// ============================================================
	iv := make([]byte, 16)
	if _, err := rand.Read(iv); err != nil {
		return fmt.Errorf("error generating IV: %w", err)
	}

	encryptedBytes, err := utils.EncryptAESCBC(sharedSecretBytes, []byte(message), iv)
	if err != nil {
		return fmt.Errorf("error encrypting: %w", err)
	}

	log.Printf("🔐 Message encrypted")

	encryptedMessageBase64 := base64.StdEncoding.EncodeToString(encryptedBytes)
	ivHex := hex.EncodeToString(iv)

	// ============================================================
	// 💾 Save to log file
	// ============================================================
	err = saveEncryptedDataToLog(
		userWalletAddress.Hex(),
		phoneNumber,
		encryptedMessageBase64,
		ephemeralPubKeyHex,
		ivHex,
	)
	if err != nil {
		log.Printf("⚠️ Could not save log: %v", err)
	}

	// ============================================================
	// 🔑 Setup transactor for completeAuthentication
	// ============================================================
	publicKeyECDSA := privateKey.Public().(*ecdsa.PublicKey)
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	
	chainID := big.NewInt(991)
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		return fmt.Errorf("failed to create transactor: %w", err)
	}

	auth.From = fromAddress
	auth.Value = big.NewInt(0)
	auth.GasLimit = uint64(5_000_000)
	auth.GasPrice = big.NewInt(1_000_000_000)

	// ============================================================
	// 📤 Send completeAuthentication transaction
	// ============================================================
	log.Printf("📤 Sending completeAuthentication...")
	
	tx, err := instance.CompleteAuthentication(
		auth,
		phoneNumber,
		encryptedBytes,
		ephemeralPubKey,
	)
	
	if err != nil {
		return fmt.Errorf("completeAuth tx failed: %w", err)
	}

	log.Printf("📝 Tx hash: %s", tx.Hash().Hex())

	// ============================================================
	// ⏳ Wait for confirmation
	// ============================================================
	ctxTimeout, cancel := context.WithTimeout(context.Background(), 3*time.Minute)
	defer cancel()

	receipt, err := bind.WaitMined(ctxTimeout, client, tx)
	if err != nil {
		return fmt.Errorf("wait mined failed: %w", err)
	}

	if receipt.Status != 1 {
		return fmt.Errorf("completeAuth failed with status: %d", receipt.Status)
	}

	log.Printf("✅ completeAuthentication SUCCESS")
	log.Printf("═══════════════════════════════════════")
	
	return nil
}

// ============================================================================
// 💾 SAVE LOG
// ============================================================================

func saveEncryptedDataToLog(
	userWalletAddress string,
	phoneNumber string,
	encryptedMessage string,
	ephemeralPubKey string,
	iv string,
) error {
	logDir := "log"
	if _, err := os.Stat(logDir); os.IsNotExist(err) {
		if err := os.Mkdir(logDir, 0755); err != nil {
			return err
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
		return err
	}

	fileName := fmt.Sprintf("%s_%s.json", userWalletAddress, time.Now().Format("20060102150405"))
	filePath := filepath.Join(logDir, fileName)

	return os.WriteFile(filePath, fileContent, 0644)
}