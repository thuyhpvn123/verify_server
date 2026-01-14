package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/meta-node-blockchain/verify_server/app"
	"github.com/meta-node-blockchain/verify_server/handler"
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
		
		fmt.Println("VERIFY_TOKEN:", VERIFY_TOKEN)
		
		if mode == "subscribe" && token == VERIFY_TOKEN {
			fmt.Println("Webhook verified successfully!")
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(challenge))
		} else {
			http.Error(w, "Verification failed", http.StatusForbidden)
		}
		
		if r.Method == http.MethodPost {
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
	// Initialize app context
	ctx, err := app.NewAppContext()
	if err != nil {
		log.Fatalf("❌ Failed to initialize application context: %v", err)
	}

	// ============================================
	// SETUP ROUTES
	// ============================================
	
	mux := http.NewServeMux()
	
	// Email webhook
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
			// WhatsApp message handler would go here
			w.WriteHeader(http.StatusOK)
		} else {
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		}
	})

	// Telegram webhook
	mux.HandleFunc("/received/telegram/message/@thuyabcbot",
		handler.HandlerTelegramMessage(
			ctx.PrivateKey,
			ctx.ContractAddr,
			ctx.RpcURL,
		))

	// Other handlers
	mux.HandleFunc("/emails", makeListEmailsHandler())
	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{
			"status":  "healthy",
			"service": "verify-server",
		})
	})

	// ============================================
	// START SERVER
	// ============================================
	
	log.Println("🚀 ========================================")
	log.Println("🚀 SERVER STARTED SUCCESSFULLY!")
	log.Println("🚀 ========================================")
	log.Println("🌐 HTTP Server:          port 9000")
	log.Println("📧 Email Webhook:        /webhook/email/inbound")
	log.Println("📱 WhatsApp webhook:     /webhook/whatsapp")
	log.Println("📱 Telegram webhook:     /received/telegram/message/@thuyabcbot")
	log.Println("📧 Email list:           /emails")
	log.Println("💚 Health check:         /health")
	log.Println("🚀 ========================================")

	if err := http.ListenAndServe(":9000", mux); err != nil {
		log.Fatalf("❌ Failed to start HTTP server: %s", err)
	}
}