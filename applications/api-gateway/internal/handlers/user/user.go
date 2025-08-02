package user

import (
	"encoding/json"
	"net/http"

	"github.com/haunshila/hypta/applications/api-gateway/internal/services/user"
	"github.com/haunshila/hypta/applications/util"
)

// ProfileRequest represents the expected profile request payload
type ProfileRequest struct {
	UserID string `json:"user_id"`
}

// ProfileHandler handles user profile requests using the service
func ProfileHandler(w http.ResponseWriter, r *http.Request) {
	// Get logger from context
	logger := util.LoggerFromContext(r.Context())

	if r.Method != http.MethodPost {
		logger.Debug("Method not allowed")
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req ProfileRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	service := user.UserService{}
	profile := service.GetProfile(req.UserID)
	if profile == nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(profile)
}
