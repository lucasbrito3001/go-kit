package httpctx

import (
	"net/http"

	"github.com/lucasbrito3001/go-kit/observability/correlation"
)

type Transport struct {
	base http.RoundTripper
}

func NewTransport(base http.RoundTripper) http.RoundTripper {
	return &Transport{base: base}
}

func (t *Transport) RoundTrip(req *http.Request) (*http.Response, error) {
	ctx := req.Context()

	if v := ctx.Value(correlation.RequestIDKey); v != nil {
		if reqID, ok := v.(string); ok {
			req.Header.Set(HeaderRequestID, reqID)
		}
	}

	if v := ctx.Value(correlation.TraceIDKey); v != nil {
		if traceID, ok := v.(string); ok {
			req.Header.Set(HeaderTraceID, traceID)
		}
	}

	return t.base.RoundTrip(req)
}
