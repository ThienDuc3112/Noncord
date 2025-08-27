package rest

import (
	"backend/internal/application/command"
	"backend/internal/application/interfaces"
	"backend/internal/domain/entities"
	"backend/internal/interface/api/rest/dto/response"
	"context"
	"errors"
	"net/http"
	"strings"

	"github.com/go-chi/render"
)

func authMiddleware(authService interfaces.AuthService) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			auth := r.Header.Get("Authorization")
			if auth == "" {
				render.Render(w, r, response.NewErrorResponse("Empty authorization header", 401, nil))
				return
			}
			parts := strings.SplitN(auth, " ", 2)
			if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
				render.Render(w, r, response.NewErrorResponse("Invalid authorization header format", 401, nil))
				return
			}

			token := parts[1]

			res, err := authService.Authenticate(r.Context(), command.AuthenticateCommand{AccessToken: token})
			if err != nil {
				var derr *entities.ChatError
				if errors.As(err, &derr) {
					render.Render(w, r, response.NewErrorResponseFromChatError(derr))
				} else {
					render.Render(w, r, response.NewErrorResponse("Internal server error", 500, err))
				}
				return
			}

			ctx := context.WithValue(r.Context(), "user", res)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
