package response

import (
	"backend/internal/domain/entities"
	"errors"
	"log"
	"net/http"

	"github.com/go-chi/render"
)

type ErrorResponse struct {
	Error      string `json:"error"`
	statusCode int32
	err        error
}

func (e *ErrorResponse) Render(w http.ResponseWriter, r *http.Request) error {
	log.Printf("[ErrorResponse] Error response: %v\nError: %v\n", e.Error, e.err)
	render.Status(r, int(e.statusCode))
	return nil
}

func NewErrorResponse(str string, statusCode int32, err error) *ErrorResponse {
	derr := &entities.ChatError{}
	if errors.As(err, &derr) {
		return NewErrorResponseFromChatError(derr)
	}

	return &ErrorResponse{
		Error:      str,
		statusCode: statusCode,
		err:        err,
	}
}

func NewErrorResponseFromChatError(err *entities.ChatError) *ErrorResponse {
	if err == nil {
		return nil
	}

	message := err.Message
	statusCode := 500
	switch err.Code {
	case entities.ErrCodeValidationError:
		statusCode = 422
	case entities.ErrCodeNoObject:
		statusCode = 404
	case entities.ErrCodeForbidden:
		statusCode = 403
	case entities.ErrCodeUnauth:
		statusCode = 401
	default:
		message = "Internal server error"
	}

	return &ErrorResponse{
		Error:      message,
		statusCode: int32(statusCode),
		err:        err,
	}

}
