package app

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/mail"
	"os"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	// "github.com/joho/godotenv"
	"github.com/meta-node-blockchain/verify_server/config"
	"github.com/meta-node-blockchain/verify_server/utils"
	"github.com/phires/go-guerrilla/backends"
	guerrillaMail "github.com/phires/go-guerrilla/mail"
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
	RpcURL    string
}

// NewAppContext khởi tạo AppContext từ config
func NewAppContext() (*AppContext, error) {
	// err := godotenv.Load()
	// if err != nil {
	// 	log.Println("⚠️  Warning: No .env file found")
	// }

	// rpcURL := os.Getenv("RPC_URL")
	// contractAddr := os.Getenv("AUTHSC_ADDRESS")
	// // contractABI := os.Getenv("CONTRACT_ABI")
	// // ✅ ĐỌC ABI TỪ FILE JSON
	// abiFilePath := os.Getenv("CONTRACT_ABI_FILE") // Ví dụ: "contracts/MyContract.json"
	cfg, err := config.LoadConfig("config.yaml")
	if err != nil {
		return nil, fmt.Errorf("error loading config: %w", err)
	}
	var contractABI string
	abiFilePath := cfg.AuthAbiPath
	contractAddr := cfg.AuthAddress
	rpcURL := cfg.RpcURL
	if abiFilePath != "" {
		// Đọc từ file
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
		RpcURL:    rpcURL,
	}, nil
}



// ============================================
// AUTHENTICATION HANDLER
// ============================================

func (ctx *AppContext) handleAuthenticationEmail(identifier string, otpString string) (bool, error) {
	log.Printf("[Auth] 🔐 Processing authentication for: %s with OTP: %s", identifier, otpString)

	service.CheckOTP(
		ctx.AdminAddress,
		ctx.MetaClient,
		ctx.ContractAddr,
		ctx.ContractABI,
		ctx.RpcURL,
		identifier,
		otpString,
		"email",
	)

	log.Printf("[Auth] ✅ Authentication request sent for: %s", identifier)
	return true, nil
}

func (ctx *AppContext) MakeReadEmailHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fileName := strings.TrimPrefix(r.URL.Path, "/emails/")
		recipient := r.URL.Query().Get("recipient")

		if fileName == "" || recipient == "" {
			http.Error(w, "Please provide fileName and recipient email in the URL", http.StatusBadRequest)
			return
		}

		encryptedData, err := os.ReadFile(fileName)
		if err != nil {
			http.Error(w, "Failed to read encrypted email file", http.StatusInternalServerError)
			return
		}

		password, err := utils.GeneratePassword(recipient)
		if err != nil {
			http.Error(w, "Failed to generate password for decryption", http.StatusInternalServerError)
			return
		}

		decryptedRawEmail, err := utils.DecryptEmail(encryptedData, password)
		if err != nil {
			http.Error(w, "Failed to decrypt email", http.StatusInternalServerError)
			return
		}

		parsedEmail, err := utils.ParseEmail(decryptedRawEmail)
		if err != nil {
			http.Error(w, "Failed to parse decrypted email", http.StatusInternalServerError)
			return
		}

		msg, _ := mail.ReadMessage(strings.NewReader(decryptedRawEmail))
		senderRaw := msg.Header.Get("From")
		addr, err := mail.ParseAddress(senderRaw)
		if err != nil {
			log.Printf("Error parsing sender email: %v", err)
			http.Error(w, "Invalid sender email format", http.StatusBadRequest)
			return
		}
		cleanSenderEmail := addr.Address

		cleanSubject := strings.TrimSpace(parsedEmail.Subject)
		otpString := strings.TrimSpace(parsedEmail.Body)

		fmt.Printf("\n--- DECRYPTED EMAIL DATA ---\n")
		fmt.Printf("  Sender: %s\n", cleanSenderEmail)
		fmt.Printf("  Recipient: %s\n", recipient)
		fmt.Printf("  Subject: %s\n", cleanSubject)
		fmt.Printf("  OTP: %s\n", otpString)
		fmt.Println("----------------------------")

		if cleanSubject == "" {
			log.Println("🔐 Authentication email detected, calling Smart Contract...")
			success, err := ctx.handleAuthenticationEmail(cleanSenderEmail, otpString)
			if err != nil {
				log.Printf("❌ Error calling smart contract: %v", err)
				http.Error(w, "Failed to call smart contract: "+err.Error(), http.StatusInternalServerError)
				return
			}

			if success {
				log.Println("✅ Smart Contract authentication successful!")
				w.Header().Set("Content-Type", "application/json")
				json.NewEncoder(w).Encode(map[string]string{"status": "Verification successful"})
				return
			} else {
				log.Println("❌ Smart Contract authentication failed.")
				http.Error(w, "Email verification failed on contract", http.StatusUnauthorized)
				return
			}
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]interface{}{
			"sender":    cleanSenderEmail,
			"recipient": recipient,
			"subject":   parsedEmail.Subject,
			"body":      parsedEmail.Body,
		})
	}
}

// ============================================
// SMTP PROCESSOR FACTORY
// ============================================
type myFooConfig struct {
		// SomeOption string `json:"some_option"` // Ví dụ về một cấu hình
	}
