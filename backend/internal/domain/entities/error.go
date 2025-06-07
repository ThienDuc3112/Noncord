package entities

import "fmt"

type ChatErrorCode string

const (
	ErrCodeValidationError = "validation_error"
)

type ChatError struct {
	Code    ChatErrorCode
	Message string
	Err     error
}

func (e *ChatError) Error() string {
	if e.Err != nil {
		return fmt.Sprintf("%s: %s: %v", e.Code, e.Message, e.Err)
	}
	return fmt.Sprintf("%s: %s", e.Code, e.Message)
}

func (e *ChatError) Unwrap() error {
	return e.Err
}

func NewError(code ChatErrorCode, message string, err error) *ChatError {
	return &ChatError{
		Code:    code,
		Message: message,
		Err:     err,
	}
}

var _ error = &ChatError{}
