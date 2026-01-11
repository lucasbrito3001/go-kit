package logger

import (
	"context"
	"log/slog"
)

type ContextHandler struct {
	slog.Handler
}

func (h ContextHandler) Handle(ctx context.Context, r slog.Record) error {
	if reqID, ok := ctx.Value(requestIDKey).(string); ok {
		r.Add("request_id", reqID)
	}
	return h.Handler.Handle(ctx, r)
}
