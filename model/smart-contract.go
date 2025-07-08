package model

type VerifyInfo struct {
	PhoneNumber string `json:"phoneNumber"` // Số điện thoại
	PublicKey   string `json:"publicKey"`   // Khóa công khai
}
