package app

import (
	"context"
	"crypto/ecdsa"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/mail"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	
	"github.com/meta-node-blockchain/verify_server/config"
	"github.com/meta-node-blockchain/verify_server/service"
	"github.com/meta-node-blockchain/verify_server/utils"
)

// ============================================================================
// 🔧 APP CONTEXT
// ============================================================================

type AppContext struct {
	PrivateKey   *ecdsa.PrivateKey
	AdminAddress common.Address
	ContractAddr common.Address
	RpcURL       string
}

func NewAppContext() (*AppContext, error) {
	cfg, err := config.LoadConfig("config.yaml")
	if err != nil {
		return nil, fmt.Errorf("error loading config: %w", err)
	}
	
	// Load private key
	privateKey, err := crypto.HexToECDSA(cfg.PrivateKeyAdmin)
	if err != nil {
		return nil, fmt.Errorf("error loading private key: %w", err)
	}

	// Derive address from private key
	publicKey := privateKey.Public().(*ecdsa.PublicKey)
	adminAddress := crypto.PubkeyToAddress(*publicKey)
	
	contractAddr := common.HexToAddress(cfg.AuthAddress)
	rpcURL := cfg.RpcURL

	log.Printf("✅ Admin address: %s", adminAddress.Hex())
	log.Printf("✅ Contract address: %s", contractAddr.Hex())
	log.Printf("✅ RPC URL: %s", rpcURL)
	
	return &AppContext{
		PrivateKey:   privateKey,
		AdminAddress: adminAddress,
		ContractAddr: contractAddr,
		RpcURL:       rpcURL,
	}, nil
}

// ============================================================================
// 🔧 EMAIL REQUEST CONTEXT
// ============================================================================

type EmailRequestContext struct {
	ID           string
	PrivateKey   *ecdsa.PrivateKey
	ContractAddr common.Address
	RpcURL       string
}

// ============================================================================
// 📧 EMAIL DATA TYPES
// ============================================================================

type InboundEmailData struct {
	From        string            `json:"from"`
	To          string            `json:"to"`
	Subject     string            `json:"subject"`
	Text        string            `json:"text"`
	TextBody    string            `json:"text_body"`
	HTML        string            `json:"html"`
	HTMLBody    string            `json:"html_body"`
	Headers     map[string]string `json:"headers"`
	MessageID   string            `json:"message_id"`
	RawEmail    string            `json:"raw_email"`
	Attachments []interface{}     `json:"attachments"`
}

// ============================================================================
// 📧 EMAIL WEBHOOK HANDLER
// ============================================================================

func (ctx *AppContext) MakeInboundEmailWebhookHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("📨 ========================================")
		log.Println("📨 INCOMING EMAIL WEBHOOK")
		log.Println("📨 ========================================")

		if r.Method != http.MethodPost {
			log.Printf("❌ Invalid method: %s", r.Method)
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		// Read body
		bodyBytes, err := io.ReadAll(r.Body)
		if err != nil {
			log.Printf("❌ Error reading body: %v", err)
			http.Error(w, "Failed to read request body", http.StatusBadRequest)
			return
		}
		defer r.Body.Close()

		// Make a copy for goroutine
		bodyCopy := make([]byte, len(bodyBytes))
		copy(bodyCopy, bodyBytes)

		contentType := r.Header.Get("Content-Type")

		// Return response immediately
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{
			"status":  "accepted",
			"message": "Email received successfully",
		})

		// Create isolated context
		requestCtx := EmailRequestContext{
			ID:           fmt.Sprintf("EMAIL-%d", time.Now().UnixNano()),
			PrivateKey:   ctx.PrivateKey,
			ContractAddr: ctx.ContractAddr,
			RpcURL:       ctx.RpcURL,
		}

		// Process in goroutine
		go processIncomingEmail(requestCtx, bodyCopy, contentType)
	}
}

// ============================================================================
// 📧 EMAIL PROCESSOR
// ============================================================================

