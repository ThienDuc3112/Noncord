package ws

import (
	"backend/internal/application/interfaces"
	"context"
	"encoding/json"
	"log/slog"
	"sync/atomic"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

const (
	WS_VERSION      = 1
	WS_AUTH_TIMEOUT = time.Second * 5

	AUTH_MESSAGE = "auth"

	AUTH_FAILED_EVENT = "auth_failed"
)

type wsPayload struct {
	EventType string `json:"eventType"`
	Payload   any    `json:"payload"`
	Version   int32  `json:"version"`
}

type client struct {
	id     uuid.UUID
	userId uuid.UUID
	conn   *websocket.Conn

	authService interfaces.AuthService

	writeChan chan any
	isClose   atomic.Bool
	auth      chan uuid.UUID
	isAuth    atomic.Bool

	unsub chan<- *client
}

func newClient(authService interfaces.AuthService, conn *websocket.Conn, unsub chan<- *client) *client {
	c := &client{
		id:   uuid.New(),
		conn: conn,

		authService: authService,

		writeChan: make(chan any, 512),
		isClose:   atomic.Bool{},
		auth:      make(chan uuid.UUID),

		unsub: unsub,
	}

	go c.writePump()
	go c.readPump()

	c.isClose.Store(false)

	ctx, cancel := context.WithTimeout(context.Background(), WS_AUTH_TIMEOUT)
	defer cancel()

	select {
	case <-ctx.Done():
		c.Close()
		return nil
	case id, ok := <-c.auth:
		if !ok {
			c.Close()
			return nil
		}
		c.userId = id
		return c
	}
}

func (c *client) Close() error {
	if c.isClose.Load() {
		return nil
	}
	c.isClose.Store(true)
	close(c.writeChan)
	return c.conn.Close()
}

func (c *client) Write(eventType string, msg any) {
	if c.isClose.Load() {
		return
	}
	c.writeChan <- wsPayload{
		EventType: eventType,
		Payload:   msg,
		Version:   WS_VERSION,
	}
}

func (c *client) writePump() {
	tickerCh := time.Tick(30 * time.Second)
	for {
		select {
		case _, ok := <-tickerCh:
			if !ok {
				return
			}
			if err := c.conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				slog.Warn("Ping failed to send", "client", c.toSlogVal())
				c.conn.Close()
				return
			}

		case msg, ok := <-c.writeChan:
			if !ok {
				slog.Warn("write channel is closed", "client", c.toSlogVal())
				_ = c.conn.WriteMessage(websocket.CloseMessage, []byte{})
				c.conn.Close()
				return
			}
			err := c.conn.WriteJSON(msg)
			if err != nil {
				slog.Warn("cannot send message", "client", c.toSlogVal(), "error", err)
				c.conn.Close()
			}
		}
	}
}

func (c *client) readPump() {
	slog.Info("Read pump started")
	defer func() {
		slog.Info("closing client", "client", c.toSlogVal())
		c.Close()
	}()

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
			slog.Warn("closing client", "error", err, "client", c.toSlogVal())
			break
		}

		var data wsPayload
		if json.Unmarshal(msg, &data) == nil {
			switch data.EventType {
			case AUTH_MESSAGE:
				str, ok := data.Payload.(string)
				if !ok {
					slog.Info("Unknown payload for auth message", "data", data, "client", c.toSlogVal())
					continue
				}

				userId := authMiddleware(c.authService, str)
				if userId == nil {
					slog.Info("attempt to authenticate token failed", "str", str, "client", c.toSlogVal())
					c.Write(AUTH_FAILED_EVENT, "Authentication failed")
					continue
				}

				slog.Info("Successfully authenticate user", "userId", userId.String(), "client", c.toSlogVal())
				if c.isAuth.CompareAndSwap(false, true) {
					c.auth <- *userId
					close(c.auth)
				}

			default:
				slog.Info("Incoming (unknown) ws message", "msg", string(msg), "client", c.toSlogVal())
			}
		} else {
			slog.Info("Unknown message received", "msg", string(msg), "client", c.toSlogVal())
		}

	}
}

func (c *client) toSlogVal() slog.Value {
	return slog.GroupValue(
		slog.Attr{Key: "conn_id", Value: slog.StringValue(c.id.String())},
		slog.Attr{Key: "user_id", Value: slog.StringValue(c.userId.String())},
	)
}
