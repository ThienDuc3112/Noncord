package rest

import (
	"backend/internal/application/command"
	"backend/internal/application/common"
	"backend/internal/application/interfaces"
	"backend/internal/application/query"
	"backend/internal/domain/entities"
	"backend/internal/interface/dto/request"
	"backend/internal/interface/dto/response"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/google/uuid"
	"github.com/gookit/goutil/arrutil"
)

type ServerController struct {
	authService       interfaces.AuthService
	serverService     interfaces.ServerService
	serverQueries     interfaces.ServerQueries
	invitationService interfaces.InviteService
	invitationQueries interfaces.InviteQueries
}

func NewServerController(
	authService interfaces.AuthService,
	serverService interfaces.ServerService,
	serverQueries interfaces.ServerQueries,
	invitationService interfaces.InviteService,
	invitationQueries interfaces.InviteQueries,
) *ServerController {
	return &ServerController{serverService: serverService, authService: authService, invitationService: invitationService, serverQueries: serverQueries, invitationQueries: invitationQueries}
}

func (c *ServerController) RegisterRoute(r chi.Router) {
	r.Route("/server", func(r chi.Router) {
		r.Use(authMiddleware(c.authService))

		r.Post("/", c.CreateServerController)
		r.Get("/", c.GetServersController)
		r.Get("/{server_id}", c.GetServerController)
		r.Patch("/{server_id}", c.UpdateServerController)
		r.Put("/{server_id}", c.UpdateServerController)
		r.Delete("/{server_id}", c.DeleteServerController)

		r.Get("/{server_id}/invitations", c.GetInvitationController)
		r.Post("/{server_id}/invitations", c.CreateInvitationController)
	})
}

// register 		godoc
//
//	@Summary		Create a server
//	@Description	Create a server
//	@Tags			Server
//	@Accept			json
//	@Produce		json
//	@Param			Authorization	header		string						true	"Bearer token"
//	@Param			payload			body		request.NewServer			true	"Data for creating server"
//	@Success		200				{object}	response.NewServerResponse	"Server id"
//	@Failure		400				{object}	response.ErrorResponse		"Bad request"
//	@Failure		401				{object}	response.ErrorResponse		"Unknown session"
//	@Failure		403				{object}	response.ErrorResponse		"Forbidden"
//	@Failure		500				{object}	response.ErrorResponse
//	@Router			/api/v1/server [post]
func (c *ServerController) CreateServerController(w http.ResponseWriter, r *http.Request) {
	log.Println("[CreateServerController] Creating server")

	body := request.NewServer{}
	if err := render.Bind(r, &body); err != nil {
		render.Render(w, r, response.ParseErrorResponse("Invalid body", http.StatusBadRequest, err))
		return
	}

	userId := extractUserId(r.Context())
	if userId == nil {
		render.Render(w, r, response.ParseErrorResponse("Cannot authenticate user", http.StatusUnauthorized, nil))
		return
	}

	server, err := c.serverService.Create(r.Context(), command.CreateServerCommand{
		UserId:       *userId,
		Name:         body.Name,
		UserNickname: "",
	})
	if err != nil {
		render.Render(w, r, response.ParseErrorResponse("Cannot create server", 500, err))
		return
	}

	render.Status(r, 200)
	render.JSON(w, r, response.NewServerResponse{
		Id: server.Result.Id,
	})
}

// register 		godoc
//
//	@Summary		Get servers by user
//	@Description	Get all servers the user is in
//	@Tags			Server
//	@Produce		json
//	@Param			Authorization	header		string	true	"Bearer token"
//	@Success		200				{object}	response.GetServersResponse
//	@Failure		401				{object}	response.ErrorResponse	"Unknown session"
//	@Failure		403				{object}	response.ErrorResponse	"Forbidden"
//	@Failure		500				{object}	response.ErrorResponse
//	@Router			/api/v1/server [get]
func (c *ServerController) GetServersController(w http.ResponseWriter, r *http.Request) {
	log.Println("[GetServersController] Getting servers by user")

	userId := extractUserId(r.Context())
	if userId == nil {
		render.Render(w, r, response.ParseErrorResponse("Cannot authenticate user", http.StatusUnauthorized, nil))
		return
	}

	servers, err := c.serverQueries.GetServersUserIn(r.Context(), query.GetServersUserIn{
		UserId: *userId,
	})
	if err != nil {
		render.Render(w, r, response.ParseErrorResponse("Cannot get servers", 500, err))
		return
	}

	render.Status(r, 200)
	render.JSON(w, r, response.GetServersResponse{
		Result: arrutil.Map(servers.Result, func(s common.Server) (target response.ServerPreview, find bool) {
			return response.ServerPreview{
				Id:        s.Id,
				Name:      s.Name,
				IconUrl:   s.IconUrl,
				BannerUrl: s.BannerUrl,
			}, true
		}),
	})
}

