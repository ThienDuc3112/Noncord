package entities

import (
	"backend/internal/domain/events"
	"time"

	"github.com/google/uuid"
)

type InvitationId uuid.UUID

type Invitation struct {
	events.Recorder

	Id             InvitationId
	ServerId       ServerId
	CreatedAt      time.Time
	ExpiresAt      *time.Time
	BypassApproval bool
	JoinLimit      int32
	JoinCount      int32
}

func (i *Invitation) UpdateExpiresAt(expiresAt *time.Time) error {
	changed := (i.ExpiresAt != nil && expiresAt == nil) || (i.ExpiresAt == nil && expiresAt != nil) || (i.ExpiresAt != nil && expiresAt != nil && !i.ExpiresAt.Equal(*expiresAt))
	if changed {
		old := i.ExpiresAt
		i.ExpiresAt = expiresAt
		i.Record(NewInvitationUpdateExpiresAt(i, old))
	}
	return nil
}

func (i *Invitation) UpdateBypassApproval(bypass bool) error {
	if i.BypassApproval != bypass {
		old := i.BypassApproval
		i.BypassApproval = bypass
		i.Record(NewInvitationUpdateBypassApproval(i, old))
	}
	return nil
}

func (i *Invitation) UpdateJoinLimit(joinLimit int32) error {
	if joinLimit != i.JoinLimit {
		old := i.JoinLimit
		i.JoinLimit = joinLimit
		i.Record(NewInvitationUpdateJoinLimit(i, old))
	}
	return nil
}

func (i *Invitation) UpdateJoinCount(joinCount int32) error {
	if i.JoinCount != joinCount {
		old := i.JoinCount
		i.JoinCount = joinCount
		i.Record(NewInvitationUpdateJoinCount(i, old))
	}
	return nil
}

func (i *Invitation) Invalidate() error {
	now := time.Now()
	old := i.ExpiresAt
	i.ExpiresAt = &now
	i.Record(NewInvitationInvalidated(i, now, old))
	return nil
}

func NewInvitation(serverId ServerId, expiresAt *time.Time, bypass bool, joinLimit int32) *Invitation {
	i := &Invitation{
		Id:             InvitationId(uuid.New()),
		ServerId:       serverId,
		CreatedAt:      time.Now(),
		ExpiresAt:      expiresAt,
		BypassApproval: bypass,
		JoinLimit:      joinLimit,
		JoinCount:      0,
	}
	i.Record(NewInvitationCreatedAt(i))
	return i
}
