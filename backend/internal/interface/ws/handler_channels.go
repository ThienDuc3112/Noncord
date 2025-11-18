package ws

import (
	"backend/internal/application/ports"
	"context"
)

func (h *Hub) channelCreatedHandler(ctx context.Context, event ports.EventMessage) error {
	return nil
}

func (h *Hub) channelDeletedHandler(ctx context.Context, event ports.EventMessage) error {
	return nil
}

func (h *Hub) channelNameUpdatedHandler(ctx context.Context, event ports.EventMessage) error {
	return nil
}

func (h *Hub) channelOverwriteUpsertedHandler(ctx context.Context, event ports.EventMessage) error {
	return nil
}

func (h *Hub) channelOverwriteDeletedHandler(ctx context.Context, event ports.EventMessage) error {
	return nil
}

func (h *Hub) channelDescriptionUpdatedHandler(ctx context.Context, event ports.EventMessage) error {
	return nil
}

func (h *Hub) channelOrderChangedHandler(ctx context.Context, event ports.EventMessage) error {
	return nil
}

func (h *Hub) channelParentCategoryChangedHandler(ctx context.Context, event ports.EventMessage) error {
	return nil
}
