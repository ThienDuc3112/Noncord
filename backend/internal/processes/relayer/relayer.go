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
	broker ports.EventsBroker
	cfg    Config
}

func New(log *slog.Logger, reader ports.OutboxReader, broker ports.EventsBroker, config Config) *Relayer {
	return &Relayer{log, reader, broker, config}
}

func (r *Relayer) step(ctx context.Context) error {
	records, err := r.reader.ClaimBatch(ctx, r.cfg.BatchSize, r.cfg.StaleAfter)
	if err != nil {
		return err
	}

	if len(records) == 0 {
		return nil
	}

	for _, rec := range records {
		header := map[string]string{
			"event_type":     rec.EventType,
			"aggregate_name": rec.AggregateName,
			"schema_version": strconv.Itoa(int(rec.SchemaVersion)),
			"occurred_at":    rec.OccurredAt.UTC().Format(time.RFC3339Nano),
			"event_id":       rec.ID.String(),
			"topic":          r.cfg.Topic,
		}

		err = r.broker.Publish(ctx, ports.EventMessage{
			AggregateId: rec.AggregateID,
			Payload:     rec.Payload,
			Headers:     header,
		})
		if err != nil {
			if rec.Attempts >= r.cfg.MaxAttempts {
				r.reader.MarkFailed(ctx, rec.ID)
			}
			return err
		}

		err = r.reader.MarkDispatched(ctx, rec.ID)
		if err != nil {
			return err
		}
	}

	return nil
}

func (r *Relayer) Run(ctx context.Context) error {
	tickerCh := time.Tick(r.cfg.PollInterval)

	for {
		select {
		case <-ctx.Done():
			r.broker.Close(ctx)
			return ctx.Err()
		case <-tickerCh:
			if err := r.step(ctx); err != nil {
				r.log.Error("relayer step failed", "err", err)
			}
		}
	}
}
