package app

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/mail"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/meta-node-blockchain/verify_server/config"
	"github.com/meta-node-blockchain/verify_server/utils"
	"github.com/meta-node-blockchain/meta-node/cmd/client"
	c_config "github.com/meta-node-blockchain/meta-node/cmd/client/pkg/config"
	service "github.com/meta-node-blockchain/verify_server/service"
)

// AppContext chứa tất cả dependencies
type AppContext struct {
	MetaClient   *client.Client
	AdminAddress common.Address
	ContractAddr string
	ContractABI  string
	RpcURL       string
}

// NewAppContext khởi tạo AppContext từ config
func NewAppContext() (*AppContext, error) {
	cfg, err := config.LoadConfig("config.yaml")
	if err != nil {
		return nil, fmt.Errorf("error loading config: %w", err)
	}
	
	var contractABI string
	abiFilePath := cfg.AuthAbiPath
	contractAddr := cfg.AuthAddress
	rpcURL := cfg.RpcURL
	
	if abiFilePath != "" {
		contractABI, err = utils.ReadABIFromFile(abiFilePath)
		if err != nil {
			return nil, fmt.Errorf("error reading ABI file: %w", err)
		}
		log.Printf("✅ Loaded ABI from file: %s", abiFilePath)
	}

	metaClient, err := client.NewClient(
		&c_config.ClientConfig{
			Version_:                cfg.MetaNodeVersion,
			PrivateKey_:             cfg.PrivateKeyAdmin,
			ParentAddress:           cfg.AdminAddress,
			ParentConnectionAddress: cfg.ParentConnectionAddress,
			ConnectionAddress_:      cfg.ConnectionAddress_,
			ParentConnectionType:    cfg.ParentConnectionType,
			ChainId:                 cfg.ChainId,
		},
	)
	if err != nil {
		return nil, fmt.Errorf("error creating meta-node client: %w", err)
	}

	adminAddress := common.HexToAddress(cfg.AdminAddress)
	log.Println("✅ Meta-Node client initialized successfully")
	
	return &AppContext{
		MetaClient:   metaClient,
		AdminAddress: adminAddress,
		ContractAddr: contractAddr,
		ContractABI:  contractABI,
		RpcURL:       rpcURL,
	}, nil
}

// ============================================
// AUTHENTICATION HANDLER
// ============================================

// func (ctx *AppContext) handleAuthenticationEmail(identifier string, otpString string) (bool, error) {
// 	log.Printf("[Auth] 🔐 Processing authentication for: %s with OTP: %s", identifier, otpString)

// 	service.CheckOTP(
// 		ctx.AdminAddress,
// 		ctx.MetaClient,
// 		ctx.ContractAddr,
// 		ctx.ContractABI,
// 		ctx.RpcURL,
// 		identifier,
// 		otpString,
// 		"email",
// 	)

// 	log.Printf("[Auth] ✅ Authentication request sent for: %s", identifier)
// 	return true, nil
// }
func (ctx *AppContext) handleAuthenticationEmail(identifier string, otpString string) (bool, error) {
	log.Printf("[Auth] 🔐 Processing authentication for: %s with OTP: %s", identifier, otpString)

	// ✅ Nhận kết quả từ CheckOTP
	result, err := service.CheckOTP(
		ctx.AdminAddress,
		ctx.MetaClient,
		ctx.ContractAddr,
		ctx.ContractABI,
		ctx.RpcURL,
		identifier,
		otpString,
		"email",
	)

	if err != nil {
		log.Printf("[Auth] ❌ Authentication failed: %v", err)
		return false, fmt.Errorf("authentication failed: %w", err)
	}

	if result == nil {
		log.Printf("[Auth] ❌ No result returned from CheckOTP")
		return false, fmt.Errorf("no result from OTP validation")
	}

	log.Printf("[Auth] ✅ Authentication successful!")
	log.Printf("[Auth]    - Public Key: %s", result.PublicKey)
	log.Printf("[Auth]    - Wallet: %s", result.Wallet.Hex())
	
	return true, nil
}
// ============================================
// INBOUND EMAIL WEBHOOK HANDLERS
// ============================================

