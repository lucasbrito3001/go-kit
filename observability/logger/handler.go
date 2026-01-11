package logger

import (
	"context"
	"log/slog"

	"github.com/lucasbrito3001/go-kit/observability/correlation"
)

type ContextHandler struct {
	slog.Handler
}

func (h ContextHandler) Handle(ctx context.Context, r slog.Record) error {
	if reqID, ok := ctx.Value(correlation.RequestIDKey).(string); ok {
		r.Add("request_id", reqID)
	}
	return h.Handler.Handle(ctx, r)
}
