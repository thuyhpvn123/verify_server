# Luồng Xử Lý Backend Của Hệ Thống Xác Thực OTP 🚀

## Key Points 📌
- 🔍 Luồng xử lý backend mô tả cách hệ thống nhận, xử lý, và xác minh mã OTP từ người dùng thông qua các kênh như WhatsApp và Telegram, sử dụng Smart Contract và Golang.
- 🔒 Luồng đảm bảo an toàn dữ liệu bằng mã hóa (AES và RSA) và tương tác với blockchain để lưu trữ và xác minh thông tin.

---

## Tổng Quan Về Luồng Xử Lý 🌐

Luồng xử lý backend trong hệ thống xác thực OTP bao gồm các bước từ khi nhận tin nhắn từ người dùng đến khi hoàn tất xác thực và lưu trữ dữ liệu trên Smart Contract. Hệ thống sử dụng Webhook để nhận tin nhắn, dịch vụ backend để xử lý, và blockchain để lưu trữ và xác minh. Dưới đây là các bước chi tiết:

---

## Chi Tiết Luồng Xử Lý 🛠️

### 1. Nhận Tin Nhắn Từ Người Dùng 📩
- **Nguồn**: Tin nhắn chứa mã OTP (One-Time Password) được người dùng gửi qua các kênh như WhatsApp hoặc Telegram.
- **Cách thức**:
  - 👤 Người dùng gửi tin nhắn đến số chatbot đã được cấu hình (ví dụ: số WhatsApp hoặc bot Telegram).
  - 🌐 Tin nhắn này được gửi dưới dạng yêu cầu HTTP POST đến endpoint của backend.
- **Kết quả**: Backend nhận được dữ liệu tin nhắn, bao gồm số điện thoại/username, nội dung OTP, và thông tin kênh (WhatsApp/Telegram).

---

### 2. Phân Tích và Trích Xuất Dữ Liệu 🔎
- **Xử lý**: Webhook backend (viết bằng Golang) phân tích yêu cầu HTTP để trích xuất thông tin quan trọng.
- **Chi tiết**:
  - 📲 Đối với WhatsApp: Parse body JSON để lấy thông tin như `message.From` (số điện thoại) và `message.Text.Body` (mã OTP).
  - 🤖 Đối với Telegram: Parse URL path để xác định bot và decode body JSON để lấy `username`, `chatID`, và `text` (mã OTP).
- **Kết quả**: Dữ liệu OTP và thông tin người gửi được chuẩn bị để kiểm tra.

---

### 3. Kiểm Tra Mã OTP Trên Smart Contract ✅
- **Xử lý**: Dữ liệu được chuyển đến dịch vụ backend để kiểm tra tính hợp lệ của OTP trên Smart Contract.
- **Chi tiết**:
  - 🔗 Backend kết nối với SMC thông qua Infura (sử dụng WebSocket).
  - 📡 Gửi yêu cầu đến Smart Contract để gọi hàm `validateOTP`, cung cấp mã OTP, số điện thoại/username, và ID bot.
  - ✔️ Smart Contract kiểm tra xem mã OTP có khớp với dữ liệu đã lưu trước đó không (dựa trên thời gian, địa chỉ, và loại kênh).
- **Kết quả**:
  - ✅ Nếu OTP hợp lệ, Smart Contract trả về khóa công khai (public key) liên kết với người dùng.
  - ❌ Nếu không hợp lệ, backend ghi log lỗi và có thể thông báo lại cho người dùng (tùy cấu hình).

---

### 4. Mã Hóa và Chuẩn Bị Dữ Liệu Xác Thực 🔐
- **Xử lý**: Nếu OTP hợp lệ, backend sử dụng khóa công khai để mã hóa dữ liệu và chuẩn bị cho bước cuối cùng.
- **Chi tiết**:
  - 📋 Tạo một bản ghi xác thực chứa thông tin như số điện thoại, khóa công khai, và thời gian.
  - 🔑 Mã hóa dữ liệu này bằng thuật toán AES (sử dụng khóa AES ngẫu nhiên) và mã hóa thêm khóa AES bằng RSA với khóa công khai của người dùng.
  - 🛡️ Quá trình này đảm bảo dữ liệu chỉ có thể được giải mã bởi bên có khóa riêng tương ứng.
- **Kết quả**: Dữ liệu đã mã hóa và sẵn sàng gửi lại Smart Contract.

---

### 5. Hoàn Tất Xác Thực và Lưu Trữ Trên Smart Contract 💾
- **Xử lý**: Dữ liệu mã hóa được gửi trở lại Smart Contract để hoàn tất quá trình xác thực.
- **Chi tiết**:
  - ✍️ Backend sử dụng private key của ứng dụng để ký giao dịch và gửi dữ liệu mã hóa đến Smart Contract thông qua hàm `completeAuthentication`.
  - ⚙️ Smart Contract nhận dữ liệu, tính toán hash của dữ liệu này (sử dụng `keccak256`), và lưu trữ hash cùng với thông tin người dùng.
  - ⏳ Giao dịch được gửi lên blockchain và chờ xác nhận.
- **Kết quả**: Hash được lưu trên Smart Contract, cho phép bên thứ ba sau này có thể xác minh tính hợp lệ của dữ liệu.

---

### 6. Cung Cấp Dữ Liệu Cho Bên Thứ Ba (Tùy Chọn) 🤝
- **Xử lý**: Sau khi xác thực thành công, người dùng hoặc hệ thống có thể chia sẻ dữ liệu đã mã hóa với bên thứ ba.
- **Chi tiết**:
  - 📤 Bên thứ ba gửi hash của dữ liệu đã nhận được đến Smart Contract để kiểm tra thông qua hàm `verifyAuthenticationHash`.
  - 🔍 Smart Contract so sánh hash này với hash đã lưu để xác nhận tính chính xác.
