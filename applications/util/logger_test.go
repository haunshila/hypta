package util

import (
	"context"
	"log/slog"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoggerFromContext(t *testing.T) {
	t.Run("LoggerExists", func(t *testing.T) {
		// Create a logger and add it to the context
		logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
		ctx := WithLogger(context.Background(), logger)

		// Extract logger from context
		result := LoggerFromContext(ctx)

		// Assert logger is not nil and matches the input logger
		assert.NotNil(t, result, "Expected logger to be found in context")
		assert.Equal(t, logger, result, "Expected logger to match the one added to context")
	})

	t.Run("LoggerNotInContext", func(t *testing.T) {
		// Create a context without a logger
		ctx := context.Background()

		// Extract logger from context
		result := LoggerFromContext(ctx)

		// Assert logger is nil
		assert.Nil(t, result, "Expected nil logger when no logger is in context")
	})

	t.Run("WrongTypeInContext", func(t *testing.T) {
		// Add a non-logger value with the same key
		ctx := context.WithValue(context.Background(), logKey("logger"), "not-a-logger")

		// Extract logger from context
		result := LoggerFromContext(ctx)

		// Assert logger is nil
		assert.Nil(t, result, "Expected nil logger when context value is not a logger")
	})
}

func TestNewLogger(t *testing.T) {
	t.Run("CreatesValidLogger", func(t *testing.T) {
		// Create a new logger
		logger := NewLogger()

		// Assert logger is not nil
		assert.NotNil(t, logger, "Expected NewLogger to return a non-nil logger")

		// Verify logger uses JSONHandler and writes to os.Stdout
		handler := logger.Handler()
		jsonHandler, ok := handler.(*slog.JSONHandler)
		assert.True(t, ok, "Expected logger to use JSONHandler")
		assert.NotNil(t, jsonHandler, "Expected JSONHandler to be non-nil")
	})
}

func TestWithLogger(t *testing.T) {
	t.Run("AddsLoggerToContext", func(t *testing.T) {
		// Create a logger and context
		logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
		ctx := WithLogger(context.Background(), logger)

		// Extract logger from context
		result := ctx.Value(logKey("logger"))

		// Assert logger is stored correctly
		assert.NotNil(t, result, "Expected logger to be stored in context")
		assert.Equal(t, logger, result, "Expected stored logger to match input logger")
	})

	t.Run("MultipleLoggers", func(t *testing.T) {
		// Create two different loggers
		logger1 := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
		logger2 := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))

		// Add first logger to context
		ctx := WithLogger(context.Background(), logger1)
		assert.Equal(t, logger1, LoggerFromContext(ctx), "Expected logger1 to be in context")

		// Add second logger to a new context
		ctx = WithLogger(ctx, logger2)
		assert.Equal(t, logger2, LoggerFromContext(ctx), "Expected logger2 to replace logger1 in context")
	})
}