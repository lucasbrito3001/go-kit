package correlation

type (
	RequestIDKeyType struct{}
	TraceIDKeyType   struct{}
)

var (
	RequestIDKey = RequestIDKeyType{}
	TraceIDKey   = TraceIDKeyType{}
)
