package rest

import (
	"backend/internal/application/interfaces"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type AuthController struct {
	authService interfaces.AuthService
}

func NewAuthController(service interfaces.AuthService) *AuthController {
	return &AuthController{authService: service}
}

func (ac *AuthController) RegisterRoute(r chi.Router) {
	r.Route("/auth", func(r chi.Router) {
		r.Post("/register", ac.RegisterController)
		r.Post("/login", ac.LoginController)
		r.Post("/logout", ac.LogoutController)
	})
}

func (ac *AuthController) RegisterController(w http.ResponseWriter, r *http.Request) {
}

func (ac *AuthController) LoginController(w http.ResponseWriter, r *http.Request) {
}

func (ac *AuthController) LogoutController(w http.ResponseWriter, r *http.Request) {
}
