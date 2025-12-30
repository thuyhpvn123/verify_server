package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/phires/go-guerrilla"
	"github.com/phires/go-guerrilla/backends"
	// guerrillaMail "github.com/phires/go-guerrilla/mail"

	handler "github.com/meta-node-blockchain/verify_server/handler"
	model "github.com/meta-node-blockchain/verify_server/model"
	"github.com/meta-node-blockchain/verify_server/app"

)

// func verifyWebhook(w http.ResponseWriter, r *http.Request) {
// 	err := godotenv.Load()
// 	if err != nil {
// 		log.Println("Warning: No .env file found")
// 	}
// 	VERIFY_TOKEN := os.Getenv("WHATSAPP_VERIFY_TOKEN")
// 	// Lấy các tham số từ query
// 	mode := r.URL.Query().Get("hub.mode")
// 	token := r.URL.Query().Get("hub.verify_token")
// 	challenge := r.URL.Query().Get("hub.challenge")

// 	// Kiểm tra xác minh webhook
// 	if mode == "subscribe" && token == VERIFY_TOKEN {
// 		fmt.Println("Webhook verified successfully!")
// 		w.WriteHeader(http.StatusOK)
// 		w.Write([]byte(challenge))
// 	} else {
// 		http.Error(w, "Verification failed", http.StatusForbidden)
// 	}
// }

// func main() {
// 	err := godotenv.Load(".env")
// 	if err != nil {
// 		log.Printf("lỗi khi tải tệp .env: %v", err)
// 	}
// 	// var method model.MessagingMethod = model.Telegram
// 	fmt.Println("Selected messaging method:", model.WhatsApp.Int())
// 	infuraURL := os.Getenv("INFURA_URL")
// 	contractAddress := os.Getenv("CONTRACT_ADDRESS")
// 	contractABI := os.Getenv("CONTRACT_ABI")
// 	fmt.Println("Server is running on port: 8080 ,")
// 	config, err := config.LoadConfig("config.yaml")
// 	if err != nil {
// 		log.Fatal("invalid configuration", err)
// 	}
// 	// Nếu không có option custom client, bạn có thể set global transport
// 	// http.DefaultTransport = customTransport
// 	client, err := client.NewClient(
// 		&c_config.ClientConfig{
// 			Version_:                config.MetaNodeVersion,
// 			PrivateKey_:             config.PrivateKeyAdmin,
// 			ParentAddress:           config.AdminAddress,
// 			ParentConnectionAddress: config.ParentConnectionAddress,
// 			// DnsLink_:                config.DnsLink(),
// 			ConnectionAddress_:   config.ConnectionAddress_,
// 			ParentConnectionType: config.ParentConnectionType,
// 			ChainId:              config.ChainId,
// 		},
// 	)
// 	if err != nil {
// 		logger.Error(fmt.Sprintf("error when create chain client %v", err))
// 	}

// 	http.HandleFunc("/webhook/whatsapp", func(w http.ResponseWriter, r *http.Request) {
// 		if r.Method == http.MethodGet {
// 			verifyWebhook(w, r) // Xác thực Webhook từ Meta
// 		} else if r.Method == http.MethodPost {
// 			// Gọi đúng handler được tạo từ ReceiveMessageWhatsapp
// 			handler.ReceiveMessageWhatsapp(common.HexToAddress(config.AdminAddress),client,contractAddress, contractABI, infuraURL)(w, r)
// 		} else {
// 			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
// 		}
// 	})

// 	// http.HandleFunc("/webhook/whatsapp", verifyWebhook)
// 	// http.HandleFunc("/webhook/whatsapp", handler.ReceiveMessageWhatsapp(contractAddress, contractABI, infuraURL))
// 	// http.HandleFunc("/received/message", handler.GetMessageTwilio(contractAddress, contractABI, infuraURL, model.WhatsApp.Int()))
// 	http.HandleFunc("/received/telegram/message/@thuyabcbot", handler.HandlerTelegramMessage(common.HexToAddress(config.AdminAddress),client,contractAddress, contractABI, infuraURL))
// 	http.ListenAndServe(":8080", nil)
// }
// ============================================
// DEPENDENCY INJECTION - APP CONTEXT
// ============================================

