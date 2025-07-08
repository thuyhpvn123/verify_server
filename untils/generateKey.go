package untils

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
)

func GeneratePublicKey() (string, error) {
	//Tạo một cặp khóa RSA với độ dài 2048-bit
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return "", err
	}
	//Trích xuất Public Key từ Private Key
	publicKeyBytes, err := x509.MarshalPKIXPublicKey(&privateKey.PublicKey)
	if err != nil {
		return "", err
	}

	// Encode thành Base64 mà không có BEGIN/END
	publicKeyBase64 := base64.StdEncoding.EncodeToString(publicKeyBytes)
	return publicKeyBase64, nil
}
