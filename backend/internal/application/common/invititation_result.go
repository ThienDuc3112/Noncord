package common

import (
	"time"

	"github.com/google/uuid"
)

type Invitation struct {
	Id             uuid.UUID
	ServerId       uuid.UUID
	CreatedAt      time.Time
	ExpiresAt      *time.Time
	BypassApproval bool
	JoinLimit      int32
	JoinCount      int32
}
