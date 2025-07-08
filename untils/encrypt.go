package untils

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"fmt"
	"io"
	"log"
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
func EncryptData(publicKeyPEM string, data []byte) (string, string, error) {
	formatPublicKeyPen := formatPEMKey(publicKeyPEM)
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

	// Mã hóa khóa AES bằng RSA
	encryptedAESKey, err := encryptRSA(formatPublicKeyPen, aesKey)
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
func encryptRSA(publicKeyPEM string, aesKey []byte) ([]byte, error) {

	block, error := pem.Decode([]byte(publicKeyPEM))
	fmt.Println(publicKeyPEM)
	if block == nil {
		log.Fatalf("Không thể decode public key: %v", error)

	}

	pub, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	rsaPub, ok := pub.(*rsa.PublicKey)
	if !ok {
		return nil, fmt.Errorf("Không phải RSA public key")
	}

	return rsa.EncryptOAEP(sha256.New(), rand.Reader, rsaPub, aesKey, nil)
}
