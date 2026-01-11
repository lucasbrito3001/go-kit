package httpctx

import "net/http"

type Transport struct {
	base http.RoundTripper
}

func NewTransport(base http.RoundTripper) http.RoundTripper {
	return &Transport{base: base}
}

func (t *Transport) RoundTrip(req *http.Request) (*http.Response, error) {
	ctx := req.Context()

	if v := ctx.Value(requestIDKey); v != nil {
		if reqID, ok := v.(string); ok {
			req.Header.Set(HeaderRequestID, reqID)
		}
	}

	if v := ctx.Value(traceIDKey); v != nil {
		if traceID, ok := v.(string); ok {
			req.Header.Set(HeaderTraceID, traceID)
		}
	}

	return t.base.RoundTrip(req)
}