// register 		godoc
//
//	@Summary		Get a server
//	@Description	Get a server by id, can only get server the user is in
//	@Tags			Server
//	@Produce		json
//	@Param			Authorization	header		string	true	"Bearer token"
//	@Param			server_id		path		string	true	"server id to fetch"
//	@Success		200				{object}	response.GetServerResponse
//	@Failure		400				{object}	response.ErrorResponse	"server_id not a valid id"
//	@Failure		401				{object}	response.ErrorResponse	"Unknown session"
//	@Failure		403				{object}	response.ErrorResponse	"Forbidden"
//	@Failure		404				{object}	response.ErrorResponse	"Server not found"
//	@Failure		500				{object}	response.ErrorResponse
//	@Router			/api/v1/server/{server_id} [get]
func (c *ServerController) GetServerController(w http.ResponseWriter, r *http.Request) {
	log.Println("[GetServerController] Getting server")

	serverId, err := uuid.Parse(chi.URLParam(r, "server_id"))
	if err != nil {
		render.Render(w, r, response.ParseErrorResponse("Invalid server id", http.StatusBadRequest, err))
		return
	}

	userId := extractUserId(r.Context())
	if userId == nil {
		render.Render(w, r, response.ParseErrorResponse("Cannot authenticate user", http.StatusUnauthorized, nil))
		return
	}

	server, err := c.serverQueries.Get(r.Context(), query.GetServer{
		ServerId: serverId,
		UserId:   userId,
	})

	if err != nil {
		render.Render(w, r, response.ParseErrorResponse("Cannot get server", 500, err))
		return
	}

	if server.Full != nil {
		render.Status(r, 200)
		// TODO: update to have role here
		render.JSON(w, r, response.GetServerResponse{
			Id:                  server.Full.Id,
			Name:                server.Full.Name,
			Description:         server.Full.Description,
			CreatedAt:           server.Full.CreatedAt,
			UpdatedAt:           server.Full.UpdatedAt,
			IconUrl:             server.Full.IconUrl,
			BannerUrl:           server.Full.BannerUrl,
			AnnouncementChannel: server.Full.AnnouncementChannel,
			DefaultRole:         server.Full.DefaultRole,
			Channels: arrutil.Map(server.Channel, func(c common.Channel) (target response.Channel, find bool) {
				return response.Channel{
					Id:             c.Id,
					CreatedAt:      c.CreatedAt,
					UpdatedAt:      c.UpdatedAt,
					Name:           c.Name,
					Description:    c.Description,
					ServerId:       c.ServerId,
					Order:          c.Order,
					ParentCategory: c.ParentCategory,
				}, true
			}),
			Roles: arrutil.Map(server.Roles, func(r common.Role) (target response.Role, find bool) {
				return response.Role{
					Id:           r.Id,
					Name:         r.Name,
					Color:        r.Color,
					Priority:     r.Priority,
					AllowMention: r.AllowMention,
					Permissions:  r.Permissions,
					ServerId:     r.ServerId,
				}, true
			}),
			SelfMembership: response.Membership{
				ServerId:      server.Membership.ServerId,
				UserId:        server.Membership.UserId,
				Nickname:      server.Membership.Nickname,
				CreatedAt:     server.Membership.CreatedAt,
				AssignedRoles: server.Membership.Roles,
			},
		})
	} else {
		render.Status(r, 404)
		render.Render(w, r, response.NewErrorResponse("Server not found", http.StatusNotFound, nil))
	}
}

