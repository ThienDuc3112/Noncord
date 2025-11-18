package ws

import (
	"backend/internal/application/ports"
	"backend/internal/domain/entities"
	"backend/internal/domain/events"
	"context"
	"log/slog"
)

func (h *Hub) messageCreatedHandler(ctx context.Context, event ports.EventMessage) error {
	h.logger.Info("Incoming message", "event", slog.GroupValue(
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
		h.logger.Warn("Unabled to parse event", "error", err)
		return err
	}

	h.logger.Info("Parsed event", "event", e)
	return nil
}

func (h *Hub) messageEditedHandler(ctx context.Context, event ports.EventMessage) error {
	return nil
}

func (h *Hub) messageDeletedHandler(ctx context.Context, event ports.EventMessage) error {
	return nil
}
