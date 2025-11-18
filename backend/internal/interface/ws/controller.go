package ws

import (
	"net/http"

	"github.com/gorilla/websocket"
)

func ServeWs(hub *Hub, w http.ResponseWriter, r *http.Request) {
	conn, err := Upgrader.Upgrade(w, r, nil)
	if err != nil {
		return
	}
	wr, err := conn.NextWriter(websocket.TextMessage)
	wr.Write([]byte("hola"))
	conn.Close()
}
