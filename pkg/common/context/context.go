package context

import "github.com/rs/zerolog"

// ReqContext stands for request context

const (
	// RequestID constant
	RequestID = "RequestID"

	// Logger constant
	Logger = "Logger"
)

// ReqContext shared in the application
type ReqContext struct {
	// RequestID is the id of the request
	RequestID string

	Logger *zerolog.Logger
}

// New creates an instance of request context
func New(r string, l *zerolog.Logger) *ReqContext {
	return &ReqContext{RequestID: r, Logger: l}
}
