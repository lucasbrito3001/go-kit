package httpctx

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/lucasbrito3001/go-kit/observability/correlation"
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

		ctx = context.WithValue(ctx, correlation.RequestIDKey, requestID)

		traceID := firstNonEmpty(
			c.GetHeader(HeaderTraceID),
			c.GetHeader(HeaderB3TraceID),
			c.GetHeader(HeaderTraceParent),
		)

		if traceID != "" {
			ctx = context.WithValue(ctx, correlation.TraceIDKey, traceID)
		}

		c.Request = req.WithContext(ctx)

		c.Writer.Header().Set(HeaderRequestID, requestID)
		if traceID != "" {
			c.Writer.Header().Set(HeaderTraceID, traceID)
		}

		c.Next()
	}
}
