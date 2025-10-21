package entities

import (
	"backend/internal/domain/events"
	"time"

	"github.com/google/uuid"
)

type MembershipId uuid.UUID

type Membership struct {
	events.Recorder

	Id        MembershipId
	ServerId  ServerId
	UserId    UserId
	Nickname  string
	CreatedAt time.Time
	Roles     map[RoleId]bool
	deleted   bool
}

func (m *Membership) Validate() error {
	if len(m.Nickname) > 128 {
		return NewError(ErrCodeValidationError, "nickname cannot exceed 128 characters", nil)
	}
	return nil
}

func (m *Membership) IsDeleted() bool {
	return m.deleted
}

func (m *Membership) AssignRole(roleId RoleId) error {
	if m.deleted {
		return NewError(ErrCodeValidationError, "cannot assign role to a deleted membership", nil)
	}
	if m.Roles == nil {
		m.Roles = make(map[RoleId]bool)
	}
	if !m.Roles[roleId] {
		m.Roles[roleId] = true
		m.Record(NewMembershipRoleAssigned(m, roleId))
	}
	return nil
}

func (m *Membership) ChangeNickname(nickname string) error {
	if m.deleted {
		return NewError(ErrCodeValidationError, "cannot change nickname of a deleted membership", nil)
	}
	if len(nickname) > 128 {
		return NewError(ErrCodeValidationError, "nickname cannot exceed 128 characters", nil)
	}
	if m.Nickname != nickname {
		old := m.Nickname
		m.Nickname = nickname
		m.Record(NewMembershipNicknameChanged(m, old))
	}
	return nil
}

func (m *Membership) Delete() error {
	// Soft delete, idempotent
	if m.deleted {
		return nil
	}
	m.deleted = true
	m.Record(NewMembershipDeleted(m))
	return nil
}

func NewMembership(sid ServerId, uid UserId, nickname string) *Membership {
	return &Membership{
		ServerId:  sid,
		UserId:    uid,
		Nickname:  nickname,
		CreatedAt: time.Now(),
		deleted:   false,
	}
}

type RoleAssignment struct {
	MembershipId MembershipId
	RoleId       RoleId
}

func NewRoleAssignment(mid MembershipId, rid RoleId) *RoleAssignment {
	return &RoleAssignment{
		MembershipId: mid,
		RoleId:       rid,
	}
}
