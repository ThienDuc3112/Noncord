package ports

import (
	"context"
	"time"

	"github.com/google/uuid"
)

type OutboxRecord struct {
	ID            uuid.UUID
	AggregateName string
	AggregateID   uuid.UUID
	EventType     string
	SchemaVersion int32
	OccurredAt    time.Time

	Payload     []byte
	Status      string
	Attempts    int32
	ClaimedAt   *time.Time
	PublishedAt *time.Time
}

type OutboxReader interface {
	ClaimBatch(ctx context.Context, limit int32, staleAfter time.Duration) ([]OutboxRecord, error)
	MarkDispatched(ctx context.Context, id uuid.UUID) error
	Requeue(ctx context.Context, id uuid.UUID) error
	MarkFailed(ctx context.Context, id uuid.UUID) error
}
