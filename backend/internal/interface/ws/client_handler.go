package ws

import (
	"backend/internal/domain/entities"
)

func (h *Hub) registerHandlers() error {
	// Messages
	if err := h.eventSubscriber.Subscribe(entities.EventMessageCreated, h.messageCreatedHandler); err != nil {
		return err
	}
	if err := h.eventSubscriber.Subscribe(entities.EventMessageEdited, h.messageEditedHandler); err != nil {
		return err
	}
	if err := h.eventSubscriber.Subscribe(entities.EventMessageDeleted, h.messageDeletedHandler); err != nil {
		return err
	}

	// Channels
	if err := h.eventSubscriber.Subscribe(entities.EventChannelCreated, h.channelCreatedHandler); err != nil {
		return err
	}
	if err := h.eventSubscriber.Subscribe(entities.EventChannelDeleted, h.channelDeletedHandler); err != nil {
		return err
	}
	if err := h.eventSubscriber.Subscribe(entities.EventChannelNameUpdated, h.channelNameUpdatedHandler); err != nil {
		return err
	}
	if err := h.eventSubscriber.Subscribe(entities.EventChannelOverwriteUpserted, h.channelOverwriteUpsertedHandler); err != nil {
		return err
	}
	if err := h.eventSubscriber.Subscribe(entities.EventChannelOverwriteDeleted, h.channelOverwriteDeletedHandler); err != nil {
		return err
	}
	if err := h.eventSubscriber.Subscribe(entities.EventChannelOrderChanged, h.channelOrderChangedHandler); err != nil {
		return err
	}
	if err := h.eventSubscriber.Subscribe(entities.EventChannelDescriptionUpdated, h.channelDescriptionUpdatedHandler); err != nil {
		return err
	}
	if err := h.eventSubscriber.Subscribe(entities.EventChannelParentCategoryChanged, h.channelParentCategoryChangedHandler); err != nil {
		return err
	}

	// Servers
	if err := h.eventSubscriber.Subscribe(entities.EventServerCreated, h.serverCreatedHandler); err != nil {
		return err
	}
	if err := h.eventSubscriber.Subscribe(entities.EventServerDeleted, h.serverDeletedHandler); err != nil {
		return err
	}
	if err := h.eventSubscriber.Subscribe(entities.EventServerNameUpdated, h.serverNameUpdatedHandler); err != nil {
		return err
	}
	if err := h.eventSubscriber.Subscribe(entities.EventServerBannerURLUpdated, h.serverBannerURLUpdatedHandler); err != nil {
		return err
	}
	if err := h.eventSubscriber.Subscribe(entities.EventServerIconURLUpdated, h.serverIconURLUpdatedHandler); err != nil {
		return err
	}

	return nil
}
