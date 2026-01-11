package logger

import (
	"log/slog"
	"os"
)

func Init(serviceName string) {
	level := parseLogLevel()

	baseHandler := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: level,
	})

	handler := &ContextHandler{handler: baseHandler}

	logger := slog.New(handler).With(
		"service", serviceName,
	)

	slog.SetDefault(logger)
}