// ============================================
// HTTP HANDLERS (với closure để inject context)
// ============================================

func makeVerifyWebhookHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		VERIFY_TOKEN := os.Getenv("WHATSAPP_VERIFY_TOKEN")
		mode := r.URL.Query().Get("hub.mode")
		token := r.URL.Query().Get("hub.verify_token")
		challenge := r.URL.Query().Get("hub.challenge")

		if mode == "subscribe" && token == VERIFY_TOKEN {
			fmt.Println("Webhook verified successfully!")
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(challenge))
		} else {
			http.Error(w, "Verification failed", http.StatusForbidden)
		}
	}
}

func makeListEmailsHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		files, err := filepath.Glob("./emails/email_*.txt.gz")
		if err != nil {
			http.Error(w, "Failed to read email directory", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(files)
	}
}


// ============================================
// MAIN FUNCTION
// ============================================

func main() {
	// Khởi tạo AppContext (dependency injection container)
	ctx, err := app.NewAppContext()
	if err != nil {
		log.Fatalf("❌ Failed to initialize application context: %v", err)
	}

	fmt.Println("Selected messaging method:", model.WhatsApp.Int())

	// ============================================
	// SMTP SERVER SETUP
	// ============================================

	smtpConfig := &guerrilla.AppConfig{
		LogFile:      "./go-guerrilla.log",
		LogLevel:     "error",
		AllowedHosts: []string{"m.pro", "payws.net", "payws.com", "metanode.co"},
		Servers: []guerrilla.ServerConfig{
			{
				IsEnabled:       true,
				ListenInterface: "0.0.0.0:2025",
				MaxClients:      5,
				Timeout:         100,
			},
		},
		BackendConfig: backends.BackendConfig{
			"save_process":       "MyFooProcessor",
			"validate_process":   "MyFooProcessor",
			"save_workers_size":  1,
			"log_received_mails": false,
		},
	}

	d := guerrilla.Daemon{Config: smtpConfig}
	
	// Inject context vào SMTP processor
	d.AddProcessor("MyFooProcessor", ctx.CreateSMTPProcessor())

	if err := d.Start(); err != nil {
		log.Fatalf("❌ Failed to start SMTP server: %s", err)
	}

	// ============================================
	// HTTP SERVER SETUP (với dependency injection)
	// ============================================

	// WhatsApp webhook
	http.HandleFunc("/webhook/whatsapp", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			makeVerifyWebhookHandler()(w, r)
		} else if r.Method == http.MethodPost {
			handler.ReceiveMessageWhatsapp(
				ctx.AdminAddress, 
				ctx.MetaClient, 
				ctx.ContractAddr, 
				ctx.ContractABI, 
				ctx.RpcURL,
			)(w, r)
		} else {
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		}
	})

	// Telegram webhook
	http.HandleFunc("/received/telegram/message/@thuyabcbot",
		handler.HandlerTelegramMessage(
			ctx.AdminAddress, 
			ctx.MetaClient, 
			ctx.ContractAddr, 
			ctx.ContractABI, 
			ctx.RpcURL,
		))

	// Email reading endpoints (inject context)
	http.HandleFunc("/emails", makeListEmailsHandler())
	http.HandleFunc("/emails/", ctx.MakeReadEmailHandler())

	// ============================================
	// START SERVERS
	// ============================================

	log.Println("🚀 ========================================")
	log.Println("🚀 SERVER STARTED SUCCESSFULLY!")
	log.Println("🚀 ========================================")
	log.Println("📧 SMTP Server:        port 2025")
	log.Println("🌐 HTTP Server:        port 9000")
	log.Println("📱 WhatsApp webhook:   /webhook/whatsapp")
	log.Println("📱 Telegram webhook:   /received/telegram/message/@thuyabcbot")
	log.Println("📧 Email list:         /emails")
	log.Println("📧 Email read:         /emails/{filename}?recipient={email}")
	log.Println("🚀 ========================================")

	if err := http.ListenAndServe(":9000", nil); err != nil {
		log.Fatalf("❌ Failed to start HTTP server: %s", err)
	}
}