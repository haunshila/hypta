package main

import (
	"context"
	"net/http"

	"github.com/haunshila/hypta/applications/api-gateway/internal/routes"
	"github.com/haunshila/hypta/applications/util"
)

func main() {

    ctx := context.Background()
	// Create logger using util package
	logger := util.NewLogger()

	// Add logger to context using util package
	ctx = util.WithLogger(ctx, logger)

	// Register routes with context
	routes.RegisterRoutes(ctx)

	// Start HTTP server
	logger.Info("API Gateway service started", "addr", ":8080")
	http.ListenAndServe(":8080", nil)
}
