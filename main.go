package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"

	handler "github.com/meta-node-blockchain/verify_server/handler"
	// model "github.com/meta-node-blockchain/verify_server/model"
	"github.com/meta-node-blockchain/verify_server/app"
)

// ============================================
// HTTP HANDLERS
// ============================================

func makeVerifyWebhookHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		VERIFY_TOKEN := os.Getenv("WHATSAPP_VERIFY_TOKEN")
		mode := r.URL.Query().Get("hub.mode")
		token := r.URL.Query().Get("hub.verify_token")
		challenge := r.URL.Query().Get("hub.challenge")
        fmt.Println("VERIFY_TOKEN:",VERIFY_TOKEN)
		if mode == "subscribe" && token == VERIFY_TOKEN {
			fmt.Println("Webhook verified successfully!")
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(challenge))
		} else {
			http.Error(w, "Verification failed", http.StatusForbidden)
		}
        if r.Method == http.MethodPost {
            // xử lý webhook message/status
            w.WriteHeader(http.StatusOK)
            return
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
    ctx, err := app.NewAppContext()
    if err != nil {
        log.Fatalf("❌ Failed to initialize application context: %v", err)
    }

    // ============================================
    // SETUP ROUTES
    // ============================================
    
    mux := http.NewServeMux()
    
    // Email webhook - accept cả root và specific path
    emailHandler := ctx.MakeInboundEmailWebhookHandler()
    mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        if r.URL.Path == "/" {
            emailHandler(w, r)
        } else {
            http.NotFound(w, r)
        }
    })
    mux.HandleFunc("/webhook/email/inbound", emailHandler)

    // WhatsApp webhook
    mux.HandleFunc("/webhook/whatsapp", func(w http.ResponseWriter, r *http.Request) {
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

    // Other handlers...
    mux.HandleFunc("/received/telegram/message/@thuyabcbot",
        handler.HandlerTelegramMessage(
            ctx.AdminAddress,
            ctx.MetaClient,
            ctx.ContractAddr,
            ctx.ContractABI,
            ctx.RpcURL,
        ))
    mux.HandleFunc("/emails", makeListEmailsHandler())
    mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Content-Type", "application/json")
        json.NewEncoder(w).Encode(map[string]string{
            "status":  "healthy",
            "service": "verify-server",
        })
    })

    log.Println("🚀 Server starting on :9000")
    if err := http.ListenAndServe(":9000", mux); err != nil {
        log.Fatalf("❌ Failed to start HTTP server: %s", err)
    }
}
// func main() {
// 	// Khởi tạo AppContext
// 	ctx, err := app.NewAppContext()
// 	if err != nil {
// 		log.Fatalf("❌ Failed to initialize application context: %v", err)
// 	}

// 	fmt.Println("Selected messaging method:", model.WhatsApp.Int())

// 	// ============================================
// 	// HTTP SERVER SETUP
// 	// ============================================

// 	// 🆕 EMAIL WEBHOOK (thay thế SMTP port 25)
// 	http.HandleFunc("/webhook/email/inbound", ctx.MakeInboundEmailWebhookHandler())

// 	// WhatsApp webhook
// 	http.HandleFunc("/webhook/whatsapp", func(w http.ResponseWriter, r *http.Request) {
// 		if r.Method == http.MethodGet {
// 			makeVerifyWebhookHandler()(w, r)
// 		} else if r.Method == http.MethodPost {
// 			handler.ReceiveMessageWhatsapp(
// 				ctx.AdminAddress,
// 				ctx.MetaClient,
// 				ctx.ContractAddr,
// 				ctx.ContractABI,
// 				ctx.RpcURL,
// 			)(w, r)
// 		} else {
// 			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
// 		}
// 	})

// 	// Telegram webhook
// 	http.HandleFunc("/received/telegram/message/@thuyabcbot",
// 		handler.HandlerTelegramMessage(
// 			ctx.AdminAddress,
// 			ctx.MetaClient,
// 			ctx.ContractAddr,
// 			ctx.ContractABI,
// 			ctx.RpcURL,
// 		))

// 	// Email reading endpoints
// 	http.HandleFunc("/emails", makeListEmailsHandler())

// 	// Health check
// 	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
// 		w.Header().Set("Content-Type", "application/json")
// 		json.NewEncoder(w).Encode(map[string]string{
// 			"status": "healthy",
// 			"service": "verify-server",
// 		})
// 	})

// 	// ============================================
// 	// START SERVER
// 	// ============================================

// 	log.Println("🚀 ========================================")
// 	log.Println("🚀 SERVER STARTED SUCCESSFULLY!")
// 	log.Println("🚀 ========================================")
// 	log.Println("🌐 HTTP Server:          port 9000")
// 	log.Println("📧 Email Webhook:        /webhook/email/inbound")
// 	log.Println("📱 WhatsApp webhook:     /webhook/whatsapp")
// 	log.Println("📱 Telegram webhook:     /received/telegram/message/@thuyabcbot")
// 	log.Println("📧 Email list:           /emails")
// 	log.Println("💚 Health check:         /health")
// 	log.Println("🚀 ========================================")
// 	log.Println("")
// 	log.Println("📌 Next steps:")
// 	log.Println("   1. Expose với Ngrok: ngrok http 9000")
// 	log.Println("   2. Config SendGrid/Mailgun Inbound Parse:")
// 	log.Println("      URL: https://YOUR-NGROK-URL.ngrok.io/webhook/email/inbound")
// 	log.Println("🚀 ========================================")

// 	if err := http.ListenAndServe(":9000", nil); err != nil {
// 		log.Fatalf("❌ Failed to start HTTP server: %s", err)
// 	}
// }