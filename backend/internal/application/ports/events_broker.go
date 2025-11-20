package ports

import (
	"context"

	"github.com/google/uuid"
)

type EventMessage struct {
	AggregateId uuid.UUID
	Payload     []byte
	Headers     map[string]any // Only string but any for amq compatibility
	EventType   string
}

type EventPublisher interface {
	Publish(ctx context.Context, msg EventMessage) error
	Close(ctx context.Context) error
}

type EventSubscriber interface {
	Subscribe(topic string, handler func(context.Context, EventMessage) error) error

	Close() error
}
