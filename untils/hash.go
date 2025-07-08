package untils

import (
	"fmt"
	"math/big"
	"strings"

	"golang.org/x/crypto/sha3"
)

func hashData(publicKey, userAddress, otp, phoneNumber string) string {
	// Ghép tất cả dữ liệu thành một chuỗi duy nhất
	data := strings.Join([]string{publicKey, userAddress, otp, phoneNumber}, "|")

	// Tạo Keccak-256 hash
	hash := sha3.NewLegacyKeccak256()
	hash.Write([]byte(data))
	hashed := hash.Sum(nil)

	// Chuyển thành hex string
	return fmt.Sprintf("%x", hashed)
}
func StringToUint256(otp string) (*big.Int, error) {
	num := new(big.Int)
	num, ok := num.SetString(otp, 10) // Base 10 conversion
	if !ok {
		return nil, fmt.Errorf("invalid number format")
	}
	return num, nil
}