// InboundEmailData - Struct để parse email từ webhook
type InboundEmailData struct {
	From        string            `json:"from"`
	To          string            `json:"to"`
	Subject     string            `json:"subject"`
	Text        string            `json:"text"`          // For some providers
	TextBody    string            `json:"text_body"`     // ✅ Add this for your provider
	HTML        string            `json:"html"`
	HTMLBody    string            `json:"html_body"`     // ✅ Add this too
	Headers     map[string]string `json:"headers"`
	MessageID   string            `json:"message_id"`    // ✅ Optional but useful
	RawEmail    string            `json:"raw_email"`
	Attachments []interface{}     `json:"attachments"`   // ✅ Optional
}
func (ctx *AppContext) MakeInboundEmailWebhookHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("📨 ========================================")
		log.Println("📨 INCOMING EMAIL WEBHOOK")
		log.Println("📨 ========================================")

		// Chỉ chấp nhận POST
		if r.Method != http.MethodPost {
			log.Printf("❌ Invalid method: %s", r.Method)
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		// Đọc raw body
		bodyBytes, err := io.ReadAll(r.Body)
		if err != nil {
			log.Printf("❌ Error reading body: %v", err)
			http.Error(w, "Failed to read request body", http.StatusBadRequest)
			return
		}
		defer r.Body.Close()

		// ✅ TRẢ RESPONSE NGAY LẬP TỨC - QUAN TRỌNG!
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{
			"status":  "accepted",
			"message": "Email received successfully",
		})

		// ✅ XỬ LÝ EMAIL TRONG GOROUTINE (không block response)
		go ctx.processIncomingEmail(bodyBytes, r.Header.Get("Content-Type"))
	}
}

