package rest

import (
	"backend/internal/application/interfaces"
	"backend/internal/interface/dto/response"
	"log/slog"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/google/uuid"
)

type UserController struct {
	userQueries interfaces.UserQueries
	authService interfaces.AuthService
}

func NewUserController(
	authService interfaces.AuthService,
	userQueries interfaces.UserQueries,
) *UserController {
	return &UserController{userQueries: userQueries, authService: authService}
}

func (c *UserController) RegisterRoute(r chi.Router) {
	r.Route("/user", func(r chi.Router) {
		r.Use(authMiddleware(c.authService))

		r.Get("/me", c.GetMe)
		r.Get("/{user_id}", c.GetUser)
	})
}

// register 		godoc
//
//	@Summary		Get own user detail
//	@Description	Get own user detail
//	@Tags			User
//	@Produce		json
//	@Param			Authorization	header		string	true	"Bearer token"
//	@Success		200				{object}	response.GetUser
//	@Failure		401				{object}	response.ErrorResponse	"Unknown session"
//	@Failure		500				{object}	response.ErrorResponse
//	@Router			/api/v1/user/me [get]
func (c *UserController) GetMe(w http.ResponseWriter, r *http.Request) {
	slog.Info("[CreateServerController] Creating server")

	userId := extractUserId(r.Context())
	if userId == nil {
		render.Render(w, r, response.ParseErrorResponse("Cannot authenticate user", http.StatusUnauthorized, nil))
		return
	}

	user, err := c.userQueries.GetBasic(r.Context(), *userId)
	if err != nil {
		render.Render(w, r, response.ParseErrorResponse("Cannot get user", 500, err))
		return
	}

	render.Status(r, 200)
	render.JSON(w, r, response.GetUser{
		Id:          user.Id,
		CreatedAt:   user.CreatedAt,
		Username:    user.Username,
		DisplayName: user.DisplayName,
		AboutMe:     user.AboutMe,
		Email:       user.Email,
		Disabled:    user.Disabled,
		AvatarUrl:   user.AvatarUrl,
		BannerUrl:   user.BannerUrl,
		Flags:       user.Flags,
	})
}

// register 		godoc
//
//	@Summary		Get user detail
//	@Description	Get user detail by user id
//	@Tags			User
//	@Produce		json
//	@Param			Authorization	header		string	true	"Bearer token"
//	@Param			user_id		path		string	true	"user id to fetch"
//	@Success		200				{object}	response.GetUser
//	@Failure		400				{object}	response.ErrorResponse	"user_id not a valid id"
//	@Failure		401				{object}	response.ErrorResponse	"Unknown session"
//	@Failure		403				{object}	response.ErrorResponse	"Forbidden"
//	@Failure		404				{object}	response.ErrorResponse	"User not found"
//	@Failure		500				{object}	response.ErrorResponse
//	@Router			/api/v1/user/{user_id} [get]
func (c *UserController) GetUser(w http.ResponseWriter, r *http.Request) {
	slog.Info("[CreateServerController] Creating server")

	userId := extractUserId(r.Context())
	if userId == nil {
		render.Render(w, r, response.ParseErrorResponse("Cannot authenticate user", http.StatusUnauthorized, nil))
		return
	}

	targetUserId, err := uuid.Parse(chi.URLParam(r, "user_id"))
	if err != nil {
		render.Render(w, r, response.ParseErrorResponse("Invalid user id", http.StatusBadRequest, err))
		return
	}

	user, err := c.userQueries.GetBasic(r.Context(), targetUserId)
	if err != nil {
		render.Render(w, r, response.ParseErrorResponse("Cannot get user", 500, err))
		return
	}

	render.Status(r, 200)
	render.JSON(w, r, response.GetUser{
		Id:          user.Id,
		CreatedAt:   user.CreatedAt,
		Username:    user.Username,
		DisplayName: user.DisplayName,
		AboutMe:     user.AboutMe,
		Email:       user.Email,
		Disabled:    user.Disabled,
		AvatarUrl:   user.AvatarUrl,
		BannerUrl:   user.BannerUrl,
		Flags:       user.Flags,
	})
}
