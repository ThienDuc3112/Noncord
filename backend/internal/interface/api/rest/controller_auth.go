package rest

import (
	"backend/internal/application/interfaces"
	"backend/internal/interface/api/rest/dto/request"
	"backend/internal/interface/api/rest/dto/response"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
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

// register 		godoc
// @Summary 		Register an account
// @Description Register an account
// @Tags 				Auth
// @Accept 			json
// @Produce 		json
// @Param       payload body request.Register true "New account data"
// @Success 		204 {object} nil "No Content"
// @Failure 		400 {object} response.ErrorResponse
// @Failure 		500 {object} response.ErrorResponse
// @Router			/api/v1/auth/register [post]
func (ac *AuthController) RegisterController(w http.ResponseWriter, r *http.Request) {
	_ = request.Register{}
	render.Status(r, http.StatusNotImplemented)
	render.JSON(w, r, response.ErrorResponse{
		Error: "Unimplmented",
	})
}

func (ac *AuthController) LoginController(w http.ResponseWriter, r *http.Request) {
}

func (ac *AuthController) LogoutController(w http.ResponseWriter, r *http.Request) {
}

func (ac *AuthController) RefreshController(w http.ResponseWriter, r *http.Request) {

}
