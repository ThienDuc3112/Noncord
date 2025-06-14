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
	ExpiresAt      *time.Time
	BypassApproval bool
	JoinLimit      int32
}

func NewInvitation(serverId ServerId, expiresAt *time.Time, bypass bool, joinLimit int32) *Invititation {
	return &Invititation{
		Id:             InvititationId(uuid.New()),
		ServerId:       serverId,
		CreatedAt:      time.Now(),
		ExpiresAt:      expiresAt,
		BypassApproval: bypass,
		JoinLimit:      joinLimit,
	}
}
