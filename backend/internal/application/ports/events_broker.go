package ports

import (
	"context"

	"github.com/google/uuid"
)

type EventMessage struct {
	AggregateId uuid.UUID
	Payload     []byte
	Headers     map[string]any
	EventType   string
}

type EventPublisher interface {
	Publish(ctx context.Context, msg EventMessage) error
	Close(ctx context.Context) error
}

type EventSubscriber interface {
	Subscribe(ctx context.Context, topic string, handler func(context.Context, EventMessage) error) error

	Close(context.Context) error
}
