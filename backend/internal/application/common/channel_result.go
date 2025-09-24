package common

import (
	"time"

	"github.com/google/uuid"
)

type Channel struct {
	Id             uuid.UUID
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      *time.Time
	Name           string
	Description    string
	ServerId       uuid.UUID
	Order          uint16
	ParentCategory *uuid.UUID
}
