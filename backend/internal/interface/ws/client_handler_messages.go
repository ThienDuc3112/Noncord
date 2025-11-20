package ws

import (
	"backend/internal/application/ports"
	"backend/internal/domain/entities"
	"backend/internal/domain/events"
	"context"
	"log/slog"
)

func (h *Hub) messageCreatedHandler(ctx context.Context, event ports.EventMessage) error {
	slog.Default().Info("Incoming message", "event", slog.GroupValue(
		slog.Attr{
			Key:   "aggregatedId",
			Value: slog.AnyValue(event.AggregateId),
		},
		slog.Attr{
			Key:   "eventType",
			Value: slog.AnyValue(event.EventType),
		},
		slog.Attr{
			Key:   "headers",
			Value: slog.AnyValue(event.Headers),
		},
	))

	e, err := events.ParseSpecificEvent[entities.MessageCreated](event.Payload, entities.EventMessageCreated, entities.MessageCreatedSchemaVersion)
	if err != nil {
		slog.Default().Warn("Unabled to parse event", "error", err)
		return err
	}

	slog.Default().Info("Parsed event", "event", e)
	h.m.RLock()
	defer h.m.RUnlock()

	if e.ChannelID != nil {
		if _, ok := h.channelSub[*e.ChannelID]; !ok {
			slog.Default().Info("no listener on channel", "channel_id", e.ChannelID.String())
			return nil
		}
		for uId := range h.channelSub[*e.ChannelID] {
			for _, c := range h.userConn[uId] {
				c.Write(map[string]any{
					"content":     e.Content,
					"author":      e.AuthorID,
					"channelId":   e.ChannelID,
					"groupId":     e.GroupID,
					"attachments": e.Attachments,
					"messageId":   e.AggregateID,
				})
			}
		}
	}

	return nil
}

func (h *Hub) messageEditedHandler(ctx context.Context, event ports.EventMessage) error {
	return nil
}

func (h *Hub) messageDeletedHandler(ctx context.Context, event ports.EventMessage) error {
	return nil
}
