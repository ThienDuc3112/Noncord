package rest

import (
	"backend/internal/application/command"
	"backend/internal/application/interfaces"
	"backend/internal/interface/rest/dto/response"
	"context"
	"log"
	"net/http"
	"strings"

	"github.com/go-chi/render"
)

func authMiddleware(authService interfaces.AuthService) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			log.Println("[authMiddleware] Verifying user")

			auth := r.Header.Get("Authorization")
			if auth == "" {
				render.Render(w, r, response.ParseErrorResponse("Empty authorization header", 401, nil))
				return
			}
			parts := strings.SplitN(auth, " ", 2)
			if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
				render.Render(w, r, response.ParseErrorResponse("Invalid authorization header format", 401, nil))
				return
			}

			token := parts[1]

			res, err := authService.Authenticate(r.Context(), command.AuthenticateCommand{AccessToken: token})
			if err != nil {
				render.Render(w, r, response.ParseErrorResponse("Internal server error", 500, err))
				return
			}

			ctx := context.WithValue(r.Context(), "user", res.User)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
