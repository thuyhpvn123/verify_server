package handler

import (
	"net/http"
)

func GetMessageTwilio(contractAddress string, contractABI string, INFURAL_URL string, messageType int) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// if r.Method != http.MethodPost {
		// 	http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		// 	return
		// }
		// msg := model.MessageTwilio{
		// 	SmsSid:      r.FormValue("SmsSid"),
		// 	ProfileName: r.FormValue("ProfileName"),
		// 	WaId:        r.FormValue("WaId"),
		// 	SmsStatus:   r.FormValue("SmsStatus"),
		// 	Body:        r.FormValue("Body"),
		// 	To:          r.FormValue("To"),
		// 	From:        r.FormValue("From"),
		// }

		// log.Printf("Received message from %s: %s", strings.TrimPrefix(msg.From, "whatsapp:+"), msg.Body)

		// w.Header().Set("Content-Type", "application/json")
		// response := map[string]string{"message": "Received"}
		// json.NewEncoder(w).Encode(response)
		// service.CheckOTP(contractAddress, contractABI, INFURAL_URL, strings.TrimPrefix(msg.From, "whatsapp:"), msg.Body, "1")
	}
}
