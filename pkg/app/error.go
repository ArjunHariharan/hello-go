package app

// ErrorType defines the type of errors
type ErrorType string

const (
	// InternalServerError should be used for internal server errors
	InternalServerError = "Internal Server Error"

	// ValidationError should be used for all validation errors
	ValidationError = "Validation Error"
)

// Error defines a custom error type
type Error struct {
	Type    ErrorType
	Message string
}

func (e *Error) Error() string {
	return e.Message
}

// NewError constructs an error object
func NewError(t ErrorType, m string) error {
	return &Error{Type: t, Message: m}
}
