package handler

import (
	"net/http"
)

func RegisterSmartContract(contractAddress string, contractABI string, INFURAL_URL string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
			return
		}

	}
}
