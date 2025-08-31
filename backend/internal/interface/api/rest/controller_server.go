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

type ServerController struct {
	serverService interfaces.ServerService
	authService   interfaces.AuthService
}

func NewServerController(serverService interfaces.ServerService, authService interfaces.AuthService) *ServerController {
	return &ServerController{serverService: serverService, authService: authService}
}

func (c *ServerController) RegisterRoute(r chi.Router) {
	r.Route("/server", func(r chi.Router) {
		r.Use(authMiddleware(c.authService))

		r.Post("/", c.CreateServerController)
		r.Get("/{server_id}", c.GetServerController)
		r.Patch("/{server_id}", c.UpdateServerController)
		r.Put("/{server_id}", c.UpdateServerController)
		r.Delete("/{server_id}", c.DeleteServerController)
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
		render.Render(w, r, response.NewErrorResponse("Invalid body", http.StatusBadRequest, err))
		return
	}

	user := extractUser(r.Context())
	if user == nil {
		render.Render(w, r, response.NewErrorResponse("Cannot authenticate user", http.StatusUnauthorized, nil))
		return
	}

	server, err := c.serverService.Create(r.Context(), command.CreateServerCommand{
		UserId: user.Id,
		Name:   body.Name,
	})
	if err != nil {
		render.Render(w, r, response.NewErrorResponse("Cannot create server", 500, err))
		return
	}

	render.Status(r, 200)
	render.JSON(w, r, response.NewServerResponse{
		Id: server.Result.Id,
	})
}

// register 		godoc
//
//	@Summary		Get a server
//	@Description	Get a server by id
//	@Tags			Server
//	@Produce		json
//	@Param			Authorization	header		string	true	"Bearer token"
//	@Param			server_id		path		string	true	"server id to fetch"
//	@Success		200				{object}	response.GetServerResponse
//	@Failure		401				{object}	response.ErrorResponse	"Unknown session"
//	@Failure		403				{object}	response.ErrorResponse	"Forbidden"
//	@Failure		404				{object}	response.ErrorResponse	"Server not found"
//	@Failure		500				{object}	response.ErrorResponse
//	@Router			/api/v1/server/{server_id} [get]
func (c *ServerController) GetServerController(w http.ResponseWriter, r *http.Request) {
	log.Println("[GetServerController] Getting server")

	serverId, err := uuid.Parse(chi.URLParam(r, "server_id"))
	if err != nil {
		render.Render(w, r, response.NewErrorResponse("Invalid server id", http.StatusBadRequest, err))
		return
	}

	user := extractUser(r.Context())
	if user == nil {
		render.Render(w, r, response.NewErrorResponse("Cannot authenticate user", http.StatusUnauthorized, nil))
		return
	}

	server, err := c.serverService.Get(r.Context(), query.GetServer{
		ServerId: serverId,
		UserId:   user.Id,
	})

	if err != nil {
		render.Render(w, r, response.NewErrorResponse("Cannot get server", 500, err))
		return
	}

	// TODO: Fetch channels and Members

	render.Status(r, 200)
	render.JSON(w, r, response.GetServerResponse{
		Id:          server.Result.Id,
		Name:        server.Result.Name,
		Description: server.Result.Description,
		CreatedAt:   server.Result.CreatedAt,
		UpdatedAt:   server.Result.UpdatedAt,
		IconUrl:     server.Result.IconUrl,
		BannerUrl:   server.Result.BannerUrl,
	})
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
func (c *ServerController) UpdateServerController(w http.ResponseWriter, r *http.Request) {
	render.Render(w, r, response.NewErrorResponse("Not implemented", http.StatusNotImplemented, nil))
}

// register     godoc
//
//	@Summary		Delete server
//	@Description	Delete server by id
//	@Tags			Server
//	@Produce		json
//	@Param			Authorization	header		string	true	"Bearer token"
//	@Param			server_id		path		int		true	"Server Id"
//	@Success		201				{object}	nil
//	@Failure		401				{object}	response.ErrorResponse	"Unknown session"
//	@Failure		403				{object}	response.ErrorResponse	"Forbidden"
//	@Failure		404				{object}	response.ErrorResponse	"Server not found"
//	@Failure		500				{object}	response.ErrorResponse	"Internal server error"
//	@Router			/api/v1/server/{server_id} [delete]
func (c *ServerController) DeleteServerController(w http.ResponseWriter, r *http.Request) {
	render.Render(w, r, response.NewErrorResponse("Not implemented", http.StatusNotImplemented, nil))
}