// register     godoc
//
//	@Summary		Update server
//	@Description	Update server
//	@Tags			Server
//	@Accept			json
//	@Produce		json
//	@Param			Authorization	header		string							true	"Bearer token"
//	@Param			payload			body		request.UpdateServer			true	"Update server"
//	@Param			server_id		path		int								true	"Server Id"
//	@Success		200				{object}	response.UpdateServerResponse	"Updated server"
//	@Failure		400				{object}	response.ErrorResponse
//	@Failure		401				{object}	response.ErrorResponse	"Unknown session"
//	@Failure		403				{object}	response.ErrorResponse	"Forbidden"
//	@Failure		404				{object}	response.ErrorResponse	"Server not found"
//	@Failure		500				{object}	response.ErrorResponse	"Internal server error"
//	@Router			/api/v1/server/{server_id} [patch]
//	@Router			/api/v1/server/{server_id} [put]
func (c *ServerController) UpdateServerController(w http.ResponseWriter, r *http.Request) {
	log.Println("[UpdateServerController] Update server")

	serverId, err := uuid.Parse(chi.URLParam(r, "server_id"))
	if err != nil {
		render.Render(w, r, response.ParseErrorResponse("Invalid server id", http.StatusBadRequest, err))
		return
	}

	body := request.UpdateServer{}
	if err := render.Bind(r, &body); err != nil {
		render.Render(w, r, response.ParseErrorResponse("Invalid body", http.StatusBadRequest, err))
		return
	}

	userId := extractUserId(r.Context())
	if userId == nil {
		render.Render(w, r, response.ParseErrorResponse("Cannot authenticate user", http.StatusUnauthorized, nil))
		return
	}

	nullableAnnouncementChannel := uuid.NullUUID{}
	if body.AnnouncementChannel != nil {
		nullableAnnouncementChannel.UUID = *body.AnnouncementChannel
		nullableAnnouncementChannel.Valid = true
	}

	server, err := c.serverService.UpdateMetadata(r.Context(), command.UpdateServerCommand{
		UserId:   *userId,
		ServerId: serverId,
		Updates: command.UpdateServerOption{
			Name:                body.Name,
			Description:         body.Description,
			IconUrl:             body.IconUrl,
			BannerUrl:           body.BannerUrl,
			NeedApproval:        body.NeedApproval,
			AnnouncementChannel: nullableAnnouncementChannel,
		},
	})
	if err != nil {
		render.Render(w, r, response.ParseErrorResponse("cannot update server", 500, err))
		return
	}

	render.Status(r, 200)
	render.JSON(w, r, response.UpdateServerResponse{
		Id:                  server.Result.Id,
		Name:                server.Result.Name,
		Description:         server.Result.Description,
		CreatedAt:           server.Result.CreatedAt,
		UpdatedAt:           server.Result.UpdatedAt,
		IconUrl:             server.Result.IconUrl,
		BannerUrl:           server.Result.BannerUrl,
		AnnouncementChannel: server.Result.AnnouncementChannel,
		DefaultRole:         server.Result.DefaultRole,
	})
}

// register     godoc
//
//	@Summary		Delete server
//	@Description	Delete server by id
//	@Tags			Server
//	@Produce		json
//	@Param			Authorization	header		string	true	"Bearer token"
//	@Param			server_id		path		int		true	"Server Id"
//	@Success		204				{object}	nil
//	@Failure		401				{object}	response.ErrorResponse	"Unknown session"
//	@Failure		403				{object}	response.ErrorResponse	"Forbidden"
//	@Failure		404				{object}	response.ErrorResponse	"Server not found"
//	@Failure		500				{object}	response.ErrorResponse	"Internal server error"
//	@Router			/api/v1/server/{server_id} [delete]
func (c *ServerController) DeleteServerController(w http.ResponseWriter, r *http.Request) {
	log.Println("[DeleteServerController] Delete server")

	serverId, err := uuid.Parse(chi.URLParam(r, "server_id"))
	if err != nil {
		render.Render(w, r, response.ParseErrorResponse("Invalid server id", http.StatusBadRequest, err))
		return
	}

	userId := extractUserId(r.Context())
	if userId == nil {
		render.Render(w, r, response.ParseErrorResponse("Cannot authenticate user", http.StatusUnauthorized, nil))
		return
	}

	err = c.serverService.Delete(r.Context(), command.DeleteServerCommand{
		UserId:   *userId,
		ServerId: serverId,
	})
	if err != nil {
		render.Render(w, r, response.NewErrorResponseFromChatError(err.(*entities.ChatError)))
		return
	}

	render.Status(r, 204)
	render.JSON(w, r, nil)
}

