package ws

import (
	"backend/internal/application/ports"
	"context"
)

func (h *Hub) serverCreatedHandler(ctx context.Context, event ports.EventMessage) error {
	return nil
}

func (h *Hub) serverDeletedHandler(ctx context.Context, event ports.EventMessage) error {
	return nil
}

func (h *Hub) serverNameUpdatedHandler(ctx context.Context, event ports.EventMessage) error {
	return nil
}

func (h *Hub) serverBannerURLUpdatedHandler(ctx context.Context, event ports.EventMessage) error {
	return nil
}

func (h *Hub) serverIconURLUpdatedHandler(ctx context.Context, event ports.EventMessage) error {
	return nil
}

func (h *Hub) serverDefaultPermissionChangedHandler(ctx context.Context, event ports.EventMessage) error {
	return nil
}
