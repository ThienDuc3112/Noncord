package ws

import (
	"context"
	"log/slog"
	"net/http"
	"time"
)

func ServeWs(hub *Hub, w http.ResponseWriter, r *http.Request) {
	slog.Default().Info("New request", "agent", r.UserAgent())
	conn, err := Upgrader.Upgrade(w, r, nil)
	if err != nil {
		return
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	go func() {
		<-ctx.Done()
		conn.Close()
	}()

	go func() {
	outer:
		for {
			_, msg, err := conn.ReadMessage()
			if err != nil {
				break outer
			}
			slog.Default().Info("Incoming message", "message", msg)
			if string(msg) == "ping" {
				conn.WriteJSON(map[string]any{"message": "pong"})
			}
		}
	}()
}
