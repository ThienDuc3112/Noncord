package common

import "github.com/google/uuid"

type Role struct {
	Id           uuid.UUID
	Name         string
	Color        uint32
	Priority     uint16
	AllowMention bool
	Permissions  []string
	ServerId     uuid.UUID
}
