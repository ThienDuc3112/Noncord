package ws

import (
	"backend/internal/application/command"
	"backend/internal/application/interfaces"
	"backend/internal/interface/dto/response"
	"log/slog"
	"net/http"
	"strings"

	"github.com/go-chi/render"
	"github.com/google/uuid"
)

func authMiddleware(authService interfaces.AuthService, w http.ResponseWriter, r *http.Request) *uuid.UUID {
	slog.Info("Verifying user")

	auth := r.Header.Get("Authorization")
	if auth == "" {
		render.Render(w, r, response.ParseErrorResponse("Empty authorization header", 401, nil))
		return nil
	}
	parts := strings.SplitN(auth, " ", 2)
	if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
		render.Render(w, r, response.ParseErrorResponse("Invalid authorization header format", 401, nil))
		return nil
	}

	token := parts[1]

	res, err := authService.Authenticate(r.Context(), command.AuthenticateCommand{AccessToken: token})
	if err != nil {
		render.Render(w, r, response.ParseErrorResponse("Internal server error", 500, err))
		return nil
	}

	return res.UserId
}
