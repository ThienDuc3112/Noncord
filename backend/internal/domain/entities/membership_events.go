package entities

import (
	"backend/internal/domain/events"
	"time"

	"github.com/google/uuid"
)

const (
	EventMembershipCreated         = "membership.created"
	EventMembershipRoleAssigned    = "membership.role_assigned"
	EventMembershipRoleUnassigned  = "membership.role_unassigned"
	EventMembershipNicknameChanged = "membership.nickname_changed"
	EventMembershipDeleted         = "membership.deleted"

	MembershipCreatedSchemaVersion         = 1
	MembershipRoleAssignedSchemaVersion    = 1
	MembershipRoleUnassignedSchemaVersion  = 1
	MembershipNicknameChangedSchemaVersion = 1
	MembershipDeletedSchemaVersion         = 1
)

// ------------- Event payloads + constructors -------------

type MembershipCreated struct {
	events.Base
	ServerID uuid.UUID `json:"server_id"`
	UserID   uuid.UUID `json:"user_id"`
	Nickname string    `json:"nickname,omitempty"`
}

func NewMembershipCreated(m *Membership) MembershipCreated {
	return MembershipCreated{
		Base:     events.NewBase("membership", uuid.UUID(m.Id), EventMembershipCreated, MembershipCreatedSchemaVersion),
		ServerID: uuid.UUID(m.ServerId),
		UserID:   uuid.UUID(m.UserId),
		Nickname: m.Nickname,
	}
}

type MembershipRoleAssigned struct {
	events.Base
	RoleID uuid.UUID `json:"role_id"`
}

func NewMembershipRoleAssigned(m *Membership, roleId RoleId) MembershipRoleAssigned {
	return MembershipRoleAssigned{
		Base:   events.NewBase("membership", uuid.UUID(m.Id), EventMembershipRoleAssigned, MembershipRoleAssignedSchemaVersion),
		RoleID: uuid.UUID(roleId),
	}
}

type MembershipRoleUnassigned struct {
	events.Base
	RoleID uuid.UUID `json:"role_id"`
}

func NewMembershipRoleUnassigned(m *Membership, roleId RoleId) MembershipRoleAssigned {
	return MembershipRoleAssigned{
		Base:   events.NewBase("membership", uuid.UUID(m.Id), EventMembershipRoleUnassigned, MembershipRoleUnassignedSchemaVersion),
		RoleID: uuid.UUID(roleId),
	}
}

type MembershipNicknameChanged struct {
	events.Base
	Old string `json:"old"`
	New string `json:"new"`
}

func NewMembershipNicknameChanged(m *Membership, old string) MembershipNicknameChanged {
	return MembershipNicknameChanged{
		Base: events.NewBase("membership", uuid.UUID(m.Id), EventMembershipNicknameChanged, MembershipNicknameChangedSchemaVersion),
		Old:  old,
		New:  m.Nickname,
	}
}

type MembershipDeleted struct {
	events.Base
	DeletedAt time.Time `json:"deleted_at"`
}

func NewMembershipDeleted(m *Membership) MembershipDeleted {
	// Soft delete timestamp captured at event creation.
	return MembershipDeleted{
		Base:      events.NewBase("membership", uuid.UUID(m.Id), EventMembershipDeleted, MembershipDeletedSchemaVersion),
		DeletedAt: time.Now(),
	}
}

func init() {
	events.Register(EventMembershipCreated, MembershipCreatedSchemaVersion, func() events.DomainEvent { return MembershipCreated{} })
	events.Register(EventMembershipRoleAssigned, MembershipRoleAssignedSchemaVersion, func() events.DomainEvent { return MembershipRoleAssigned{} })
	events.Register(EventMembershipNicknameChanged, MembershipNicknameChangedSchemaVersion, func() events.DomainEvent { return MembershipNicknameChanged{} })
	events.Register(EventMembershipDeleted, MembershipDeletedSchemaVersion, func() events.DomainEvent { return MembershipDeleted{} })
}