// ✅ Tách xử lý email ra hàm riêng
func (ctx *AppContext) processIncomingEmail(bodyBytes []byte, contentType string) {
	log.Printf("📦 Processing email asynchronously...")
	log.Printf("📦 Raw Body Length: %d bytes", len(bodyBytes))

	var emailData InboundEmailData

	// Parse theo content type
	if strings.Contains(contentType, "application/json") {
		log.Println("🔍 Parsing as JSON...")
		err := json.Unmarshal(bodyBytes, &emailData)
		if err != nil {
			log.Printf("❌ JSON parse error: %v", err)
			return
		}
		
		// ✅ Handle different field names
		if emailData.Text == "" && emailData.TextBody != "" {
			emailData.Text = emailData.TextBody
		}
		if emailData.HTML == "" && emailData.HTMLBody != "" {
			emailData.HTML = emailData.HTMLBody
		}
	} else {
		log.Printf("⚠️ Unsupported content type: %s", contentType)
		return
	}

	log.Println("📧 ========================================")
	log.Println("📧 PARSED EMAIL DATA:")
	log.Println("📧 ========================================")
	log.Printf("   From:    %s", emailData.From)
	log.Printf("   To:      %s", emailData.To)
	log.Printf("   Subject: %s", emailData.Subject)
	log.Printf("   Text:    %s", emailData.Text)
	log.Println("📧 ========================================")

	// Parse email addresses
	senderEmail, err := extractEmailAddress(emailData.From)
	if err != nil {
		log.Printf("❌ Invalid sender email: %v", err)
		return
	}

	recipientEmail, err := extractEmailAddress(emailData.To)
	if err != nil {
		log.Printf("❌ Invalid recipient email: %v", err)
		return
	}

	// Clean data
	cleanSubject := strings.TrimSpace(emailData.Subject)
	otpString := emailData.Text
	if otpString == "" {
		otpString = emailData.TextBody
	}
	// ✅ Loại bỏ TẤT CẢ whitespace characters
	otpString = strings.TrimSpace(otpString)           // Trim đầu/cuối
	otpString = strings.ReplaceAll(otpString, "\r\n", "") // Windows line ending
	otpString = strings.ReplaceAll(otpString, "\n", "")   // Unix line ending
	otpString = strings.ReplaceAll(otpString, "\r", "")   // Old Mac line ending
	otpString = strings.ReplaceAll(otpString, " ", "")    // Spaces
	otpString = strings.ReplaceAll(otpString, "\t", "")   // Tabs
	otpString = strings.TrimSpace(otpString)           // Trim lại lần nữa để chắc chắn

	log.Printf("   OTP/Body (raw):    '%s' (len: %d)", emailData.TextBody, len(emailData.TextBody))
	log.Printf("   OTP/Body (cleaned): '%s' (len: %d)", otpString, len(otpString))

	log.Println("🔍 ========================================")
	log.Println("🔍 EXTRACTED DATA:")
	log.Println("🔍 ========================================")
	log.Printf("   Clean Sender:    %s", senderEmail)
	log.Printf("   Clean Recipient: %s", recipientEmail)
	log.Printf("   Clean Subject:   %s", cleanSubject)
	log.Printf("   OTP/Body:        '%s' (len: %d)", otpString, len(otpString))
	log.Println("🔍 ========================================")

	// ============================================
	// KIỂM TRA AUTHENTICATION EMAIL
	// ============================================
	if cleanSubject == "" {
		log.Println("🔐 ========================================")
		log.Println("🔐 AUTHENTICATION EMAIL DETECTED!")
		log.Println("🔐 ========================================")
		log.Printf("🔐 Sender: %s", senderEmail)
		log.Printf("🔐 OTP: %s", otpString)
		
		success, err := ctx.handleAuthenticationEmail(senderEmail, otpString)
		if err != nil {
			log.Printf("❌ Error processing authentication: %v", err)
			return
		}
		if success {
			log.Println("✅ ========================================")
			log.Println("✅ AUTHENTICATION SUCCESSFUL!")
			log.Println("✅ ========================================")
		}
		return
	}

	// ============================================
	// EMAIL THƯỜNG
	// ============================================
	log.Println("📧 Normal email received, storing...")
	
	password, err := utils.GeneratePassword(recipientEmail)
	if err != nil {
		log.Printf("⚠️  Warning: Failed to generate password: %v", err)
	} else {
		emailContent := fmt.Sprintf("From: %s\nTo: %s\nSubject: %s\n\n%s", 
			emailData.From, emailData.To, emailData.Subject, emailData.Text)
		
		encryptedEmail, err := utils.EncryptEmail(emailContent, password)
		if err != nil {
			log.Printf("⚠️  Warning: Failed to encrypt email: %v", err)
		} else {
			err = utils.SaveEmailLocally(encryptedEmail)
			if err != nil {
				log.Printf("⚠️  Warning: Failed to save email: %v", err)
			} else {
				log.Println("✅ Email encrypted and saved successfully")
			}
		}
	}

	log.Println("✅ ========================================")
	log.Println("✅ EMAIL PROCESSING COMPLETED")
	log.Println("✅ ========================================")
}
// MakeInboundEmailWebhookHandler - Handler nhận email qua HTTP POST
// func (ctx *AppContext) MakeInboundEmailWebhookHandler() http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		log.Println("📨 ========================================")
// 		log.Println("📨 INCOMING EMAIL WEBHOOK")
// 		log.Println("📨 ========================================")

// 		// Chỉ chấp nhận POST
// 		if r.Method != http.MethodPost {
// 			log.Printf("❌ Invalid method: %s", r.Method)
// 			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
// 			return
// 		}

// 		// Log headers
// 		log.Println("📋 Request Headers:")
// 		for name, values := range r.Header {
// 			for _, value := range values {
// 				log.Printf("   %s: %s", name, value)
// 			}
// 		}

// 		// Đọc raw body
// 		bodyBytes, err := io.ReadAll(r.Body)
// 		if err != nil {
// 			log.Printf("❌ Error reading body: %v", err)
// 			http.Error(w, "Failed to read request body", http.StatusBadRequest)
// 			return
// 		}
// 		defer r.Body.Close()

// 		log.Printf("📦 Raw Body Length: %d bytes", len(bodyBytes))
// 		log.Printf("📦 Raw Body Content:\n%s", string(bodyBytes))

// 		// Parse Content-Type để xử lý đúng format
// 		contentType := r.Header.Get("Content-Type")
// 		log.Printf("📋 Content-Type: %s", contentType)

// 		var emailData InboundEmailData

// 		// ============================================
// 		// XỬ LÝ THEO ĐỊNH DẠNG
// 		// ============================================

