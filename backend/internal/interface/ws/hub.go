package ws

import (
	"backend/internal/application/interfaces"
	"backend/internal/application/ports"
	"context"
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

type Hub struct {
	userConn   map[uuid.UUID]map[uuid.UUID]*client
	serverSub  map[uuid.UUID]map[uuid.UUID]bool
	channelSub map[uuid.UUID]map[uuid.UUID]bool

	visibilityService interfaces.VisibilityQueries
	authService       interfaces.AuthService
	eventSubscriber   ports.EventSubscriber

	unsubChan chan *client

	nicknameCache ports.CacheStore
	userResolver  ports.UserResolver

	m sync.RWMutex
}

var Upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

func NewHub(ctx context.Context, authService interfaces.AuthService, visibilityQueries interfaces.VisibilityQueries, eventReader ports.EventSubscriber, cacheStore ports.CacheStore, userResolver ports.UserResolver) (*Hub, error) {
	hub := &Hub{
		userConn:   make(map[uuid.UUID]map[uuid.UUID]*client),
		serverSub:  make(map[uuid.UUID]map[uuid.UUID]bool),
		channelSub: make(map[uuid.UUID]map[uuid.UUID]bool),

		visibilityService: visibilityQueries,
		authService:       authService,
		eventSubscriber:   eventReader,

		unsubChan: make(chan *client, 1024),

		nicknameCache: cacheStore,
		userResolver:  userResolver,
	}

	if err := hub.registerHandlers(); err != nil {
		return nil, err
	}
	go hub.unsubLoop(ctx)

	return hub, nil
}

func (h *Hub) Register(ctx context.Context, conn *websocket.Conn) error {
	c := newClient(h.authService, conn, h.unsubChan)
	if c == nil {
		return fmt.Errorf("Unauth")
	}
	userId := c.userId

	chans, err := h.visibilityService.GetVisibleChannels(ctx, userId)
	if err != nil {
		c.Close()
		return err
	}
	servers, err := h.visibilityService.GetVisibleServers(ctx, userId)
	if err != nil {
		c.Close()
		return err
	}

	h.m.Lock()

	for _, cId := range chans {
		if _, ok := h.channelSub[cId]; !ok {
			h.channelSub[cId] = make(map[uuid.UUID]bool)
		}
		h.channelSub[cId][userId] = true
	}

	for _, sId := range servers {
		if _, ok := h.serverSub[sId]; !ok {
			h.serverSub[sId] = make(map[uuid.UUID]bool)
		}
		h.serverSub[sId][userId] = true
	}

	if _, ok := h.userConn[userId]; !ok {
		h.userConn[userId] = make(map[uuid.UUID]*client)
	}
	h.userConn[userId][c.id] = c

	h.m.Unlock()

	c.Write(initializedEvent, map[string]any{"subscribedFrom": time.Now()})

	return nil
}

func (h *Hub) unsubLoop(ctx context.Context) {
outer:
	for {
		select {
		case <-ctx.Done():
			break outer
		case c, ok := <-h.unsubChan:
			if !ok {
				break
			}
			if c == nil {
				continue
			}
			h.m.Lock()
			if _, ok = h.userConn[c.userId]; ok {
				if _, ok = h.userConn[c.userId][c.id]; ok {
					delete(h.userConn[c.userId], c.id)
				}
				if len(h.userConn[c.userId]) == 0 {
					delete(h.userConn, c.userId)
					for _, v := range h.serverSub {
						delete(v, c.userId)
					}
					for _, v := range h.channelSub {
						delete(v, c.userId)
					}
				}
			}
			h.m.Unlock()
		}
	}
}
