package entities

import (
	"time"

	"github.com/google/uuid"
)

type InvititationId uuid.UUID

type Invititation struct {
	Id             InvititationId
	ServerId       ServerId
	CreatedAt      time.Time
	ExpiredAt      *time.Time
	BypassApproval bool
	JoinLimit      int32
}
