package postgres

import (
	"backend/internal/domain/events"
	"backend/internal/infra/db/postgres/gen"
	"context"
	"encoding/json"
)

func pullAndPushEvents(ctx context.Context, q *gen.Queries, evts []events.DomainEvent) error {
	for _, evt := range evts {
		payload, err := json.Marshal(evt)
		if err != nil {
			return err
		}

		base := evt.GetBase()
		if _, err := q.InsertEventToOutbox(ctx, gen.InsertEventToOutboxParams{
			ID:            base.EventID,
			AggregateName: base.AggregateName,
			AggregateID:   base.AggregateID,
			EventType:     base.EventType,
			SchemaVersion: int32(base.SchemaVersion),
			OccurredAt:    base.OccurredAt,
			Payload:       payload,
		}); err != nil {
			return err
		}
	}

	return nil
}
