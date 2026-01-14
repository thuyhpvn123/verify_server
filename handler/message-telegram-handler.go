package handler

import (
	"context"
	"crypto/ecdsa"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"regexp"
	"strings"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/meta-node-blockchain/verify_server/service"
)

// ============================================================================
// 🔒 DUPLICATE REQUEST PREVENTION
// ============================================================================

var (
	processedUpdates = make(map[int]time.Time)
	updatesMutex     sync.RWMutex
)

func init() {
	go func() {
		ticker := time.NewTicker(5 * time.Minute)
		for range ticker.C {
			updatesMutex.Lock()
			now := time.Now()
			for updateID, timestamp := range processedUpdates {
				if now.Sub(timestamp) > 10*time.Minute {
					delete(processedUpdates, updateID)
				}
			}
			updatesMutex.Unlock()
		}
	}()
}

func isUpdateProcessed(updateID int) bool {
	updatesMutex.RLock()
	defer updatesMutex.RUnlock()
	_, exists := processedUpdates[updateID]
	return exists
}

func markUpdateProcessed(updateID int) {
	updatesMutex.Lock()
	defer updatesMutex.Unlock()
	processedUpdates[updateID] = time.Now()
}

// ============================================================================
// 📱 TELEGRAM TYPES
// ============================================================================

type TelegramUpdate struct {
	UpdateID int `json:"update_id"`
	Message  struct {
		MessageID int `json:"message_id"`
		From      struct {
			ID        int64  `json:"id"`
			FirstName string `json:"first_name"`
			Username  string `json:"username"`
		} `json:"from"`
		Chat struct {
			ID   int64  `json:"id"`
			Type string `json:"type"`
		} `json:"chat"`
		Text string `json:"text"`
		Date int64  `json:"date"`
	} `json:"message"`
}

// ============================================================================
// 🔧 REQUEST CONTEXT
// ============================================================================

type RequestContext struct {
	ID           string
	PrivateKey   *ecdsa.PrivateKey
	ContractAddr common.Address
	RPCURL       string
}

// ============================================================================
// 📱 TELEGRAM WEBHOOK HANDLER
// ============================================================================

func HandlerTelegramMessage(
	privateKey *ecdsa.PrivateKey,
	contractAddress common.Address,
	RPC_HTTP_URL string,
) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		requestID := fmt.Sprintf("TG-%d", time.Now().UnixNano())

		log.Printf("═══════════════════════════════════════")
		log.Printf("📱 [%s] TELEGRAM WEBHOOK RECEIVED", requestID)
		log.Printf("═══════════════════════════════════════")

		// Read body
		bodyBytes, err := io.ReadAll(r.Body)
		if err != nil {
			log.Printf("❌ [%s] Error reading body: %v", requestID, err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		r.Body.Close()

		// Make a copy for goroutine
		bodyCopy := make([]byte, len(bodyBytes))
		copy(bodyCopy, bodyBytes)

		log.Printf("📦 [%s] Body length: %d bytes", requestID, len(bodyCopy))

		// Return response immediately
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"status": "ok"})

		// Create isolated context
		ctx := RequestContext{
			ID:           requestID,
			PrivateKey:   privateKey,
			ContractAddr: contractAddress,
			RPCURL:       RPC_HTTP_URL,
		}

		// Process in goroutine
		go processTelegramUpdate(ctx, bodyCopy)
	}
}

// ============================================================================
// 🔧 OTP UTILITIES
// ============================================================================

func cleanOTP(otp string) string {
	cleaned := strings.TrimSpace(otp)
	cleaned = strings.ReplaceAll(cleaned, " ", "")
	cleaned = strings.ReplaceAll(cleaned, "\t", "")
	cleaned = strings.ReplaceAll(cleaned, "\n", "")
	cleaned = strings.ReplaceAll(cleaned, "\r", "")
	cleaned = strings.ReplaceAll(cleaned, "\u00a0", "")
	cleaned = strings.ReplaceAll(cleaned, "\u200b", "")
	
	re := regexp.MustCompile(`[^\d]`)
	cleaned = re.ReplaceAllString(cleaned, "")
	
	return cleaned
}

func validateOTP(otp string) error {
	if otp == "" {
		return fmt.Errorf("OTP is empty")
	}
	
	matched, _ := regexp.MatchString(`^\d+$`, otp)
	if !matched {
		return fmt.Errorf("OTP contains non-digit characters")
	}
	
	if len(otp) < 4 {
		return fmt.Errorf("OTP too short (%d digits, minimum 4)", len(otp))
	}
	
	if len(otp) > 8 {
		return fmt.Errorf("OTP too long (%d digits, maximum 8)", len(otp))
	}
	
	return nil
}

