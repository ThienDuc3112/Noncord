package ws

import (
	"log/slog"
	"sync/atomic"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

type client struct {
	id     uuid.UUID
	userId uuid.UUID
	conn   *websocket.Conn

	writeChan chan any
	isClose   atomic.Bool
}

func newClient(userId uuid.UUID, conn *websocket.Conn) *client {
	c := &client{
		id:     uuid.New(),
		userId: userId,
		conn:   conn,

		writeChan: make(chan any, 512),
		isClose:   atomic.Bool{},
	}
	c.isClose.Store(false)

	go c.writePump()
	go c.readPump()

	return c
}

func (c *client) Close() error {
	if c.isClose.Load() {
		return nil
	}
	c.isClose.Store(true)
	close(c.writeChan)
	return c.conn.Close()
}

func (c *client) Write(msg any) {
	if c.isClose.Load() {
		return
	}
	c.writeChan <- msg
}

func (c *client) writePump() {
	for msg := range c.writeChan {
		err := c.conn.WriteJSON(msg)
		if websocket.IsUnexpectedCloseError(err) {
			slog.Warn("write message to a closed client", "client", c.toSlogVal())
			break
		} else if err != nil {
			slog.Warn("cannot send message", "client", c.toSlogVal(), "error", err)
		}
	}
}

func (c *client) readPump() {
	defer c.Close()

	c.conn.SetReadLimit(512)
	c.conn.SetReadDeadline(time.Now().Add(60 * time.Second))
	c.conn.SetPongHandler(func(string) error {
		// extend deadline on pong
		c.conn.SetReadDeadline(time.Now().Add(60 * time.Second))
		return nil
	})

	for !c.isClose.Load() {
		_, msg, err := c.conn.ReadMessage()
		if err != nil {
			break
		}

		_ = msg
	}
}

func (c *client) toSlogVal() slog.Value {
	return slog.GroupValue(
		slog.Attr{Key: "conn_id", Value: slog.StringValue(c.id.String())},
		slog.Attr{Key: "user_id", Value: slog.StringValue(c.userId.String())},
	)
}
