package rest

import (
	"backend/internal/application/interfaces"
	"backend/internal/application/query"
	"backend/internal/interface/api/rest/dto/response"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/google/uuid"
)

type InvitationController struct {
	serverService     interfaces.ServerService
	authService       interfaces.AuthService
	invitationService interfaces.InviteService
}

func NewInvitationController(serverService interfaces.ServerService, authService interfaces.AuthService, invitationService interfaces.InviteService) *ServerController {
	return &ServerController{serverService: serverService, authService: authService, invitationService: invitationService}
}

func (c *InvitationController) RegisterRoute(r chi.Router) {
	r.Route("/invitation", func(r chi.Router) {
		r.Use(authMiddleware(c.authService))

		r.Post("/{invitation_id}/join", c.GetInvitationController)
		r.Get("/{invitation_id}", c.GetInvitationController)
	})
}

// register 		godoc
//
//	@Summary		Get invitation detail
//	@Description	Get an invitation detail by invitation id
//	@Tags			Invitation
//	@Produce		json
//	@Param			Authorization	header		string	true	"Bearer token"
//	@Param			invitation_id		path		string	true	"invite id to fetch"
//	@Success		200				{object}	response.GetInvitationResponse
//	@Failure		404				{object}	response.ErrorResponse	"Invitation not found"
//	@Failure		500				{object}	response.ErrorResponse
//	@Router			/api/v1/invitation/{invitation_id} [get]
func (c *InvitationController) GetInvitationController(w http.ResponseWriter, r *http.Request) {
	log.Println("[GetInvitationController] Getting invitation")

	invitationId, err := uuid.Parse(chi.URLParam(r, "invitation_id"))
	if err != nil {
		render.Render(w, r, response.NewErrorResponse("Invalid invitation id", http.StatusBadRequest, err))
		return
	}

	invitation, err := c.invitationService.GetInvitationById(r.Context(), query.GetInvitation{
		InvitationId: invitationId,
	})
	if err != nil {
		render.Render(w, r, response.NewErrorResponse("Cannot get invitation", 500, err))
		return
	}

	user := extractUser(r.Context())
	if user == nil {
		render.Render(w, r, response.NewErrorResponse("Cannot authenticate user", http.StatusUnauthorized, nil))
		return
	}

	server, err := c.serverService.Get(r.Context(), query.GetServer{
		ServerId: invitation.Result.ServerId,
		UserId:   user.Id,
	})
	if err != nil {
		render.Render(w, r, response.NewErrorResponse("Cannot get invitation info", 500, err))
		return
	}

	// TODO: Fetch channels and Members

	render.Status(r, 200)
	render.JSON(w, r, response.GetInvitationResponse{
		Id: invitation.Result.Id,
		Server: response.ServerPreview{
			Id:        server.Preview.Id,
			Name:      server.Preview.Name,
			IconUrl:   server.Preview.IconUrl,
			BannerUrl: server.Preview.BannerUrl,
		},
	})
}

// register 		godoc
//
//	@Summary		Join a server
//	@Description	Join a server through an invitation
//	@Tags			Invitation
//	@Produce		json
//	@Param			Authorization	header		string	true	"Bearer token"
//	@Param			invitation_id		path		string	true	"invite id to join server"
//	@Success		200				{object}	nil
//	@Failure		404				{object}	response.ErrorResponse	"Invitation not found"
//	@Failure		500				{object}	response.ErrorResponse
//	@Router			/api/v1/invitation/{invitation_id}/join [post]
func (c *InvitationController) JoinServerController(w http.ResponseWriter, r *http.Request) {
	log.Println("[JoinServerController] Joining server")

	invitationId, err := uuid.Parse(chi.URLParam(r, "invitation_id"))
	if err != nil {
		render.Render(w, r, response.NewErrorResponse("Invalid invitation id", http.StatusBadRequest, err))
		return
	}

	_, err = c.invitationService.GetInvitationById(r.Context(), query.GetInvitation{
		InvitationId: invitationId,
	})
	if err != nil {
		render.Render(w, r, response.NewErrorResponse("Cannot get invitation", 500, err))
		return
	}

	user := extractUser(r.Context())
	if user == nil {
		render.Render(w, r, response.NewErrorResponse("Cannot authenticate user", http.StatusUnauthorized, nil))
		return
	}

	render.Render(w, r, response.NewErrorResponse("Have not implemented", http.StatusNotImplemented, nil))
}
