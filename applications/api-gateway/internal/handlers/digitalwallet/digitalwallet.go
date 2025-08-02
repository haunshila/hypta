package digitalwallet

import (
	"encoding/json"
	"net/http"

	"github.com/haunshila/hypta/applications/api-gateway/internal/services/digitalwallet"
	"github.com/haunshila/hypta/applications/util"
)

// LoginRequest represents the expected login payload
type LoginRequest struct {
	UserID   string `json:"user_id"`
	Password string `json:"password"`
}

// LoginHandler handles digital wallet login requests using the service
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	// Get logger from context
	logger := util.LoggerFromContext(r.Context())

	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		if logger != nil {
			logger.Warn("Invalid method", "method", r.Method)
		}
		return
	}

	var req LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		if logger != nil {
			logger.Error("Invalid request payload", "error", err)
		}
		return
	}

	service := digitalwallet.DigitalWalletService{}
	if service.Authenticate(req.UserID, req.Password) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message":"Login successful"}`))
		if logger != nil {
			logger.Info("Login successful", "user_id", req.UserID)
		}
	} else {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		if logger != nil {
			logger.Warn("Invalid credentials", "user_id", req.UserID)
		}
	}
}
