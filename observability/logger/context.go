package logger

import (
	"context"
	"log/slog"

	"github.com/lucasbrito3001/go-kit/observability/correlation"
)

func FromContext(ctx context.Context) *slog.Logger {
	if reqID, ok := ctx.Value(correlation.RequestIDKey).(string); ok {
		return slog.With("request_id", reqID)
	}
	return slog.Default()
}
