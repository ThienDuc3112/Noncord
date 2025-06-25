package entities

import (
	"errors"
	"fmt"
)

type ChatErrorCode string

const (
	ErrCodeDepFail          = "service_fail" // Dependency failure (db, 3rd party packages)
	ErrCodeInvalidAssertion = "assert_fail"  // Assertion that should've never been false is false

	ErrCodeValidationError = "validation_error"      // Validation error
	ErrCodeNoObject        = "object_none_existance" // Specified resource don't exist
	ErrCodeInvalidAction   = "invalid_action"        // Action prohibited by the service
	ErrCodeLogicFailure    = "logic_fail"            // Explicitly defined failure case (invalid password, unauthorized access, etc...)
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

// Get underlying ChatError or if not exist, wrap the err and return it
func GetErrOrDefault(err error, code ChatErrorCode, message string) *ChatError {
	if err == nil {
		return nil
	}
	var derr *ChatError
	if errors.As(err, &derr) {
		return derr
	} else {
		return NewError(code, message, err)
	}
}

var _ error = &ChatError{}
