package response

import (
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
	log.Printf("Error response: %v\nError: %v\n", e.Error, e.err)
	render.Status(r, int(e.statusCode))
	return nil
}

func NewErrorResponse(str string, statusCode int32, err error) *ErrorResponse {
	return &ErrorResponse{
		Error:      str,
		statusCode: statusCode,
		err:        err,
	}
}
