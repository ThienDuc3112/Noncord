package entities

import (
	"time"

	"github.com/google/uuid"
)

type InvitationId uuid.UUID

type Invitation struct {
	Id             InvitationId
	ServerId       ServerId
	CreatedAt      time.Time
	ExpiresAt      *time.Time
	BypassApproval bool
	JoinLimit      int32
	JoinCount      int32
}

func (i *Invitation) UpdateExpiresAt(expiresAt *time.Time) error {
	i.ExpiresAt = expiresAt
	return nil
}

func (i *Invitation) UpdateBypassApproval(bypass bool) error {
	i.BypassApproval = bypass
	return nil
}

func (i *Invitation) UpdateJoinLimit(joinLimit int32) error {
	i.JoinLimit = joinLimit
	return nil
}

func (i *Invitation) UpdateJoinCount(joinCount int32) error {
	i.JoinCount = joinCount
	return nil
}

func NewInvitation(serverId ServerId, expiresAt *time.Time, bypass bool, joinLimit int32) *Invitation {
	return &Invitation{
		Id:             InvitationId(uuid.New()),
		ServerId:       serverId,
		CreatedAt:      time.Now(),
		ExpiresAt:      expiresAt,
		BypassApproval: bypass,
		JoinLimit:      joinLimit,
		JoinCount:      0,
	}
}
