package ws

import (
	"context"
	"log/slog"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

func ServeWs(hub *Hub, w http.ResponseWriter, r *http.Request) {
	slog.Default().Info("New request", "agent", r.UserAgent())
	conn, err := Upgrader.Upgrade(w, r, nil)
	if err != nil {
		return
	}
	wr, err := conn.NextWriter(websocket.TextMessage)
	wr.Write([]byte("hola"))
	wr.Close()

	ctx, cancel := context.WithTimeout(r.Context(), 60*time.Second)
	defer cancel()
	go func() {
	outer:
		for {
			select {
			case <-ctx.Done():
				break outer
			default:
				_, msg, err := conn.ReadMessage()
				if err != nil {
					break outer
				}
				slog.Default().Info("Incoming message", "message", msg)
				if string(msg) == "ping" {
					conn.WriteJSON(map[string]any{"message": "pong"})
				}
			}
		}
	}()
	<-ctx.Done()
	conn.Close()
}
