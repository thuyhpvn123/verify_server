go run main.go

ngrok http 8080

https://wa.me/15551850221



sau khi tạo bot thì set bot cho url webhook
curl -s "https://api.telegram.org/bot8267654371:AAFyljaL4bQ55W12x3b5e7WPJQnb6XAqd2s/setWebhook?url=https://microclimatic-darcey-hysterogenic.ngrok-free.dev/received/telegram/message/@thuyabcbot&secret_token=ibesecret"
curl -s "https://api.telegram.org/bot8132349832:AAGxEdTlZ45K84KaZbnmxRRglTT-0tusY9M/setWebhook?url=https://microclimatic-darcey-hysterogenic.ngrok-free.dev/received/telegram/message/@thuyabcbot&secret_token=ibesecret"
curl -s "https://api.telegram.org/bot8547508784:AAFKTp7kEKxwKFQ1NT1KxQGcvttl6lR1q8o/setWebhook?url=https://microclimatic-darcey-hysterogenic.ngrok-free.dev/received/telegram/message/@thuyabcbot&secret_token=ibesecret"

kiểm tra set được chưa
curl -s "https://api.telegram.org/bot8267654371:AAFyljaL4bQ55W12x3b5e7WPJQnb6XAqd2s/getWebhookInfo"


agent:
  authtoken: 36NSUEF95CXJKqYRRZuJ9mkbdPt_d7WupJ6WVqKgpGF1jASW

url ngrok:
https://microclimatic-darcey-hysterogenic.ngrok-free.dev


url local server 
https://microclimatic-darcey-hysterogenic.ngrok-free.dev/received/telegram/message/@thuyabcbot

nhắn tin cho bot telegram
curl -X POST   -H "Content-Type: application/json"   -d '{"message":{"chat":{"id":12345},"text":"133049-0901328063","date":1736077846}}'   https://microclimatic-darcey-hysterogenic.ngrok-free.dev/received/telegram/message/@thuyabcbot

curl -X POST "https://api.telegram.org/bot8267654371:AAFyljaL4bQ55W12x3b5e7WPJQnb6XAqd2s/sendMessage" \
  -H "Content-Type: application/json" \
  -d '{
        "chat_id": "2137283886",
        "text": 985872-0901328063"
      }'



curl -X POST "https://api.telegram.org/bot8132349832:AAFyljaL4bQ55W12x3b5e7WPJQnb6XAqd2s/sendMessage" \
  -H "Content-Type: application/json" \
  -d '{
        "chat_id": "2137283886",
        "text": 985872-0901328063"
      }'


lay chatID:
curl -s "https://api.telegram.org/bot8267654371:AAFyljaL4bQ55W12x3b5e7WPJQnb6XAqd2s/getUpdates"


Ngrok mo cong: ngrok http 9000
-> xem request qua ngrok: http://127.0.0.1:4040

1. test telegram:
requestAuthentication:
_identifier = "0356065027"
_walletAddress = 0xa1aa9efa7d314994b080a8f377126355eb8d8db1
_publicKey = 82885ad9781e513a34bfb7b7bbd435cc3aedfdacebf03aed80c7fbb8306d81a338401af75e67fe0ffc30b61330819d7d
_typeMethod = 1

2. goi OTP lay otp - 756458
3. nhan vao bot thuyabcbot: 756458-0356065027