// 		if strings.Contains(contentType, "application/json") {
// 			// Format 1: JSON (SendGrid Inbound Parse với JSON)
// 			log.Println("🔍 Parsing as JSON...")
// 			err = json.Unmarshal(bodyBytes, &emailData)
// 			if err != nil {
// 				log.Printf("❌ JSON parse error: %v", err)
// 				http.Error(w, "Invalid JSON format", http.StatusBadRequest)
// 				return
// 			}
// 			// ✅ Handle different field names from different email providers
// 			if emailData.Text == "" && emailData.TextBody != "" {
// 				emailData.Text = emailData.TextBody
// 			}
// 			if emailData.HTML == "" && emailData.HTMLBody != "" {
// 				emailData.HTML = emailData.HTMLBody
// 			}
// 		} else if strings.Contains(contentType, "multipart/form-data") || strings.Contains(contentType, "application/x-www-form-urlencoded") {
// 			// Format 2: Form data (SendGrid/Mailgun default)
// 			log.Println("🔍 Parsing as Form Data...")
// 			err = r.ParseMultipartForm(10 << 20) // 10MB max
// 			if err != nil {
// 				log.Printf("❌ Form parse error: %v", err)
// 				http.Error(w, "Failed to parse form", http.StatusBadRequest)
// 				return
// 			}

// 			emailData = InboundEmailData{
// 				From:    r.FormValue("from"),
// 				To:      r.FormValue("to"),
// 				Subject: r.FormValue("subject"),
// 				Text:    r.FormValue("text"),
// 				HTML:    r.FormValue("html"),
// 			}

// 			// Log tất cả form fields
// 			log.Println("📋 Form Fields:")
// 			for key, values := range r.Form {
// 				log.Printf("   %s: %v", key, values)
// 			}

// 		} else {
// 			// Format 3: Raw email
// 			log.Println("🔍 Treating as raw email...")
// 			emailData.RawEmail = string(bodyBytes)
// 		}

// 		// ============================================
// 		// LOG DỮ LIỆU EMAIL
// 		// ============================================

// 		log.Println("📧 ========================================")
// 		log.Println("📧 PARSED EMAIL DATA:")
// 		log.Println("📧 ========================================")
// 		log.Printf("   From:    %s", emailData.From)
// 		log.Printf("   To:      %s", emailData.To)
// 		log.Printf("   Subject: %s", emailData.Subject)
// 		log.Printf("   Text:    %s", emailData.Text)
// 		log.Printf("   HTML:    %s", emailData.HTML)
// 		if emailData.RawEmail != "" {
// 			log.Printf("   Raw Email (first 500 chars): %s", 
// 				truncateString(emailData.RawEmail, 500))
// 		}
// 		log.Println("📧 ========================================")

// 		// ============================================
// 		// XỬ LÝ EMAIL
// 		// ============================================

// 		// Parse email address từ "From" field
// 		senderEmail, err := extractEmailAddress(emailData.From)
// 		if err != nil {
// 			log.Printf("❌ Invalid sender email: %v", err)
// 			http.Error(w, "Invalid sender email format", http.StatusBadRequest)
// 			return
// 		}

// 		recipientEmail, err := extractEmailAddress(emailData.To)
// 		if err != nil {
// 			log.Printf("❌ Invalid recipient email: %v", err)
// 			http.Error(w, "Invalid recipient email format", http.StatusBadRequest)
// 			return
// 		}

// 		cleanSubject := strings.TrimSpace(emailData.Subject)
// 		// ✅ Extract and clean OTP - remove newlines and whitespace
// 		otpString := strings.TrimSpace(emailData.Text)
// 		otpString = strings.ReplaceAll(otpString, "\r\n", "")
// 		otpString = strings.ReplaceAll(otpString, "\n", "")
// 		otpString = strings.ReplaceAll(otpString, "\r", "")
// 		otpString = strings.TrimSpace(otpString)

// 		log.Println("🔍 ========================================")
// 		log.Println("🔍 EXTRACTED DATA:")
// 		log.Println("🔍 ========================================")
// 		log.Printf("   Clean Sender:    %s", senderEmail)
// 		log.Printf("   Clean Recipient: %s", recipientEmail)
// 		log.Printf("   Clean Subject:   %s", cleanSubject)
// 		log.Printf("   OTP/Body:        %s", otpString)
// 		log.Println("🔍 ========================================")

