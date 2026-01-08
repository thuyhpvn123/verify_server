package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	model "github.com/meta-node-blockchain/verify_server/model"
	service "github.com/meta-node-blockchain/verify_server/service"
	"github.com/meta-node-blockchain/meta-node/cmd/client"
	"github.com/ethereum/go-ethereum/common"


)

func ReceiveMessageWhatsapp(fromAddress common.Address ,client *client.Client,contractAddress string, contractABI string, INFURAL_URL string) http.HandlerFunc {
	fmt.Println("aaaaaaaaaa")
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("📱 ========================================")
        log.Println("📱 WHATSAPP WEBHOOK RECEIVED")
        log.Println("📱 ========================================")
        log.Printf("📱 Method: %s", r.Method)
        log.Printf("📱 URL: %s", r.URL.String())
        
        // Log ALL headers
        log.Println("📋 Headers:")
        for name, values := range r.Header {
            for _, value := range values {
                log.Printf("   %s: %s", name, value)
            }
        }
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

						// convert phone number
						phoneNumberFromWhatsapp := message.From

						if len(phoneNumberFromWhatsapp) > 2 && phoneNumberFromWhatsapp[:2] == "84" {
							phoneNumberFromWhatsapp = "0" + phoneNumberFromWhatsapp[2:]
							// fmt.Printf("Converted phone number:: %s\n", phoneNumberFromWhatsapp)
						}

						service.CheckOTP(fromAddress,client,contractAddress, contractABI, INFURAL_URL, phoneNumberFromWhatsapp, message.Text.Body, "1")
					}
				}
			}
			w.WriteHeader(http.StatusOK)
		} else {
			w.WriteHeader(http.StatusNotFound)
		}
	}
}
