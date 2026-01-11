package httpctx

import "net/http"

func NewClient(base *http.Client) *http.Client {
	if base == nil {
		base = http.DefaultClient
	}

	transport := base.Transport
	if transport == nil {
		transport = http.DefaultTransport
	}

	base.Transport = NewTransport(transport)
	return base
}
