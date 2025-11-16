package relayer

import (
	"backend/internal/application/ports"
	"context"
	"log/slog"
	"strconv"
	"time"
)

type Config struct {
	BatchSize    int32
	StaleAfter   time.Duration // lease duration, e.g. 60 * time.Second
	MaxAttempts  int32         // e.g. 8
	PollInterval time.Duration // e.g. 100 * time.Millisecond
	Topic        string
	ShardCount   int32
}

type Relayer struct {
	log    *slog.Logger
	reader ports.OutboxReader
	broker ports.EventPublisher
	cfg    Config
}

func New(log *slog.Logger, reader ports.OutboxReader, broker ports.EventPublisher, config Config) *Relayer {
	return &Relayer{log, reader, broker, config}
}

func (r *Relayer) step(ctx context.Context) (int32, error) {
	records, err := r.reader.ClaimBatch(ctx, r.cfg.BatchSize, r.cfg.StaleAfter)
	if err != nil {
		return 0, err
	}

	if len(records) == 0 {
		return 0, nil
	}

	for i, rec := range records {
		header := map[string]any{
			"event_type":     rec.EventType,
			"aggregate_name": rec.AggregateName,
			"schema_version": strconv.Itoa(int(rec.SchemaVersion)),
			"occurred_at":    rec.OccurredAt.UTC().Format(time.RFC3339Nano),
			"event_id":       rec.ID.String(),
			"topic":          r.cfg.Topic,
		}

		err = r.broker.Publish(ctx, ports.EventMessage{
			AggregateId: rec.AggregateID,
			EventType:   rec.EventType,
			Payload:     rec.Payload,
			Headers:     header,
		})
		if err != nil {
			if rec.Attempts >= r.cfg.MaxAttempts {
				r.reader.MarkFailed(ctx, rec.ID)
			}
			return int32(i), err
		}

		err = r.reader.MarkDispatched(ctx, rec.ID)
		if err != nil {
			return int32(i), err
		}
	}

	return int32(len(records)), nil
}

func (r *Relayer) Run(ctx context.Context) error {
	defer r.broker.Close(ctx)
	tickerCh := time.Tick(r.cfg.PollInterval)

	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-tickerCh:
			count, err := r.step(ctx)
			if err != nil {
				r.log.Error("relayer step failed", "err", err, "count", count)
			} else if count > 0 {
				r.log.Log(ctx, slog.LevelInfo, "Delivered events", "count", count)
			}
		}
	}
}
