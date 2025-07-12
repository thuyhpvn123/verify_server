# Hướng Dẫn Triển Khai Bot WhatsApp API

Hướng dẫn cách đăng ký, cấu hình và triển khai bot sử dụng WhatsApp Business API với Webhook.

## 1. Giới thiệu

WhatsApp Business API cho phép doanh nghiệp giao tiếp với khách hàng qua WhatsApp một cách tự động và hiệu quả. có thể sử dụng API để gửi tin nhắn, nhận phản hồi và quản lý hội thoại từ khách hàng thông qua webhook.

## 2. Yêu cầu hệ thống

Trước khi bắt đầu, cần chuẩn bị:

- Tài khoản Facebook Developer.
- Một Facebook Business profile
- Một ứng dụng Facebook được liên kết với WhatsApp Business API.
- Máy chủ để lắng nghe các webhook từ WhatsApp (có thể là máy chủ bất kỳ có thể kết nối internet).
- Token truy cập API từ Facebook Developer.

## 3. Chuẩn bị

Để bắt đầu, cần tạo tài khoản Facebook Developer, Facebook business profile và một ứng dụng WhatsApp:

### Lưu ý
Bạn cần một tài khoản Facebook chưa vi phạm các chính sách của Facebook để  có thể tạo được Facebook Business profile truy cập đường link này để check tài khoản facebook của bạn có bị vi phạm hay ko [https://www.facebook.com/accountquality/](https://www.facebook.com/accountquality/).
Nếu tài khoản của bạn bị thông báo đã vi phạm các chính sách của Facebook thì hãy sử dụng một tài khoản facebook khác để đăng kí Facebook business profile.

1. Truy cập: [https://business.facebook.com/business/loginpage/?login_options[0]=FB&login_options[1]=IG&login_options[2]=SSO&config_ref=biz_login_tool_flavor_mbs&create_business_portfolio_for_bm=1](https://business.facebook.com/business/loginpage/?login_options[0]=FB&login_options[1]=IG&login_options[2]=SSO&config_ref=biz_login_tool_flavor_mbs&create_business_portfolio_for_bm=1)
2. Tạo một Facebook Business profile bằng tài khoản Facebook hiện có của bạn

3. Truy cập: [https://developers.facebook.com/apps/create/](https://deGoelopers.facebook.com/apps/create/).
4. Tạo một ứng dụng mới và chọn loại **Business**.
5. Sau khi tạo, sẽ có **App ID** và **App Secret**. Thông tin này sẽ được sử dụng để cấu hình kết nối với WhatsApp API.

## 4. Đăng ký WhatsApp Business API

1. Vào trang quản lý ứng dụng của Facebook Developer.
2. Truy cập mục **WhatsApp > Getting Started**.
3. Kết nối hoặc tạo tài khoản WhatsApp Business của bạn.
4. Cấu hình số điện thoại mà bot sẽ sử dụng để nhận và gửi tin nhắn.
5. sẽ nhận được một **Token**. Token này sẽ được dùng để thực hiện các yêu cầu API.

## 5. Cấu hình webhook

Webhook sẽ cho phép ứng dụng của nhận tin nhắn từ người dùng WhatsApp.

### Bước 1: Thiết lập webhook

1. Truy cập **Webhooks** trên trang quản lý ứng dụng Facebook Developer.
2. Cung cấp URL của server Webhook (ví dụ: `https://your-server.com/webhook`).

3. Xác thực webhook bằng cách cung cấp `verify_token` mà bạn tạo.

### Lưu ý
    verify_token sẽ được dùng làm biến môi trường để  giao tiếp với webhooks
    
4. cài đặt URL end point để lắng nghe
    - Ví dụ bạn chạy server ngrok có địa chỉ như sau `https://d6b1d60aeee4.ngrok-free.app`
    - thì cài đặt là `https://d6b1d60aeee4.ngrok-free.app/webhook/whatsapp`

5. sau khi thiết lập xong cấu hình webhook bạn chọn **WhatSapp Business Account** trong seclect box mục sản phẩm

6. Kéo xuống tìm message bấm đăng ký để có thể gửi nhận tin nhắn từ Whatsapp 

### Bước 2: Xác thực webhook

Khi cấu hình webhook, Facebook sẽ gửi một yêu cầu GET để xác thực. Đảm bảo rằng máy chủ của phản hồi yêu cầu theo định dạng sau:

```bash
GET /webhook?hub.verify_token=your_token&hub.challenge=CHALLENGE_ACCEPTED&hub.mode=subscribe
```

## Lưu ý quan trọng
1. Quyền truy cập: Kiểm tra kỹ quyền truy cập của ứng dụng trên Facebook Developer để đảm bảo nó có đủ quyền để truy cập API WhatsApp.
2. Xử lý lỗi: Hãy đảm bảo rằng hệ thống của bạn xử lý được các trường hợp như token hết hạn, lỗi xác thực webhook, hoặc lỗi kết nối API.
3. Chạy thử: Sau khi hoàn thành cấu hình, chạy thử bot WhatsApp bằng cách gửi tin nhắn từ tài khoản WhatsApp cá nhân đến số điện thoại đã đăng ký.




# Bảng Kết Luận Tổng Quan: Triển Khai Bot WhatsApp API

| **Giai đoạn**              | **Mô tả**                                                                                 | **Hành động chính**                                                                                                  | **Kết quả**                                                                                  |
|----------------------------|-------------------------------------------------------------------------------------------|---------------------------------------------------------------------------------------------------------------------|---------------------------------------------------------------------------------------------|
| **1. Giới thiệu**          | WhatsApp Business API giúp doanh nghiệp giao tiếp tự động với khách hàng qua WhatsApp.   | Hiểu mục đích: gửi tin nhắn, nhận phản hồi, quản lý hội thoại qua Webhook.                                          | Nắm rõ vai trò của API trong giao tiếp doanh nghiệp.                                        |
| **2. Yêu cầu hệ thống**    | Chuẩn bị các thành phần cần thiết trước khi triển khai.                                   | - Tài khoản Facebook Developer.<br>- Ứng dụng Business.<br>- Máy chủ Webhook.<br>- Token API.                       | Đảm bảo đủ điều kiện để bắt đầu đăng ký và triển khai.                                      |
| **3. Chuẩn bị**            | Thiết lập tài khoản và ứng dụng để kết nối WhatsApp API.                                  | - Truy cập [developers.facebook.com](https://developers.facebook.com).<br>- Tạo ứng dụng Business.<br>- Lấy App ID, App Secret. | Có App ID và App Secret để cấu hình API.                                                   |
| **4. Đăng ký WhatsApp API**| Quy trình đăng ký tài khoản WhatsApp Business API.                                       | - Vào WhatsApp > Getting Started.<br>- Kết nối tài khoản Business.<br>- Cấu hình số điện thoại.<br>- Lấy Token.     | Số điện thoại được đăng ký và Token sẵn sàng sử dụng.                                      |
| **5. Cấu hình Webhook**    | Thiết lập Webhook để nhận tin nhắn từ người dùng WhatsApp.                               | - Thiết lập URL Webhook.<br>- Xác thực bằng `verify_token`.<br>- Chọn sự kiện (tin nhắn, trạng thái).               | Webhook hoạt động, sẵn sàng nhận tin nhắn từ WhatsApp.                                     |
| **- Bước 1: Thiết lập**   | Cấu hình URL và sự kiện trên Facebook Developer.                                         | - Cung cấp URL (ví dụ: `https://your-server.com/webhook`).<br>- Chọn sự kiện nhận tin nhắn.                        | Webhook được liên kết với ứng dụng WhatsApp.                                               |
| **- Bước 2: Xác thực**    | Xác minh Webhook bằng yêu cầu GET từ Facebook.                                           | - Máy chủ phản hồi: `GET /webhook?hub.verify_token=your_token&hub.challenge=CHALLENGE_ACCEPTED&hub.mode=subscribe`. | Webhook được xác thực thành công, sẵn sàng hoạt động.                                      |
| **Lưu ý quan trọng**       | Đảm bảo triển khai trơn tru và xử lý lỗi hiệu quả.