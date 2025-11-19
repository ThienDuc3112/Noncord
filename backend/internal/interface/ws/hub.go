package ws

import (
	"backend/internal/application/interfaces"
	"backend/internal/application/ports"
	"context"
	"sync"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

type Hub struct {
	userConn   map[uuid.UUID]map[uuid.UUID]*client
	serverSub  map[uuid.UUID]map[uuid.UUID]bool
	channelSub map[uuid.UUID]map[uuid.UUID]bool

	permissionService interfaces.PermissionService
	eventSubscriber   ports.EventSubscriber

	m sync.RWMutex
}

var Upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func NewHub(ctx context.Context, permService interfaces.PermissionService, eventReader ports.EventSubscriber) (*Hub, error) {
	hub := &Hub{
		userConn:   make(map[uuid.UUID]map[uuid.UUID]*client),
		serverSub:  make(map[uuid.UUID]map[uuid.UUID]bool),
		channelSub: make(map[uuid.UUID]map[uuid.UUID]bool),

		permissionService: permService,
		eventSubscriber:   eventReader,
	}

	if err := hub.registerHandlers(ctx); err != nil {
		return nil, err
	}

	return hub, nil
}

func (h *Hub) Register(ctx context.Context, conn *websocket.Conn, userId uuid.UUID) error {
	chans, err := h.permissionService.GetVisibleChannels(ctx, userId)
	if err != nil {
		return err
	}
	servers, err := h.permissionService.GetVisibleServers(ctx, userId)
	if err != nil {
		return err
	}

	c := newClient(userId, conn)
	h.m.Lock()
	defer h.m.Unlock()

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

	return nil
}
