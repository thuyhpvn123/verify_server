# Tài Liệu Đăng Ký Bot Telegram  
# Hệ Thống Xác Thực OTP Qua Telegram 🚀

## Thông Tin Đăng Ký Bot Telegram 📋
- **Tên Bot**: `@OTPVerificationBot` (Tên mẫu, có thể thay đổi theo nhu cầu).
- **Mục Đích**: Xác thực mã OTP (One-Time Password) từ người dùng thông qua Telegram, tích hợp với Smart Contract để đảm bảo an toàn và minh bạch.
- **Token API**: Được cấp bởi BotFather sau khi đăng ký (sẽ cập nhật sau khi hoàn tất).
- **Webhook URL**: [URL của backend để nhận tin nhắn từ Telegram, ví dụ: `https://your-backend.com/telegram/webhook`].

---

## Tổng Quan Hệ Thống 🌐
Hệ thống xác thực OTP qua Telegram được thiết kế để nhận tin nhắn chứa mã OTP từ người dùng, xử lý dữ liệu bằng backend viết bằng Golang, và tương tác với Smart Contract trên blockchain để xác minh và lưu trữ thông tin. Bot Telegram đóng vai trò là giao diện để người dùng gửi mã OTP, đảm bảo trải nghiệm đơn giản và an toàn.



## Mục Đích 📌
Tạo và đăng ký một bot Telegram thông qua `@BotFather`, để dùng API Token để sử dụng bot nhận tin nhắn OTP.

---

## Các Bước Đăng Ký Bot Telegram 🛠️

### 1. Truy Cập BotFather
- **Hành động**: Mở ứng dụng Telegram trên điện thoại hoặc máy tính.
- **Cách thực hiện**:
  - Tìm kiếm `@BotFather` trong thanh tìm kiếm của Telegram.
  - Nhấn vào `@BotFather` để bắt đầu trò chuyện.
- **Kết quả**: Bạn sẽ thấy giao diện trò chuyện với BotFather, kèm theo thông báo chào mừng và danh sách lệnh khả dụng.

---

### 2. Tạo Một Bot Mới
- **Hành động**: Gửi lệnh để yêu cầu tạo bot mới.
- **Cách thực hiện**:
  - Gõ lệnh `/newbot` và nhấn Enter.
  - BotFather sẽ yêu cầu bạn đặt tên cho bot (ví dụ: "MyAwesomeBot").
  - Nhập tên bot (tên này sẽ hiển thị cho người dùng, không cần ký tự đặc biệt).
- **Kết quả**: BotFather sẽ yêu cầu bạn đặt username cho bot.

---

### 3. Đặt Username Cho Bot
- **Hành động**: Cung cấp username độc nhất cho bot.
- **Cách thực hiện**:
  - Username phải kết thúc bằng từ `Bot` (ví dụ: `@MyAwesome_bot`).
  - Nhập username mong muốn (phải là duy nhất, không trùng với bot khác).
  - Nếu username đã được sử dụng, BotFather sẽ thông báo để bạn thử lại.
- **Kết quả**: 
  - Nếu thành công, BotFather sẽ gửi thông báo xác nhận kèm theo **API Token**.

---

### 4. Nhận API Token
- **Hành động**: Lưu lại API Token do BotFather cung cấp.
- **Chi tiết**:
  - API Token là một chuỗi ký tự dài (ví dụ: `123456:ABC-DEF1234567890`).
  - Token này dùng để xác thực và điều khiển bot thông qua Telegram API.
- **Kết quả**: Bạn đã có API Token để sử dụng bot.

---

### 5. Kiểm Tra Bot
- **Hành động**: Kiểm tra bot đã hoạt động chưa.
- **Cách thực hiện**:
  - Tìm kiếm bot bằng username (ví dụ: `@MyAwesomeBot`) trên Telegram.
  - Nhấn "Start" để kiểm tra phản hồi mặc định.
- **Kết quả**: Nếu bot hiển thị và phản hồi, quá trình đăng ký đã hoàn tất.

---

## Ví Dụ Quy Trình
1. Gửi `/newbot` đến `@BotFather`.
2. Nhập tên: `MyAwesomeBot`.
3. Nhập username: `@MyAwesome_bot`.
4. BotFather trả về: API Token

### Hướng Dẫn Tạo Bot Telegram với BotFather

| Bước | Hành Động | Cách Thực Hiện | Kết Quả |
|------|-----------|----------------|---------|
| 1 | Truy cập BotFather | Mở Telegram, tìm @BotFather | Giao diện BotFather hiện ra |
| 2 | Tạo bot mới | Gửi lệnh /newbot | Nhập tên bot (VD: "MyAwesomeBot") |
| 3 | Đặt username | Cung cấp username độc nhất | Username kết thúc bằng Bot (VD: @MyAwesomeBot) |
| 4 | Nhận API Token | Lưu API Token | Token dạng 123456:ABC-DEF1234567890 |
| 5 | Kiểm tra bot | Tìm bot và nhấn "Start" | Bot phản hồi, đăng ký hoàn tất |

