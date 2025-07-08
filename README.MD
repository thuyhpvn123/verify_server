go run main.go

ngrok http 8080

curl -F "url=https://<URL_NGROK_CỦA_BẠN>/received/telegram/message/@<TÊN_BOT_CỦA_BẠN>" "<https://api.telegram.org/bot><TOKEN_BOT_CỦA_BẠN>/setWebhook"
