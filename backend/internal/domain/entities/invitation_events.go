package entities

import (
	"backend/internal/domain/events"
	"time"

	"github.com/google/uuid"
)

const (
	EventInvitationCreated              = "invitation.created"
	EventInvitationUpdateExpiresAt      = "invitation.expires_at_updated"
	EventInvitationUpdateBypassApproval = "invitation.bypass_approval_updated"
	EventInvitationUpdateJoinLimit      = "invitation.join_limit_updated"
	EventInvitationUpdateJoinCount      = "invitation.join_count_updated"
	EventInvitationInvalidated          = "invitation.invalidated"

	InvitationCreatedSchemaVersion              = 1
	InvitationUpdateExpiresAtSchemaVersion      = 1
	InvitationUpdateBypassApprovalSchemaVersion = 1
	InvitationUpdateJoinLimitSchemaVersion      = 1
	InvitationUpdateJoinCountSchemaVersion      = 1
	InvitationInvalidatedSchemaVersion          = 1
)

type InvitationCreatedAt struct {
	events.Base
	ServerId       uuid.UUID  `json:"serverId"`
	ExpiresAt      *time.Time `json:"expiresAt"`
	BypassApproval bool       `json:"bypassApproval"`
	JoinLimit      int32      `json:"joinLimit"`
}

func NewInvitationCreatedAt(i *Invitation) InvitationCreatedAt {
	return InvitationCreatedAt{
		Base:           events.NewBase("invitation", uuid.UUID(i.Id), EventInvitationCreated, InvitationCreatedSchemaVersion),
		ServerId:       uuid.UUID(i.ServerId),
		ExpiresAt:      i.ExpiresAt,
		BypassApproval: i.BypassApproval,
		JoinLimit:      i.JoinLimit,
	}
}

type InvitationUpdateExpiresAt struct {
	events.Base
	Old *time.Time `json:"old"`
	New *time.Time `json:"new"`
}

func NewInvitationUpdateExpiresAt(i *Invitation, old *time.Time) InvitationUpdateExpiresAt {
	return InvitationUpdateExpiresAt{
		Base: events.NewBase("invitation", uuid.UUID(i.Id), EventInvitationUpdateExpiresAt, InvitationUpdateExpiresAtSchemaVersion),
		Old:  old,
		New:  i.ExpiresAt,
	}
}

type InvitationUpdateBypassApproval struct {
	events.Base
	Old bool `json:"old"`
	New bool `json:"new"`
}

func NewInvitationUpdateBypassApproval(i *Invitation, old bool) InvitationUpdateBypassApproval {
	return InvitationUpdateBypassApproval{
		Base: events.NewBase("invitation", uuid.UUID(i.Id), EventInvitationUpdateBypassApproval, InvitationUpdateBypassApprovalSchemaVersion),
		Old:  old,
		New:  i.BypassApproval,
	}
}

type InvitationUpdateJoinLimit struct {
	events.Base
	Old int32 `json:"old"`
	New int32 `json:"new"`
}

func NewInvitationUpdateJoinLimit(i *Invitation, old int32) InvitationUpdateJoinLimit {
	return InvitationUpdateJoinLimit{
		Base: events.NewBase("invitation", uuid.UUID(i.Id), EventInvitationUpdateJoinLimit, InvitationUpdateJoinLimitSchemaVersion),
		Old:  old,
		New:  i.JoinLimit,
	}
}

type InvitationUpdateJoinCount struct {
	events.Base
	Old int32 `json:"old"`
	New int32 `json:"new"`
}

func NewInvitationUpdateJoinCount(i *Invitation, old int32) InvitationUpdateJoinCount {
	return InvitationUpdateJoinCount{
		Base: events.NewBase("invitation", uuid.UUID(i.Id), EventInvitationUpdateJoinCount, InvitationUpdateJoinCountSchemaVersion),
		Old:  old,
		New:  i.JoinCount,
	}
}

type InvitationInvalidated struct {
	events.Base
	At           time.Time  `json:"at"`
	OldExpiresAt *time.Time `json:"oldExpiresAt"`
}

func NewInvitationInvalidated(i *Invitation, at time.Time, oldExpiresAt *time.Time) InvitationInvalidated {
	return InvitationInvalidated{
		Base:         events.NewBase("invitation", uuid.UUID(i.Id), EventInvitationInvalidated, InvitationInvalidatedSchemaVersion),
		At:           at,
		OldExpiresAt: oldExpiresAt,
	}
}

func init() {}
