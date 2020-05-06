package context

// ReqContext stands for request context
type ReqContext struct {
	// RequestID is the id of the request
	RequestID string
}

// New creates an instance of request context
func New(r string) *ReqContext {
	return &ReqContext{RequestID: r}
}
