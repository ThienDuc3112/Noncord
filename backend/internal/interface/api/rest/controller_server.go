package rest

import (
	// "backend/internal/application/command"
	"backend/internal/application/interfaces"
	// "backend/internal/domain/entities"
	// "backend/internal/interface/api/rest/dto/request"
	"backend/internal/interface/api/rest/dto/response"
	// "errors"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

type ServerController struct {
	serverService interfaces.ServerService
	authService   interfaces.AuthService
}

func NewServerController(serverService interfaces.ServerService, authService interfaces.AuthService) *ServerController {
	return &ServerController{serverService: serverService, authService: authService}
}

func (sc *ServerController) RegisterRoute(r chi.Router) {
	r.Route("/server", func(r chi.Router) {
		r.Use(authMiddleware(sc.authService))

		r.Post("/", sc.CreateServerController)
		r.Get("/{server_id}", sc.GetServerController)
		r.Patch("/{server_id}", sc.UpdateServerController)
		r.Put("/{server_id}", sc.UpdateServerController)
		r.Delete("/{server_id}", sc.DeleteServerController)
	})
}

// register 		godoc
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
//	@Failure		422				{object}	response.ErrorResponse		"Invalid name"
//	@Failure		500				{object}	response.ErrorResponse
//	@Router			/api/v1/server [post]

func (ac *ServerController) CreateServerController(w http.ResponseWriter, r *http.Request) {
	render.Render(w, r, response.NewErrorResponse("Not implemented", http.StatusNotImplemented, nil))
}

// register 		godoc
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
//	@Failure		422				{object}	response.ErrorResponse	"Not valid server id"
//	@Failure		500				{object}	response.ErrorResponse
//	@Router			/api/v1/server/{servier_id} [get]

func (ac *ServerController) GetServerController(w http.ResponseWriter, r *http.Request) {
	render.Render(w, r, response.NewErrorResponse("Not implemented", http.StatusNotImplemented, nil))
}

// register     godoc
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
//	@Failure		422				{object}	response.ErrorResponse	"Invalid server id or invalid payload"
//	@Failure		500				{object}	response.ErrorResponse	"Internal server error"
//	@Router			/api/v1/server/{server_id} [patch]

func (ac *ServerController) UpdateServerController(w http.ResponseWriter, r *http.Request) {
	render.Render(w, r, response.NewErrorResponse("Not implemented", http.StatusNotImplemented, nil))
}

// register     godoc
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
//	@Failure		422				{object}	response.ErrorResponse	"Not valid server id"
//	@Failure		500				{object}	response.ErrorResponse	"Internal server error"
//	@Router			/api/v1/server/{server_id} [delete]

func (ac *ServerController) DeleteServerController(w http.ResponseWriter, r *http.Request) {
	render.Render(w, r, response.NewErrorResponse("Not implemented", http.StatusNotImplemented, nil))
}
