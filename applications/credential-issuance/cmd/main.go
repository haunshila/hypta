package main

import (
	"context"

	"github.com/haunshila/hypta/applications/util"
)

func main() {

    ctx := context.Background()
	// Create logger using util package
	logger := util.NewLogger()

	// Add logger to context using util package
	ctx = util.WithLogger(ctx, logger)

	// Start server
	logger.Info("credential issuance service started")
}