func (ctx *AppContext) CreateSMTPProcessor() func() backends.Decorator {
	return func() backends.Decorator {
		// config := &struct{}{}
		config := &myFooConfig{}

		initFunc := backends.InitializeWith(func(backendConfig backends.BackendConfig) error {
			// Trích xuất cấu hình từ backendConfig
			configType := backends.BaseConfig(&myFooConfig{})
			bcfg, err := backends.Svc.ExtractConfig(backendConfig, configType)
			if err != nil {
				return err
			}
			*config = *(bcfg.(*myFooConfig))
			return nil
		})

		backends.Svc.AddInitializer(initFunc)

		return func(p backends.Processor) backends.Processor {
			return backends.ProcessWith(func(e *guerrillaMail.Envelope, task backends.SelectTask) (backends.Result, error) {

				if task == backends.TaskValidateRcpt {
					if len(e.RcptTo) == 0 {
						return backends.NewResult("550 No recipient provided"), nil
					}
					recipient := e.RcptTo[0].String()
					recipientName := strings.Split(recipient, "@")[0]
					
					if recipientName == "" || !utils.IsValidRecipientName(recipientName) {
						log.Printf("Invalid recipient format: %s", recipient)
						return backends.NewResult("554 Invalid recipient email format"), nil
					}

					log.Printf("✅ Recipient validated: %s", recipientName)
					return backends.NewResult("250 Recipient OK"), nil
				}

				if task == backends.TaskSaveMail {
					if len(e.RcptTo) == 0 {
						return backends.NewResult("550 No recipient provided"), nil
					}
					
					recipient := e.RcptTo[0].String()
					recipientName := strings.Split(recipient, "@")[0]
					
					if recipientName == "" {
						log.Printf("Invalid recipient format: %s", recipient)
						return backends.NewResult("554 Invalid recipient email format"), nil
					}

					ip := e.RemoteIP
					sender := e.MailFrom.String()
					senderDomain := utils.ExtractDomain(sender)

					if len(e.Data.String()) > 1024*1024 {
						return backends.NewResult("552 Error: Message size exceeds 1MB limit"), nil
					}

					dkimResult, err := utils.CheckDKIM([]byte(e.Data.String()), senderDomain)
					if ip != "127.0.0.1" && !dkimResult {
						if err != nil {
							log.Printf("DKIM error: %v", err)
						}

						log.Printf("DKIM failed, fallback to SPF and DMARC checks")

						spfResult, spfErr := utils.CheckSPF(ip, senderDomain)
						if spfErr != nil || !spfResult {
							return backends.NewResult(fmt.Sprintf("554 SPF failed: %v", spfErr)), nil
						}

						dmarcResult, dmarcErr := utils.CheckDMARC(senderDomain)
						if dmarcErr != nil || !dmarcResult {
							return backends.NewResult(fmt.Sprintf("554 DMARC failed: %v", dmarcErr)), nil
						}
					}

					password, err := utils.GeneratePassword(recipient)
					if err != nil {
						log.Printf("Error generating password: %v", err)
						return backends.NewResult("554 Error generating password"), nil
					}

					encryptedEmail, err := utils.EncryptEmail(e.Data.String(), password)
					if err != nil {
						log.Printf("Error encrypting email: %v", err)
						return backends.NewResult("554 Error encrypting email"), nil
					}

					err = utils.SaveEmailLocally(encryptedEmail)
					if err != nil {
						log.Printf("Error saving email locally: %v", err)
						return backends.NewResult("554 Error saving email locally"), nil
					}

					parsedEmail, err := utils.ParseEmail(e.Data.String())
					if err != nil {
						log.Printf("Error parsing email: %v", err)
						return backends.NewResult("554 Error parsing email"), nil
					}

					addr, err := mail.ParseAddress(e.MailFrom.String())
					if err != nil {
						log.Printf("Error parsing sender address: %v", err)
						return backends.NewResult("554 Invalid sender address"), nil
					}
					senderEmail := addr.Address

					cleanSubject := strings.TrimSpace(parsedEmail.Subject)
					otpString := strings.TrimSpace(parsedEmail.Body)

					if cleanSubject == "" {
						log.Printf(">>> 🔐 Authentication email detected from [%s]", senderEmail)
						
						success, err := ctx.handleAuthenticationEmail(senderEmail, otpString)
						if err != nil {
							log.Printf("--> ❌ Error processing authentication for [%s]: %v", senderEmail, err)
							return backends.NewResult("554 Error processing authentication"), nil
						}
						
						if success {
							log.Printf("--> ✅ Authentication request sent for [%s]", senderEmail)
						}

						return backends.NewResult("250 OK: Authentication request received"), nil
					}

					log.Println(">>> Received normal email, continuing processing...")

					subject := parsedEmail.Subject
					body := utils.SanitizeEmailHTML(parsedEmail.Body)

					log.Printf("📧 Email received - Subject: %s, Body length: %d bytes", subject, len(body))
					return backends.NewResult("250 OK: Email received and stored successfully"), nil
				}
				
				return p.Process(e, task)
			})
		}
	}
}