// ============================================================================
// 📱 TELEGRAM MESSAGE PROCESSOR
// ============================================================================

func processTelegramUpdate(ctx RequestContext, bodyBytes []byte) {
	log.Printf("📦 [%s] Processing update...", ctx.ID)

	// Parse update
	var update TelegramUpdate
	if err := json.Unmarshal(bodyBytes, &update); err != nil {
		log.Printf("❌ [%s] JSON parse error: %v", ctx.ID, err)
		log.Printf("📦 [%s] Raw body: %s", ctx.ID, string(bodyBytes))
		return
	}

	// Check if already processed
	if isUpdateProcessed(update.UpdateID) {
		log.Printf("⚠️ [%s] DUPLICATE DETECTED! UpdateID %d - SKIPPING", ctx.ID, update.UpdateID)
		return
	}

	markUpdateProcessed(update.UpdateID)

	log.Printf("📱 [%s] Message Info:", ctx.ID)
	log.Printf("   UpdateID: %d ✅", update.UpdateID)
	log.Printf("   MessageID: %d", update.Message.MessageID)
	log.Printf("   From: %s (@%s, ID: %d)",
		update.Message.From.FirstName,
		update.Message.From.Username,
		update.Message.From.ID)
	log.Printf("   Text (raw): '%s'", update.Message.Text)

	// Validate message
	if update.Message.Text == "" {
		log.Printf("⚠️ [%s] Empty message - skipping", ctx.ID)
		return
	}

	// Extract and COPY to local variables
	text := strings.TrimSpace(update.Message.Text)
	
	var phoneNumber, otpRaw string

	// Parse formats
	if strings.Contains(text, "-") {
		parts := strings.Split(text, "-")
		if len(parts) == 2 {
			otpRaw = strings.TrimSpace(parts[0])
			phoneNumber = strings.TrimSpace(parts[1])
			log.Printf("ℹ️  [%s] Format detected: OTP-Phone", ctx.ID)
		} else {
			log.Printf("❌ [%s] Invalid hyphen format", ctx.ID)
			return
		}
	} else {
		parts := strings.Fields(text)
		log.Printf("🔍 [%s] Parsed %d field(s): %v", ctx.ID, len(parts), parts)

		if len(parts) == 1 {
			otpRaw = parts[0]
			phoneNumber = update.Message.From.Username
			if phoneNumber == "" {
				phoneNumber = fmt.Sprintf("%d", update.Message.From.ID)
			}
			log.Printf("ℹ️  [%s] Single field mode", ctx.ID)
		} else if len(parts) >= 2 {
			phoneNumber = parts[0]
			otpRaw = parts[1]
			log.Printf("ℹ️  [%s] Format detected: Phone OTP", ctx.ID)
		} else {
			log.Printf("❌ [%s] Invalid format", ctx.ID)
			return
		}
	}

	log.Printf("🔍 [%s] Extracted:", ctx.ID)
	log.Printf("   Phone: '%s'", phoneNumber)
	log.Printf("   OTP (raw): '%s'", otpRaw)

	// Clean and validate OTP
	otp := cleanOTP(otpRaw)
	log.Printf("   OTP (cleaned): '%s' (length: %d)", otp, len(otp))

	if err := validateOTP(otp); err != nil {
		log.Printf("❌ [%s] OTP validation failed: %v", ctx.ID, err)
		return
	}

	log.Printf("✅ [%s] OTP validation passed", ctx.ID)
	log.Printf("🔐 [%s] Calling CheckOTP with botID='telegram'...", ctx.ID)

	// Call CheckOTP
	result, err := service.CheckOTP(
		context.Background(),
		ctx.PrivateKey,
		ctx.ContractAddr,
		ctx.RPCURL,
		phoneNumber,
		otp,
		"telegram", // ← botID
	)

	if err != nil {
		log.Printf("❌ [%s] CheckOTP failed: %v", ctx.ID, err)
		log.Printf("═══════════════════════════════════════")
		return
	}

	log.Printf("✅ [%s] ✅ AUTHENTICATION SUCCESSFUL!", ctx.ID)
	log.Printf("✅ [%s] PublicKey: %s", ctx.ID, result.PublicKey)
	log.Printf("✅ [%s] Wallet: %s", ctx.ID, result.Wallet.Hex())
	log.Printf("═══════════════════════════════════════")
}