package httpctx

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func firstNonEmpty(values ...string) string {
	for _, v := range values {
		if v != "" {
			return v
		}
	}
	return ""
}

func GinContextMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		req := c.Request
		ctx := req.Context()

		requestID := c.GetHeader(HeaderRequestID)
		if requestID == "" {
			requestID = uuid.NewString()
		}

		ctx = context.WithValue(ctx, requestIDKey, requestID)

		traceID := firstNonEmpty(
			c.GetHeader(HeaderTraceID),
			c.GetHeader(HeaderB3TraceID),
			c.GetHeader(HeaderTraceParent),
		)

		if traceID != "" {
			ctx = context.WithValue(ctx, traceIDKey, traceID)
		}

		c.Request = req.WithContext(ctx)

		c.Writer.Header().Set(HeaderRequestID, requestID)
		if traceID != "" {
			c.Writer.Header().Set(HeaderTraceID, traceID)
		}

		c.Next()
	}
}
