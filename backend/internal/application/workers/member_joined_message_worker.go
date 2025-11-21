package workers

import (
	"backend/internal/application/command"
	"backend/internal/application/interfaces"
	"backend/internal/application/ports"
	"backend/internal/domain/entities"
	"backend/internal/domain/events"
	"context"
	"log/slog"
)

type announcementWorker struct {
	messageSvc interfaces.MessageService
}

func (w *announcementWorker) SendJoinMessageHandler(ctx context.Context, event ports.EventMessage) error {
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

	m, err := events.ParseSpecificEvent[entities.MembershipCreated](event.Payload, entities.EventMembershipCreated, entities.MembershipCreatedSchemaVersion)
	if err != nil {
		slog.Default().Warn("Unabled to parse event", "error", err)
		return err
	}

	err = w.messageSvc.CreateSystemMessage(ctx, command.CreateSystemMessageCommand{
		ServerId: m.ServerID,
		Content:  ("<@" + m.UserID.String() + "> joined the server!"),
	})
	if err != nil {
		slog.Default().Warn("Unabled to send message", "error", err)
	}

	return nil
}
