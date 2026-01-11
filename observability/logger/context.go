package logger

import (
	"context"
	"log/slog"
)

type requestIDKeyType struct{}

var requestIDKey = requestIDKeyType{}

func FromContext(ctx context.Context) *slog.Logger {
	if reqID, ok := ctx.Value(requestIDKey).(string); ok {
		return slog.With("request_id", reqID)
	}
	return slog.Default()
}
