package entities

import "time"

type Membership struct {
	ServerId  ServerId
	UserId    UserId
	Nickname  string
	CreatedAt time.Time
}

type RoleAssignment struct {
	UserId    UserId
	ServerId  ServerId
	RoleId    RoleId
	CreatedAt time.Time
}