func processIncomingEmail(ctx EmailRequestContext, bodyBytes []byte, contentType string) {
	log.Printf("📦 [%s] Processing email asynchronously...", ctx.ID)
	log.Printf("📦 [%s] Raw Body Length: %d bytes", ctx.ID, len(bodyBytes))

	var emailData InboundEmailData

	if strings.Contains(contentType, "application/json") {
		log.Printf("🔍 [%s] Parsing as JSON...", ctx.ID)
		err := json.Unmarshal(bodyBytes, &emailData)
		if err != nil {
			log.Printf("❌ [%s] JSON parse error: %v", ctx.ID, err)
			return
		}
		
		if emailData.Text == "" && emailData.TextBody != "" {
			emailData.Text = emailData.TextBody
		}
		if emailData.HTML == "" && emailData.HTMLBody != "" {
			emailData.HTML = emailData.HTMLBody
		}
	} else {
		log.Printf("⚠️ [%s] Unsupported content type: %s", ctx.ID, contentType)
		return
	}

	log.Printf("📧 [%s] From: %s, To: %s, Subject: %s", ctx.ID, emailData.From, emailData.To, emailData.Subject)

	// Extract email addresses
	senderEmail, err := extractEmailAddress(emailData.From)
	if err != nil {
		log.Printf("❌ [%s] Invalid sender email: %v", ctx.ID, err)
		return
	}

	recipientEmail, err := extractEmailAddress(emailData.To)
	if err != nil {
		log.Printf("❌ [%s] Invalid recipient email: %v", ctx.ID, err)
		return
	}

	cleanSubject := strings.TrimSpace(emailData.Subject)
	
	// Extract and clean OTP
	otpString := emailData.Text
	if otpString == "" {
		otpString = emailData.TextBody
	}
	
	otpString = strings.TrimSpace(otpString)
	otpString = strings.ReplaceAll(otpString, "\r\n", "")
	otpString = strings.ReplaceAll(otpString, "\n", "")
	otpString = strings.ReplaceAll(otpString, "\r", "")
	otpString = strings.ReplaceAll(otpString, " ", "")
	otpString = strings.ReplaceAll(otpString, "\t", "")

	log.Printf("🔍 [%s] Sender: %s, OTP: '%s' (len: %d)", ctx.ID, senderEmail, otpString, len(otpString))

	// Check if this is authentication email (empty subject)
	if cleanSubject == "" {
		log.Printf("🔐 [%s] AUTHENTICATION EMAIL DETECTED!", ctx.ID)
		log.Printf("🔐 [%s] Sender: %s, OTP: %s", ctx.ID, senderEmail, otpString)
		
		// Handle authentication
		success, err := handleAuthenticationEmail(ctx, senderEmail, otpString)
		if err != nil {
			log.Printf("❌ [%s] Error: %v", ctx.ID, err)
			return
		}
		if success {
			log.Printf("✅ [%s] AUTHENTICATION SUCCESSFUL!", ctx.ID)
		}
		return
	}

	// Normal email - store it
	log.Printf("📧 [%s] Normal email, storing...", ctx.ID)
	
	password, err := utils.GeneratePassword(recipientEmail)
	if err != nil {
		log.Printf("⚠️ [%s] Password gen failed: %v", ctx.ID, err)
	} else {
		emailContent := fmt.Sprintf("From: %s\nTo: %s\nSubject: %s\n\n%s", 
			emailData.From, emailData.To, emailData.Subject, emailData.Text)
		
		encryptedEmail, err := utils.EncryptEmail(emailContent, password)
		if err != nil {
			log.Printf("⚠️ [%s] Encrypt failed: %v", ctx.ID, err)
		} else {
			err = utils.SaveEmailLocally(encryptedEmail)
			if err != nil {
				log.Printf("⚠️ [%s] Save failed: %v", ctx.ID, err)
			} else {
				log.Printf("✅ [%s] Email saved", ctx.ID)
			}
		}
	}

	log.Printf("✅ [%s] EMAIL PROCESSING COMPLETED", ctx.ID)
}

// ============================================================================
// 🔐 AUTHENTICATION HANDLER
// ============================================================================

func handleAuthenticationEmail(ctx EmailRequestContext, identifier string, otpString string) (bool, error) {
	log.Printf("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	log.Printf("[Auth-%s] 🔐 STARTING EMAIL AUTHENTICATION", ctx.ID)
	log.Printf("[Auth-%s]    Identifier: %s", ctx.ID, identifier)
	log.Printf("[Auth-%s]    OTP: %s", ctx.ID, otpString)
	log.Printf("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")

	log.Printf("[Auth-%s] 📞 CALLING service.CheckOTP...", ctx.ID)
	log.Printf("[Auth-%s]    Contract: %s", ctx.ID, ctx.ContractAddr.Hex())
	log.Printf("[Auth-%s]    BotID: 'email'", ctx.ID)
	
	result, err := service.CheckOTP(
		context.Background(),
		ctx.PrivateKey,
		ctx.ContractAddr,
		ctx.RpcURL,
		identifier,
		otpString,
		"email", // ← botID
	)
	
	log.Printf("[Auth-%s] 📞 service.CheckOTP RETURNED", ctx.ID)
	log.Printf("[Auth-%s]    Error: %v", ctx.ID, err)
	log.Printf("[Auth-%s]    Result: %v", ctx.ID, result)

	if err != nil {
		log.Printf("[Auth-%s] ❌ Authentication failed: %v", ctx.ID, err)
		return false, fmt.Errorf("authentication failed: %w", err)
	}

	if result == nil {
		log.Printf("[Auth-%s] ❌ No result returned from CheckOTP", ctx.ID)
		return false, fmt.Errorf("no result from OTP validation")
	}

	log.Printf("[Auth-%s] ✅ Authentication successful!", ctx.ID)
	log.Printf("[Auth-%s]    - Public Key: %s", ctx.ID, result.PublicKey)
	log.Printf("[Auth-%s]    - Wallet: %s", ctx.ID, result.Wallet.Hex())
	log.Printf("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	
	return true, nil
}

// ============================================================================
// 🛠️ UTILITY FUNCTIONS
// ============================================================================

func extractEmailAddress(emailStr string) (string, error) {
	emailStr = strings.TrimSpace(emailStr)
	
	if !strings.Contains(emailStr, "<") {
		addr, err := mail.ParseAddress(emailStr)
		if err != nil {
			return emailStr, nil
		}
		return addr.Address, nil
	}
	
	addr, err := mail.ParseAddress(emailStr)
	if err != nil {
		return "", fmt.Errorf("invalid email format: %w", err)
	}
	
	return addr.Address, nil
}