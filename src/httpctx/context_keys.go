package httpctx

type (
	requestIDKeyType struct{}
	traceIDKeyType   struct{}
)

var (
	requestIDKey = requestIDKeyType{}
	traceIDKey   = traceIDKeyType{}
)
