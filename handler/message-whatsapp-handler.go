package handler

import (
	"WhatsappVerifyOTP/model"
	"WhatsappVerifyOTP/service"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func ReceiveMessageWhatsapp(contractAddress string, contractABI string, INFURAL_URL string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Đọc dữ liệu JSON từ request body
		body, err := io.ReadAll(r.Body)
		if err != nil {
			log.Println("❌ Lỗi đọc request body:", err)
			http.Error(w, "Bad Request", http.StatusBadRequest)
			return
		}
		defer r.Body.Close()

		var req model.WebhookRequest
		if err := json.Unmarshal(body, &req); err != nil {
			log.Println("❌ Lỗi parse JSON:", err)
			http.Error(w, "Bad Request", http.StatusBadRequest)
			return
		}

		// Kiểm tra tin nhắn
		if req.Object == "whatsapp_business_account" {
			for _, entry := range req.Entry {
				for _, change := range entry.Changes {
					for _, message := range change.Value.Messages {
						fmt.Printf("📩 Tin nhắn từ %s: %s\n", message.From, message.Text.Body)
						service.CheckOTP(contractAddress, contractABI, INFURAL_URL, message.From, message.Text.Body, "1")
					}
				}
			}
			w.WriteHeader(http.StatusOK)
		} else {
			w.WriteHeader(http.StatusNotFound)
		}
	}
}
