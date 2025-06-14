package entities

import "time"

type Membership struct {
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
	UserId    UserId
	ServerId  ServerId
	RoleId    RoleId
	CreatedAt time.Time
}

func NewRoleAssignment(sid ServerId, uid UserId, rid RoleId) *RoleAssignment {
	return &RoleAssignment{
		UserId:    uid,
		ServerId:  sid,
		RoleId:    rid,
		CreatedAt: time.Now(),
	}
}
