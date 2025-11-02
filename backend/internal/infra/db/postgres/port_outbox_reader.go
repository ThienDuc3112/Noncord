package postgres

import (
	"backend/internal/application/ports"
	"backend/internal/infra/db/postgres/gen"
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/gookit/goutil/arrutil"
	"github.com/jackc/pgx/v5/pgxpool"
)

type PGOutboxReader struct {
	pool *pgxpool.Pool
	q    *gen.Queries
}

func NewPGOutboxReader(pool *pgxpool.Pool) ports.OutboxReader {
	return &PGOutboxReader{pool, gen.New(pool)}
}

func (r *PGOutboxReader) ClaimBatch(ctx context.Context, limit int32, staleAfter time.Duration) ([]ports.OutboxRecord, error) {
	batch, err := r.q.ClaimOutboxBatch(ctx, gen.ClaimOutboxBatchParams{
		StaleAfter: staleAfter,
		Limit:      limit,
	})
	if err != nil {
		return nil, err
	}

	return arrutil.Map(batch, func(row gen.Outbox) (target ports.OutboxRecord, find bool) {
		return ports.OutboxRecord{
			ID:            row.ID,
			AggregateName: row.AggregateName,
			AggregateID:   row.AggregateID,
			EventType:     row.EventType,
			SchemaVersion: row.SchemaVersion,
			OccurredAt:    row.OccurredAt,
			Status:        row.Status,
			Attempts:      row.Attempts,
			ClaimedAt:     row.ClaimedAt,
			PublishedAt:   row.PublishedAt,
			Payload:       row.Payload,
		}, true
	}), nil
}

func (r *PGOutboxReader) MarkDispatched(ctx context.Context, id uuid.UUID) error {
	return r.q.MarkedOutboxDispatched(ctx, id)
}

func (r *PGOutboxReader) Requeue(ctx context.Context, id uuid.UUID) error {
	return r.q.RequeueOutbox(ctx, id)
}

func (r *PGOutboxReader) MarkFailed(ctx context.Context, id uuid.UUID) error {
	return r.q.MarkedOutboxFailed(ctx, id)
}
