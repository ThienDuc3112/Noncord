package events

import (
	"encoding/json"
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
