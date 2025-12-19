package ws

import (
	"backend/internal/application/interfaces"
	"log/slog"
	"net/http"
)

func ServeWs(hub *Hub, authService interfaces.AuthService, w http.ResponseWriter, r *http.Request) {
	// userId := authMiddleware(authService, w, r)
	// if userId == nil {
	// 	return
	// }

	conn, err := Upgrader.Upgrade(w, r, nil)
	if err != nil {
		return
	}

	if err = hub.Register(r.Context(), conn); err != nil {
		slog.Error("Fail to register to the ws hub", "error", err)
		return
	}
}