- **Kết quả**: Bên thứ ba nhận được kết quả xác minh (hợp lệ hoặc không hợp lệ).

---

## Sơ Đồ Luồng Xử Lý 📊
```plaintext
Người Dùng 👤 → [Tin Nhắn OTP] 📩 → Backend (Webhook) 🌐
           ↓
Phân Tích Dữ Liệu 🔎 → Trích Xuất OTP và Thông Tin 📋
           ↓
Kiểm Tra OTP trên Smart Contract ✅
           ↓
[OTP Hợp Lệ] ✔️ → Mã Hóa Dữ Liệu (AES + RSA) 🔐
           ↓
Gửi Dữ Liệu Mã Hóa đến Smart Contract 📡
           ↓
Lưu Hash và Hoàn Tất Xác Thực 💾
           ↓
[Cung Cấp Dữ Liệu] 📤 → Bên Thứ Ba (Xác Minh qua Smart Contract) 🤝
```
---

## Ví Dụ Luồng

- **Bắt đầu**: Người dùng gửi tin nhắn "123456" qua WhatsApp đến số chatbot.
- **Bước 1**: Backend nhận tin nhắn, trích xuất "123456" và số điện thoại của người dùng.
- **Bước 2**: Gửi yêu cầu kiểm tra "123456" lên Smart Contract.
- **Bước 3**: Smart Contract xác nhận OTP hợp lệ và trả về khóa công khai.
- **Bước 4**: Backend mã hóa dữ liệu xác thực (bao gồm số điện thoại và thời gian) bằng khóa công khai.
- **Bước 5**: Gửi dữ liệu mã hóa trở lại Smart Contract, lưu hash.
- **Kết quả**: Hash được lưu trên blockchain, người dùng có thể chia sẻ dữ liệu với bên thứ ba để xác minh.

---

## Lưu Ý Quan Trọng

- **An Toàn Dữ Liệu**: Luồng đảm bảo rằng tất cả dữ liệu nhạy cảm được mã hóa trước khi gửi lên blockchain hoặc chia sẻ với bên thứ ba.
- **Hiệu Suất**: Cần xử lý đồng thời nhiều yêu cầu từ người dùng để tránh tắc nghẽn.
- **Xử Lý Lỗi**: Nếu bất kỳ bước nào thất bại (ví dụ: OTP không hợp lệ, kết nối blockchain gián đoạn), hệ thống cần ghi log và thông báo cho người dùng hoặc quản trị viên.
- **Cấu Hình**: Đảm bảo các biến môi trường như `INFURA_URL`, `CONTRACT_ADDRESS`, và `WHATSAPP_VERIFY_TOKEN` được bảo mật và không lộ ra ngoài.

---
# Bảng Tổng Hợp Luồng Xử Lý Backend Hệ Thống Xác Thực OTP

| **Bước**                              | **Xử Lý**                                                                 | **Chi Tiết**                                                                                                     | **Kết Quả**                                                                                      |
|---------------------------------------|---------------------------------------------------------------------------|------------------------------------------------------------------------------------------------------------------|--------------------------------------------------------------------------------------------------|
| 1. Nhận Tin Nhắn Từ Người Dùng        | Webhook nhận tin nhắn từ WhatsApp/Telegram                                | - Tin nhắn gửi qua HTTP POST đến endpoint <br>- Trích xuất số điện thoại/username, OTP, kênh | Dữ liệu tin nhắn (số điện thoại/username, OTP, kênh) được nhận bởi backend               |
| 2. Phân Tích và Trích Xuất Dữ Liệu    | Backend (Golang) phân tích yêu cầu HTTP                                   | - WhatsApp: Parse JSON (`message.From`, `message.Text.Body`)<br>- Telegram: Parse URL và JSON (`username`, `chatID`, `text`) | Dữ liệu OTP và thông tin người gửi được chuẩn bị để kiểm tra                             |
| 3. Kiểm Tra Mã OTP Trên Smart Contract| Backend gửi yêu cầu kiểm tra OTP đến Smart Contract                       | - Kết nối qua Infura (WebSocket)<br>- Gọi hàm `validateOTP` với OTP, số điện thoại/username, ID bot<br>- SMC kiểm tra tính hợp lệ | - OTP hợp lệ: Trả về public key<br>- OTP không hợp lệ: Ghi log lỗi, thông báo (tùy chọn) |
| 4. Mã Hóa và Chuẩn Bị Dữ Liệu         | Backend mã hóa dữ liệu xác thực nếu OTP hợp lệ                            | - Tạo bản ghi (số điện thoại, public key, thời gian)<br>- Mã hóa bằng AES (khóa ngẫu nhiên) + RSA (public key)<br>- Chỉ giải mã bằng private key | Dữ liệu mã hóa sẵn sàng gửi lại Smart Contract                                           |
| 5. Hoàn Tất Xác Thực và Lưu Trữ       | Gửi dữ liệu mã hóa đến Smart Contract và lưu trữ                          | - Backend ký giao dịch bằng private key<br>- Gọi `completeAuthentication`<br>- SMC tính hash (`keccak256`) và lưu trữ | Hash và thông tin người dùng được lưu trên blockchain                                    |
| 6. Cung Cấp Dữ Liệu Cho Bên Thứ Ba   | Chia sẻ dữ liệu mã hóa và xác minh (tùy chọn)                             | - Bên thứ ba gửi hash đến SMC qua `verifyHash`<br>- SMC so sánh hash với dữ liệu đã lưu                          | Bên thứ ba nhận kết quả xác minh (hợp lệ/không hợp lệ)                                   |
