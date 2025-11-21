package workers

import (
	"backend/internal/application/interfaces"
	"backend/internal/application/ports"
	"backend/internal/domain/entities"
	"log/slog"
)

// Recommend a dedicated event subscriber for this
func NewWorker(messageSvc interfaces.MessageService, eventReader ports.EventSubscriber) error {
	if err := eventReader.Subscribe(entities.EventMembershipCreated, (&announcementWorker{messageSvc}).SendJoinMessageHandler); err != nil {
		return err
	}
	slog.Info("Attach worker to event subscriber successfully")
	return nil
}
