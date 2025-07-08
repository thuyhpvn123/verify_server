package main

import (
	"WhatsappVerifyOTP/handler"
	"WhatsappVerifyOTP/model"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func verifyWebhook(w http.ResponseWriter, r *http.Request) {
	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: No .env file found")
	}
	VERIFY_TOKEN := os.Getenv("WHATSAPP_VERIFY_TOKEN")
	// Lấy các tham số từ query
	mode := r.URL.Query().Get("hub.mode")
	token := r.URL.Query().Get("hub.verify_token")
	challenge := r.URL.Query().Get("hub.challenge")

	// Kiểm tra xác minh webhook
	if mode == "subscribe" && token == VERIFY_TOKEN {
		fmt.Println("Webhook verified successfully!")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(challenge))
	} else {
		http.Error(w, "Verification failed", http.StatusForbidden)
	}
}

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Printf("lỗi khi tải tệp .env: %v", err)
	}
	// var method model.MessagingMethod = model.Telegram
	// fmt.Println("Selected messaging method:", model.WhatsApp.Int())
	infuraURL := os.Getenv("INFURA_URL")
	contractAddress := os.Getenv("CONTRACT_ADDRESS")
	contractABI := os.Getenv("CONTRACT_ABI")
	fmt.Println("Server is running on port: 8080 ,")
	http.HandleFunc("/webhook/whatsapp", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			verifyWebhook(w, r) // Xác thực Webhook từ Meta
		} else if r.Method == http.MethodPost {
			// Gọi đúng handler được tạo từ ReceiveMessageWhatsapp
			handler.ReceiveMessageWhatsapp(contractAddress, contractABI, infuraURL)(w, r)
		} else {
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		}
	})

	// http.HandleFunc("/webhook/whatsapp", verifyWebhook)
	// http.HandleFunc("/webhook/whatsapp", handler.ReceiveMessageWhatsapp(contractAddress, contractABI, infuraURL))
	http.HandleFunc("/received/message", handler.GetMessageTwilio(contractAddress, contractABI, infuraURL, model.WhatsApp.Int()))
	http.HandleFunc("/received/telegram/message/@OTPCong_bot", handler.HandlerTelegramMessage(contractAddress, contractABI, infuraURL))
	http.HandleFunc("/received/telegram/message/@OTPLocNguyen_2_Bot", handler.HandlerTelegramMessage(contractAddress, contractABI, infuraURL))

	http.ListenAndServe(":8080", nil)
}