// register     godoc
//
//	@Summary		Get server's invitations
//	@Description	Get all server's invitations
//	@Tags			Server
//	@Produce		json
//	@Param			Authorization	header		string	true	"Bearer token"
//	@Param			server_id		path		int		true	"Server Id"
//	@Success		200				{object}	response.Invitation
//	@Failure		401				{object}	response.ErrorResponse	"Unknown session"
//	@Failure		403				{object}	response.ErrorResponse	"Forbidden"
//	@Failure		404				{object}	response.ErrorResponse	"Server not found"
//	@Failure		500				{object}	response.ErrorResponse	"Internal server error"
//	@Router			/api/v1/server/{server_id}/invitations [get]
func (c *ServerController) GetInvitationController(w http.ResponseWriter, r *http.Request) {
	log.Println("[GetInvitationController] Get invites")

	serverId, err := uuid.Parse(chi.URLParam(r, "server_id"))
	if err != nil {
		render.Render(w, r, response.ParseErrorResponse("Invalid server id", http.StatusBadRequest, err))
		return
	}

	userId := extractUserId(r.Context())
	if userId == nil {
		render.Render(w, r, response.ParseErrorResponse("Cannot authenticate user", http.StatusUnauthorized, nil))
		return
	}

	invs, err := c.invitationQueries.GetInvitationsByServerId(r.Context(), query.GetInvitationsByServerId{
		ServerId: serverId,
		UserId:   *userId,
	})
	if err != nil {
		render.Render(w, r, response.NewErrorResponseFromChatError(err.(*entities.ChatError)))
		return
	}

	render.Status(r, 200)
	render.JSON(w, r, response.GetServerInvitationsResponse{
		Result: arrutil.Map(invs.Result, func(i *common.Invitation) (target response.Invitation, find bool) {
			return response.Invitation{
				Id:             i.Id,
				ServerId:       i.ServerId,
				CreatedAt:      i.CreatedAt,
				ExpiresAt:      i.ExpiresAt,
				BypassApproval: i.BypassApproval,
				JoinLimit:      i.JoinLimit,
				JoinCount:      i.JoinCount,
			}, true
		}),
	})
}

// register     godoc
//
//	@Summary		Create invitation
//	@Description	Get an invitation on a server
//	@Tags			Server
//	@Produce		json
//	@Param			Authorization	header		string					true	"Bearer token"
//	@Param			server_id		path		int						true	"Server Id"
//	@Param			payload			body		request.NewInvitation	true	"Data for creating invitation"
//	@Success		200				{object}	response.Invitation
//	@Failure		401				{object}	response.ErrorResponse	"Unknown session"
//	@Failure		403				{object}	response.ErrorResponse	"Forbidden"
//	@Failure		404				{object}	response.ErrorResponse	"Server not found"
//	@Failure		500				{object}	response.ErrorResponse	"Internal server error"
//	@Router			/api/v1/server/{server_id}/invitations [post]
func (c *ServerController) CreateInvitationController(w http.ResponseWriter, r *http.Request) {
	log.Println("[CreateInvitationController] Creating invite")

	serverId, err := uuid.Parse(chi.URLParam(r, "server_id"))
	if err != nil {
		render.Render(w, r, response.ParseErrorResponse("Invalid server id", http.StatusBadRequest, err))
		return
	}

	body := request.NewInvitation{}
	if err := render.Bind(r, &body); err != nil {
		render.Render(w, r, response.ParseErrorResponse("Invalid body", http.StatusBadRequest, err))
		return
	}

	userId := extractUserId(r.Context())
	if userId == nil {
		render.Render(w, r, response.ParseErrorResponse("Cannot authenticate user", http.StatusUnauthorized, nil))
		return
	}

	inv, err := c.invitationService.CreateInvitation(r.Context(), command.CreateInvitationCommand{
		ServerId:       serverId,
		UserId:         *userId,
		ExpiresAt:      body.ExpiresAt,
		BypassApproval: body.BypassApproval,
		JoinLimit:      body.JoinLimit,
	})
	if err != nil {
		render.Render(w, r, response.NewErrorResponseFromChatError(err.(*entities.ChatError)))
		return
	}

	render.Status(r, 200)
	render.JSON(w, r, response.Invitation{
		Id:             inv.Result.Id,
		ServerId:       inv.Result.ServerId,
		CreatedAt:      inv.Result.CreatedAt,
		ExpiresAt:      inv.Result.ExpiresAt,
		BypassApproval: inv.Result.BypassApproval,
		JoinLimit:      inv.Result.JoinLimit,
		JoinCount:      inv.Result.JoinCount,
	})
}
