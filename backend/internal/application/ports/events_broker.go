package ports

import (
	"context"

	"github.com/google/uuid"
)

type EventMessage struct {
	AggregateId uuid.UUID
	Payload     []byte
	Headers     map[string]string
}

type EventsBroker interface {
	Publish(ctx context.Context, msg EventMessage) error
	Close(ctx context.Context) error
}
