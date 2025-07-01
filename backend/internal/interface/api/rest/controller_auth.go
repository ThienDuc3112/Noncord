package rest

import (
	"backend/internal/application/command"
	"backend/internal/application/interfaces"
	"backend/internal/domain/entities"
	"backend/internal/interface/api/rest/dto/request"
	"backend/internal/interface/api/rest/dto/response"
	"errors"
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
		r.Post("/refresh", ac.RefreshController)
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
	body := request.Register{}
	if err := render.Bind(r, &body); err != nil {
		render.Render(w, r, response.NewErrorResponse("Invalid body", http.StatusBadRequest, err))
		return
	}

	_, err := ac.authService.Register(r.Context(), command.RegisterCommand{
		Username: body.Username,
		Email:    body.Email,
		Password: body.Password,
	})
	if err != nil {
		var e *entities.ChatError
		if errors.As(err, &e) {
			switch e.Code {
			case entities.ErrCodeValidationError, entities.ErrCodeNoObject:
				render.Render(w, r, response.NewErrorResponse(e.Message, http.StatusBadRequest, e))
			case entities.ErrCodeForbidden:
				render.Render(w, r, response.NewErrorResponse(e.Message, http.StatusForbidden, e))
			case entities.ErrCodeUnauth:
				render.Render(w, r, response.NewErrorResponse(e.Message, http.StatusUnauthorized, e))
			default:
				render.Render(w, r, response.NewErrorResponse("Internal server error", http.StatusInternalServerError, e))
			}
			return
		} else {
			render.Render(w, r, response.NewErrorResponse("Internal server error", http.StatusInternalServerError, err))
		}

		return
	}

	render.NoContent(w, r)
}

// register     godoc
// @Summary     Login
// @Description Logging in an account without sso
// @Tags        Auth
// @Accept      json
// @Produce     json
// @Param       payload body request.Login true "New account data"
// @Success 		200 {object} response.LoginResponse "Access token"
// @Header      200 {string} Cookie "refreshToken=abcd1234; HttpOnly; Path=/api/v1/auth/refresh"
// @Failure     400 {object} response.ErrorResponse "Missing field"
// @Failure     401 {object} response.ErrorResponse "Wrong credential"
// @Failure     403 {object} response.ErrorResponse "SSO enabled account"
// @Failure     500 {object} response.ErrorResponse "Internal server error"
// @Router      /api/v1/auth/login [post]
func (ac *AuthController) LoginController(w http.ResponseWriter, r *http.Request) {
	body := request.Login{}
	if err := render.Bind(r, &body); err != nil {
		render.Render(w, r, response.NewErrorResponse("Invalid body", http.StatusBadRequest, err))
		return
	}

	tokens, err := ac.authService.Login(r.Context(), command.LoginCommand{
		Username:  body.Username,
		Password:  body.Password,
		UserAgent: r.UserAgent(),
	})

	if err != nil {
		var e *entities.ChatError
		if errors.As(err, &e) {
			switch e.Code {
			case entities.ErrCodeValidationError, entities.ErrCodeNoObject:
				render.Render(w, r, response.NewErrorResponse(e.Message, http.StatusBadRequest, e))
			case entities.ErrCodeForbidden:
				render.Render(w, r, response.NewErrorResponse(e.Message, http.StatusForbidden, e))
			case entities.ErrCodeUnauth:
				render.Render(w, r, response.NewErrorResponse(e.Message, http.StatusUnauthorized, e))
			default:
				render.Render(w, r, response.NewErrorResponse("Internal server error", http.StatusInternalServerError, e))
			}
			return
		} else {
			render.Render(w, r, response.NewErrorResponse("Internal server error", http.StatusInternalServerError, err))
		}

		return
	}

	res := response.LoginResponse{
		AccessToken:  tokens.AccessToken,
		RefreshToken: tokens.RefreshToken,
	}

	render.Status(r, http.StatusOK)
	render.JSON(w, r, res)
}

// register     godoc
// @Summary     Logout
// @Description Invalidate the current session
// @Tags        Auth
// @Produce     json
// @Param       Cookie header string true "refreshToken=\<Refresh token here\>"
// @Success 		204 {object} nil "No Content"
// @Header      204 {string} Cookie "refreshToken=; HttpOnly; Path=/api/v1/auth/refresh"
// @Failure     401 {object} response.ErrorResponse "Unknown session"
// @Failure     500 {object} response.ErrorResponse "Internal server error"
// @Router      /api/v1/auth/logout [post]
func (ac *AuthController) LogoutController(w http.ResponseWriter, r *http.Request) {
	render.Render(w, r, response.NewErrorResponse("Invalid body", http.StatusNotImplemented, nil))
}

// register     godoc
// @Summary     Refresh
// @Description Rotate current refresh token
// @Tags        Auth
// @Produce     json
// @Param       Cookie header string true "refreshToken=\<Refresh token here\>"
// @Success 		204 {object} nil "No Content"
// @Header      204 {string} Cookie "refreshToken=abcd1234; HttpOnly; Path=/api/v1/auth/refresh"
// @Failure     401 {object} response.ErrorResponse "Unknown session"
// @Failure     500 {object} response.ErrorResponse "Internal server error"
// @Router      /api/v1/auth/refresh [post]
func (ac *AuthController) RefreshController(w http.ResponseWriter, r *http.Request) {
	render.Render(w, r, response.NewErrorResponse("Invalid body", http.StatusNotImplemented, nil))
}
