package util

import (
	"context"
	"log/slog"
	"os"
)
type logKey string

// LoggerFromContext extracts slog.Logger from context, returns nil if not found.
func LoggerFromContext(ctx context.Context) *slog.Logger {
	logger, ok := ctx.Value(logKey("logger")).(*slog.Logger)
	if !ok {
		return nil
	}
	return logger
}

// NewLogger returns a new slog.Logger with JSON handler.
func NewLogger() *slog.Logger {
	return slog.New(slog.NewJSONHandler(os.Stdout, nil))
}

// WithLogger adds a slog.Logger to the context.
func WithLogger(ctx context.Context, logger *slog.Logger) context.Context {
	return context.WithValue(ctx, logKey("logger"), logger)
}
