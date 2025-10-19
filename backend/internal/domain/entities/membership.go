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
}

func NewMembership(sid ServerId, uid UserId, nickname string) *Membership {
	return &Membership{
		ServerId:  sid,
		UserId:    uid,
		Nickname:  nickname,
		CreatedAt: time.Now(),
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
