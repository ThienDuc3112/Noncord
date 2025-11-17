package ws

import (
	"github.com/google/uuid"
	"golang.org/x/net/websocket"
)

type Hub struct {
	channelUserMapper map[uuid.UUID]uuid.UUID
	userConn          map[uuid.UUID]*websocket.Conn
}

type Client struct {
}
