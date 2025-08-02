package routes

import (
	"context"
	"net/http"

	dw "github.com/haunshila/hypta/applications/api-gateway/internal/handlers/digitalwallet"
	user "github.com/haunshila/hypta/applications/api-gateway/internal/handlers/user"
)

// Pass context to handlers via request context
func RegisterRoutes(ctx context.Context) {
	http.HandleFunc("/digital-wallet/login", func(w http.ResponseWriter, r *http.Request) {
		dw.LoginHandler(w, r.WithContext(ctx))
	})
	http.HandleFunc("/user/profile", func(w http.ResponseWriter, r *http.Request) {
		user.ProfileHandler(w, r.WithContext(ctx))
	})
}
