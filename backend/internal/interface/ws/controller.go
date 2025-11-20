package ws

import (
	"backend/internal/application/interfaces"
	"backend/internal/interface/rest/dto/response"
	"net/http"

	"github.com/go-chi/render"
)

func ServeWs(hub *Hub, authService interfaces.AuthService, w http.ResponseWriter, r *http.Request) {
	userId := authMiddleware(authService, w, r)
	if userId == nil {
		return
	}

	conn, err := Upgrader.Upgrade(w, r, nil)
	if err != nil {
		return
	}

	if err = hub.Register(r.Context(), conn, *userId); err != nil {
		render.Render(w, r, response.ParseErrorResponse("Cannot register user", 500, err))
		return
	}
}
