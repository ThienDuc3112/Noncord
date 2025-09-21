package common

import (
	"time"

	"github.com/google/uuid"
)

type Membership struct {
	ServerId  uuid.UUID
	UserId    uuid.UUID
	Nickname  string
	CreatedAt time.Time
}