// 		// ============================================
// 		// KIỂM TRA AUTHENTICATION EMAIL (subject rỗng)
// 		// ============================================

// 		if cleanSubject == "" {
// 			log.Println("🔐 ========================================")
// 			log.Println("🔐 AUTHENTICATION EMAIL DETECTED!")
// 			log.Println("🔐 ========================================")
// 			log.Printf("🔐 Sender: %s", senderEmail)
// 			log.Printf("🔐 OTP: %s", otpString)
			
// 			success, err := ctx.handleAuthenticationEmail(senderEmail, otpString)
// 			if err != nil {
// 				log.Printf("❌ Error processing authentication: %v", err)
				
// 				// ✅ Return error details to client
// 				w.Header().Set("Content-Type", "application/json")
// 				w.WriteHeader(http.StatusBadRequest)
// 				json.NewEncoder(w).Encode(map[string]string{
// 					"status":  "error",
// 					"message": err.Error(),
// 					"sender":  senderEmail,
// 				})
// 				return
// 			}
// 			if success {
// 				log.Println("✅ ========================================")
// 				log.Println("✅ AUTHENTICATION SUCCESSFUL!")
// 				log.Println("✅ ========================================")
				
// 				w.Header().Set("Content-Type", "application/json")
// 				json.NewEncoder(w).Encode(map[string]string{
// 					"status":  "success",
// 					"message": "Authentication email processed successfully",
// 					"sender":  senderEmail,
// 				})
// 				return
// 			}
// 		}

// 		// ============================================
// 		// EMAIL THƯỜNG (có subject)
// 		// ============================================

// 		log.Println("📧 Normal email received, storing...")
		
// 		// Encrypt và lưu email (optional)
// 		password, err := utils.GeneratePassword(recipientEmail)
// 		if err != nil {
// 			log.Printf("⚠️  Warning: Failed to generate password: %v", err)
// 		} else {
// 			emailContent := fmt.Sprintf("From: %s\nTo: %s\nSubject: %s\n\n%s", 
// 				emailData.From, emailData.To, emailData.Subject, emailData.Text)
			
// 			encryptedEmail, err := utils.EncryptEmail(emailContent, password)
// 			if err != nil {
// 				log.Printf("⚠️  Warning: Failed to encrypt email: %v", err)
// 			} else {
// 				err = utils.SaveEmailLocally(encryptedEmail)
// 				if err != nil {
// 					log.Printf("⚠️  Warning: Failed to save email: %v", err)
// 				} else {
// 					log.Println("✅ Email encrypted and saved successfully")
// 				}
// 			}
// 		}

// 		// Response
// 		w.Header().Set("Content-Type", "application/json")
// 		json.NewEncoder(w).Encode(map[string]interface{}{
// 			"status":    "success",
// 			"message":   "Email received and processed",
// 			"sender":    senderEmail,
// 			"recipient": recipientEmail,
// 			"subject":   cleanSubject,
// 		})

// 		log.Println("✅ ========================================")
// 		log.Println("✅ EMAIL PROCESSING COMPLETED")
// 		log.Println("✅ ========================================")
// 	}
// }

// ============================================
// HELPER FUNCTIONS
// ============================================

// extractEmailAddress - Trích xuất email từ string "Name <email@domain.com>"
func extractEmailAddress(emailStr string) (string, error) {
	emailStr = strings.TrimSpace(emailStr)
	
	// Nếu đã là email thuần
	if !strings.Contains(emailStr, "<") {
		addr, err := mail.ParseAddress(emailStr)
		if err != nil {
			return emailStr, nil // Trả về nguyên bản nếu parse lỗi
		}
		return addr.Address, nil
	}
	
	// Parse "Name <email@domain.com>"
	addr, err := mail.ParseAddress(emailStr)
	if err != nil {
		return "", fmt.Errorf("invalid email format: %w", err)
	}
	
	return addr.Address, nil
}

// truncateString - Cắt string cho logging
func truncateString(s string, maxLen int) string {
	if len(s) <= maxLen {
		return s
	}
	return s[:maxLen] + "... (truncated)"
}