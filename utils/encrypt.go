package utils

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"fmt"
	"io"
	"strings"
)

func formatPEMKey(key string) string {
	header := "-----BEGIN PUBLIC KEY-----\n"
	footer := "\n-----END PUBLIC KEY-----"

	// Đảm bảo key không có header/footer cũ (tránh trùng lặp)
	key = strings.ReplaceAll(key, "-----BEGIN PUBLIC KEY-----", "")
	key = strings.ReplaceAll(key, "-----END PUBLIC KEY-----", "")

	// Chuẩn hóa: Thêm xuống dòng mỗi 64 ký tự để đảm bảo đúng format PEM
	var formattedKey strings.Builder
	for i := 0; i < len(key); i += 64 {
		end := i + 64
		if end > len(key) {
			end = len(key)
		}
		formattedKey.WriteString(key[i:end] + "\n")
	}

	return header + formattedKey.String() + footer
}

// EncryptData mã hóa dữ liệu bằng AES và mã hóa khóa AES bằng RSA
func EncryptData(publicKeyHex string, data []byte) (string, string, error) {
	// Tạo khóa AES ngẫu nhiên 32 byte (AES-256)
	aesKey := make([]byte, 32)
	if _, err := rand.Read(aesKey); err != nil {
		return "", "", err
	}

	// Mã hóa dữ liệu bằng AES
	encryptedData, _, err := encryptAES(data, aesKey)
	if err != nil {
		return "", "", err
	}

	// Mã hóa khóa AES bằng RSA, truyền vào chuỗi hex public key
	encryptedAESKey, err := encryptRSA(publicKeyHex, aesKey)
	if err != nil {
		return "", "", err
	}

	// Trả về dữ liệu mã hóa và khóa AES mã hóa dưới dạng base64
	return base64.StdEncoding.EncodeToString(encryptedData), base64.StdEncoding.EncodeToString(encryptedAESKey), nil
}

func encryptAES(plainText []byte, key []byte) ([]byte, []byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, nil, err
	}

	ciphertext := make([]byte, aes.BlockSize+len(plainText))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return nil, nil, err
	}

	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(ciphertext[aes.BlockSize:], plainText)

	return ciphertext, iv, nil
}

// Sửa lại hàm encryptRSA để xử lý chuỗi hex
func encryptRSA(publicKeyBase64 string, aesKey []byte) ([]byte, error) {
	// 1. Giải mã chuỗi public key từ Base64 về dạng byte
	derBytes, err := base64.StdEncoding.DecodeString(publicKeyBase64)
	if err != nil {
		return nil, fmt.Errorf("không thể giải mã base64 public key: %v", err)
	}

	// 2. Parse các byte (định dạng DER) để lấy public key
	pub, err := x509.ParsePKIXPublicKey(derBytes)
	if err != nil {
		return nil, fmt.Errorf("không thể parse PKIX public key: %v", err)
	}

	// 3. Chuyển đổi sang dạng RSA public key
	rsaPub, ok := pub.(*rsa.PublicKey)
	if !ok {
		return nil, fmt.Errorf("key không phải là RSA public key")
	}

	// 4. Mã hóa khóa AES bằng RSA public key
	return rsa.EncryptOAEP(sha256.New(), rand.Reader, rsaPub, aesKey, nil)
}
