package entities

import (
	"time"

	"github.com/google/uuid"
)

type RoleId uuid.UUID

type Role struct {
	Id           RoleId
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    *time.Time
	Name         string
	Color        uint32
	Priority     uint8
	AllowMention bool
	Permissions  ServerPermissionBits
	ServerId     ServerId
}

type UserRole struct {
	UserId    UserId
	RoleId    RoleId
	CreatedAt time.Time
}
