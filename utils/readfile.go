package utils

import (
	"fmt"
	"os"
)

// ReadABIFromFile đọc file JSON ABI và trả về string nguyên gốc (giữ nguyên format)
func ReadABIFromFile(filePath string) (string, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return "", fmt.Errorf("failed to read ABI file: %w", err)
	}

	return string(data), nil
}