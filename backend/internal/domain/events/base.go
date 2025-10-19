package events

import (
	"time"

	"github.com/google/uuid"
)

type DomainEvent interface {
	GetBase() Base
}

type Base struct {
	EventID       uuid.UUID `json:"event_id"`
	AggregateName string    `json:"aggregate"`      // e.g., "server"
	AggregateID   uuid.UUID `json:"aggregate_id"`   // the root's UUID
	EventType     string    `json:"type"`           // e.g., "server.created", "server.name_updated"
	SchemaVersion int       `json:"schema_version"` // payload schema version (1, 2, ...)
	OccurredAt    time.Time `json:"occurred_at"`    // when the domain change happened
}

func (b Base) GetBase() Base { return b }

// NewBase builds a Base with sensible defaults. Call from your event constructors.
func NewBase(aggregateName string, aggregateID uuid.UUID, eventType string, schemaVersion int) Base {
	return Base{
		EventID:       uuid.New(),
		AggregateName: aggregateName,
		AggregateID:   aggregateID,
		EventType:     eventType,
		SchemaVersion: schemaVersion,
		OccurredAt:    time.Now().UTC(),
	}
}

type Recorder struct {
	events []DomainEvent
}

func (r *Recorder) Record(e DomainEvent)       { r.events = append(r.events, e) }
func (r *Recorder) PullsEvents() []DomainEvent { out := r.events; r.events = nil; return out }
