package ws

import (
	"backend/internal/domain/entities"
	"context"
)

func (h *Hub) registerHandlers(ctx context.Context) error {
	// Messages
	if err := h.eventSubscriber.Subscribe(ctx, entities.EventMessageCreated, h.messageCreatedHandler); err != nil {
		return err
	}
	if err := h.eventSubscriber.Subscribe(ctx, entities.EventMessageEdited, h.messageEditedHandler); err != nil {
		return err
	}
	if err := h.eventSubscriber.Subscribe(ctx, entities.EventMessageDeleted, h.messageDeletedHandler); err != nil {
		return err
	}

	// Channels
	if err := h.eventSubscriber.Subscribe(ctx, entities.EventChannelCreated, h.channelCreatedHandler); err != nil {
		return err
	}
	if err := h.eventSubscriber.Subscribe(ctx, entities.EventChannelDeleted, h.channelDeletedHandler); err != nil {
		return err
	}
	if err := h.eventSubscriber.Subscribe(ctx, entities.EventChannelNameUpdated, h.channelNameUpdatedHandler); err != nil {
		return err
	}
	if err := h.eventSubscriber.Subscribe(ctx, entities.EventChannelOverwriteUpserted, h.channelOverwriteUpsertedHandler); err != nil {
		return err
	}
	if err := h.eventSubscriber.Subscribe(ctx, entities.EventChannelOverwriteDeleted, h.channelOverwriteDeletedHandler); err != nil {
		return err
	}
	if err := h.eventSubscriber.Subscribe(ctx, entities.EventChannelOrderChanged, h.channelOrderChangedHandler); err != nil {
		return err
	}
	if err := h.eventSubscriber.Subscribe(ctx, entities.EventChannelDescriptionUpdated, h.channelDescriptionUpdatedHandler); err != nil {
		return err
	}
	if err := h.eventSubscriber.Subscribe(ctx, entities.EventChannelParentCategoryChanged, h.channelParentCategoryChangedHandler); err != nil {
		return err
	}

	// Servers
	if err := h.eventSubscriber.Subscribe(ctx, entities.EventServerCreated, h.serverCreatedHandler); err != nil {
		return err
	}
	if err := h.eventSubscriber.Subscribe(ctx, entities.EventServerDeleted, h.serverDeletedHandler); err != nil {
		return err
	}
	if err := h.eventSubscriber.Subscribe(ctx, entities.EventServerNameUpdated, h.serverNameUpdatedHandler); err != nil {
		return err
	}
	if err := h.eventSubscriber.Subscribe(ctx, entities.EventServerBannerURLUpdated, h.serverBannerURLUpdatedHandler); err != nil {
		return err
	}
	if err := h.eventSubscriber.Subscribe(ctx, entities.EventServerIconURLUpdated, h.serverIconURLUpdatedHandler); err != nil {
		return err
	}
	if err := h.eventSubscriber.Subscribe(ctx, entities.EventServerDefaultPermissionChanged, h.serverDefaultPermissionChangedHandler); err != nil {
		return err
	}

	return nil
}
