package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	model "github.com/meta-node-blockchain/verify_server/model"
	service "github.com/meta-node-blockchain/verify_server/service"
	"github.com/meta-node-blockchain/meta-node/cmd/client"
	"github.com/ethereum/go-ethereum/common"

)

func HandlerTelegramMessage(fromAddress common.Address ,client *client.Client,contractAddress string, contractABI string, INFURAL_URL string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		// --- Phần 1: Đọc và phân tích tin nhắn từ Telegram ---
		var telegramMessage model.TelegramMessage
		urlPath := r.URL.Path
		partsUrl := strings.Split(urlPath, "/")
		botUsername := partsUrl[len(partsUrl)-1]

		if err := json.NewDecoder(r.Body).Decode(&telegramMessage); err != nil {
			log.Println("❌ Lỗi decode JSON:", err)
			http.Error(w, "Lỗi parse request", http.StatusBadRequest)
			return
		}

		// --- Phần 2: Lọc tin nhắn cũ và tin nhắn lệnh ---
		messageTimestamp := int64(telegramMessage.Message.Date)
		currentTimestamp := time.Now().Unix()

		// Bỏ qua tin nhắn cũ hơn 60 giây
		if currentTimestamp-messageTimestamp > 60 {
			return
		}

		text := telegramMessage.Message.Text
		// Bỏ qua các lệnh (bắt đầu bằng "/")
		if strings.HasPrefix(text, "/") {
			return
		}

		// --- Phần 3: Tách OTP và Số Điện Thoại từ tin nhắn ---
		parts := strings.Split(text, "-")

		// Kiểm tra xem tin nhắn có đúng định dạng "<OTP>-<Số Điện Thoại>" không
		if len(parts) != 2 {
			log.Printf("⚠️ Sai định dạng tin nhắn từ Chat ID %d. Mong muốn: '<OTP>-<SĐT>', Nhận được: '%s'", telegramMessage.Message.Chat.ID, text)
			fmt.Fprintf(w, "OK")
			return
		}

		// Gán OTP và số điện thoại từ các phần đã tách
		otp := parts[0]
		userPhoneNumber := parts[1]

		// --- Phần 4: Gửi dữ liệu đã tách đến Smart Contract để xác thực ---
		service.CheckOTP(fromAddress,client,contractAddress, contractABI, INFURAL_URL, userPhoneNumber, otp, botUsername)

		fmt.Fprintf(w, "OK")
	}
}
