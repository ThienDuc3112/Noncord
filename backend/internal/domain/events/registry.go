package events

import (
	"encoding/json"
	"fmt"
)

type Factory func() DomainEvent

// event_type -> schema_version -> factory
var reg = map[string]map[int]Factory{}

func Register(eventType string, schemaVersion int, f Factory) {
	if reg[eventType] == nil {
		reg[eventType] = make(map[int]Factory)
	}
	reg[eventType][schemaVersion] = f
}

func New(eventType string, schemaVersion int) (DomainEvent, bool) {
	if m, ok := reg[eventType]; ok {
		if f, ok := m[schemaVersion]; ok {
			return f(), true
		}
	}
	return nil, false
}

func ParseEvent(payload []byte) (Base, error) {
	var event Base
	err := json.Unmarshal(payload, &event)
	return event, err
}

func ParseSpecificEvent[T DomainEvent](payload []byte, eventType string, schemaVersion int) (T, error) {
	var e T
	if _, ok := reg[eventType]; !ok {
		return e, fmt.Errorf("event type not found")
	}
	fac, ok := reg[eventType][schemaVersion]
	if !ok {
		return e, fmt.Errorf("event version not found")
	}

	e = fac().(T)
	if err := json.Unmarshal(payload, &e); err != nil {
		return e, err
	}

	return e, nil
}
