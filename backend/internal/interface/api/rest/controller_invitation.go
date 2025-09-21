package rest

import (
	"backend/internal/application/command"
	"backend/internal/application/interfaces"
	"backend/internal/application/query"
	"backend/internal/interface/api/rest/dto/request"
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
	memberService     interfaces.MembershipService
}

func NewInvitationController(serverService interfaces.ServerService, authService interfaces.AuthService, invitationService interfaces.InviteService, memberService interfaces.MembershipService) *InvitationController {
	return &InvitationController{serverService: serverService, authService: authService, invitationService: invitationService, memberService: memberService}
}

func (c *InvitationController) RegisterRoute(r chi.Router) {
	r.Route("/invitations", func(r chi.Router) {
		r.Use(authMiddleware(c.authService))

		r.Post("/{invitation_id}/join", c.GetInvitationController)
		r.Get("/{invitation_id}", c.GetInvitationController)
		r.Put("/{invitation_id}", c.UpdateInvitationController)
		r.Patch("/{invitation_id}", c.UpdateInvitationController)
		r.Delete("/{invitation_id}", c.DeleteInvitationController)
	})
}

// register 		godoc
//
//	@Summary		Get invitation detail
//	@Description	Get an invitation detail by invitation id
//	@Tags			Invitation
//	@Produce		json
//	@Param			Authorization	header		string	true	"Bearer token"
//	@Param			invitation_id	path		string	true	"invite id to fetch"
//	@Success		200				{object}	response.GetInvitationResponse
//	@Failure		400				{object}	response.ErrorResponse	"Invalid invitation id"
//	@Failure		404				{object}	response.ErrorResponse	"Invitation not found"
//	@Failure		500				{object}	response.ErrorResponse
//	@Router			/api/v1/invitations/{invitation_id} [get]
func (c *InvitationController) GetInvitationController(w http.ResponseWriter, r *http.Request) {
	log.Println("[GetInvitationController] Getting invitation")

	invitationId, err := uuid.Parse(chi.URLParam(r, "invitation_id"))
	if err != nil {
		render.Render(w, r, response.ParseErrorResponse("Invalid invitation id", http.StatusBadRequest, err))
		return
	}

	invitation, err := c.invitationService.GetInvitationById(r.Context(), query.GetInvitation{
		InvitationId: invitationId,
	})
	if err != nil {
		render.Render(w, r, response.ParseErrorResponse("Cannot get invitation", 500, err))
		return
	}

	user := extractUser(r.Context())
	if user == nil {
		render.Render(w, r, response.ParseErrorResponse("Cannot authenticate user", http.StatusUnauthorized, nil))
		return
	}

	server, err := c.serverService.Get(r.Context(), query.GetServer{
		ServerId: invitation.Result.ServerId,
		UserId:   user.Id,
	})
	if err != nil {
		render.Render(w, r, response.ParseErrorResponse("Cannot get invitation info", 500, err))
		return
	}

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
//	@Param			invitation_id	path		string	true	"invite id to join server"
//	@Success		200				{object}	response.JoinServerResponse
//	@Failure		400				{object}	response.ErrorResponse	"Invalid invitation id"
//	@Failure		404				{object}	response.ErrorResponse	"Invitation not found"
//	@Failure		500				{object}	response.ErrorResponse
//	@Router			/api/v1/invitations/{invitation_id}/join [post]
func (c *InvitationController) JoinServerController(w http.ResponseWriter, r *http.Request) {
	log.Println("[JoinServerController] Joining server")

	invitationId, err := uuid.Parse(chi.URLParam(r, "invitation_id"))
	if err != nil {
		render.Render(w, r, response.ParseErrorResponse("Invalid invitation id", http.StatusBadRequest, err))
		return
	}

	user := extractUser(r.Context())
	if user == nil {
		render.Render(w, r, response.ParseErrorResponse("Cannot authenticate user", http.StatusUnauthorized, nil))
		return
	}

	membership, err := c.memberService.JoinServer(r.Context(), command.JoinServerCommand{
		UserId:       user.Id,
		InvitationId: invitationId,
	})
	if err != nil {
		render.Render(w, r, response.ParseErrorResponse("Cannot join server", 500, err))
		return
	}

	server, err := c.serverService.Get(r.Context(), query.GetServer{
		ServerId: membership.Result.ServerId,
		UserId:   membership.Result.UserId,
	})

	render.Status(r, 200)
	render.JSON(w, r, response.JoinServerResponse{
		Server: response.ServerPreview{
			Id:        server.Preview.Id,
			Name:      server.Preview.Name,
			IconUrl:   server.Preview.IconUrl,
			BannerUrl: server.Preview.BannerUrl,
		},
		Membership: response.Membership{
			ServerId:  membership.Result.ServerId,
			UserId:    membership.Result.UserId,
			Nickname:  membership.Result.Nickname,
			CreatedAt: membership.Result.CreatedAt,
		},
	})
}

// register 		godoc
//
//	@Summary		Update invitation
//	@Description	Update invitation detail by invitation id
//	@Tags			Invitation
//	@Produce		json
//	@Param			Authorization	header		string						true	"Bearer token"
//	@Param			invitation_id	path		string						true	"invite id to fetch"
//	@Param			payload			body		request.UpdateInvitation	true	"payload"
//	@Success		200				{object}	response.Invitation			"Updated invitation"
//	@Failure		400				{object}	response.ErrorResponse		"Invalid body or invalid invitation id"
//	@Failure		401				{object}	response.ErrorResponse		"Cannot authenticate user"
//	@Failure		403				{object}	response.ErrorResponse		"Forbidden action"
//	@Failure		404				{object}	response.ErrorResponse		"Invitation not found"
//	@Failure		500				{object}	response.ErrorResponse
//	@Router			/api/v1/invitations/{invitation_id} [put]
//	@Router			/api/v1/invitations/{invitation_id} [patch]
func (c *InvitationController) UpdateInvitationController(w http.ResponseWriter, r *http.Request) {
	log.Println("[GetInvitationController] Getting invitation")

	invitationId, err := uuid.Parse(chi.URLParam(r, "invitation_id"))
	if err != nil {
		render.Render(w, r, response.ParseErrorResponse("Invalid invitation id", http.StatusBadRequest, err))
		return
	}

	invitation, err := c.invitationService.GetInvitationById(r.Context(), query.GetInvitation{
		InvitationId: invitationId,
	})
	if err != nil {
		render.Render(w, r, response.ParseErrorResponse("Cannot get invitation", 500, err))
		return
	}

	body := request.UpdateInvitation{}
	if err := render.Bind(r, &body); err != nil {
		render.Render(w, r, response.ParseErrorResponse("Invalid body", http.StatusBadRequest, err))
		return
	}

	user := extractUser(r.Context())
	if user == nil {
		render.Render(w, r, response.ParseErrorResponse("Cannot authenticate user", http.StatusUnauthorized, nil))
		return
	}

	newInv, err := c.invitationService.UpdateInvitation(r.Context(), command.UpdateInvitationCommand{
		InvitationId: invitation.Result.Id,
		UserId:       user.Id,
		Updates:      command.UpdateInvitationOption(body),
	})
	if err != nil {
		render.Render(w, r, response.ParseErrorResponse("Cannot update invitation info", 500, err))
		return
	}

	render.Status(r, 200)
	render.JSON(w, r, response.Invitation{
		Id:             newInv.Result.Id,
		ServerId:       newInv.Result.ServerId,
		CreatedAt:      newInv.Result.CreatedAt,
		ExpiresAt:      newInv.Result.ExpiresAt,
		BypassApproval: newInv.Result.BypassApproval,
		JoinLimit:      newInv.Result.JoinLimit,
		JoinCount:      newInv.Result.JoinCount,
	})
}

// register 		godoc
//
//	@Summary		Invalidate invitation
//	@Description	Invalidate an invitation by invitation id
//	@Tags			Invitation
//	@Produce		json
//	@Param			Authorization	header		string	true	"Bearer token"
//	@Param			invitation_id	path		string	true	"invite id to fetch"
//	@Success		204				{object}	nil
//	@Failure		400				{object}	response.ErrorResponse	"Invalid invitation id"
//	@Failure		401				{object}	response.ErrorResponse	"Cannot authenticate user"
//	@Failure		403				{object}	response.ErrorResponse	"Forbidden action"
//	@Failure		404				{object}	response.ErrorResponse	"Invitation not found"
//	@Failure		500				{object}	response.ErrorResponse
//	@Router			/api/v1/invitations/{invitation_id} [delete]
func (c *InvitationController) DeleteInvitationController(w http.ResponseWriter, r *http.Request) {
	log.Println("[DeleteInvitationController] Invalidating invitation")

	invitationId, err := uuid.Parse(chi.URLParam(r, "invitation_id"))
	if err != nil {
		render.Render(w, r, response.ParseErrorResponse("Invalid invitation id", http.StatusBadRequest, err))
		return
	}

	invitation, err := c.invitationService.GetInvitationById(r.Context(), query.GetInvitation{
		InvitationId: invitationId,
	})
	if err != nil {
		render.Render(w, r, response.ParseErrorResponse("Cannot get invitation", 500, err))
		return
	}

	user := extractUser(r.Context())
	if user == nil {
		render.Render(w, r, response.ParseErrorResponse("Cannot authenticate user", http.StatusUnauthorized, nil))
		return
	}

	err = c.invitationService.InvalidateInvitation(r.Context(), command.InvalidateInvitationCommand{
		InvitationId: invitation.Result.Id,
		UserId:       user.Id,
	})
	if err != nil {
		render.Render(w, r, response.ParseErrorResponse("Cannot get invitation info", 500, err))
		return
	}

	render.NoContent(w, r)
}
