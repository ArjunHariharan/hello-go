package core

// ErrorType defines the type of errors
type ErrorType string

const (
	// InternalServerError should be used for internal server errors
	InternalServerError = "Internal Server Error"

	// ValidationError should be used for all validation errors
	ValidationError = "Validation Error"
)

// ErrorCode defines the error codes sent to the client
type ErrorCode string

const (
	// InvalidRequestCode is used to indicate error in dto validation
	InvalidRequestCode = "INVALID_REQUEST"

	// InteralServerCode is used to indicate internal server errors
	InteralServerCode = "INTERNAL_SERVER_ERROR"
)

// Error defines a custom error type
type Error struct {
	Type    ErrorType
	Code    ErrorCode
	Message string
}

func (e *Error) Error() string {
	return e.Message
}

// NewError constructs an error object
func NewError(t ErrorType, c ErrorCode, m string) error {
	return &Error{Type: t, Code: c, Message: m}
}
