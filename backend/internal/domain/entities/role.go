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
	Priority     uint16
	AllowMention bool
	Permissions  ServerPermissionBits
	ServerId     ServerId
}

func (r *Role) Validate() error {
	if r.Name == "" {
		return NewError(ErrCodeValidationError, "name cannot be empty", nil)
	}
	if len(r.Name) > 64 {
		return NewError(ErrCodeValidationError, "name cannot exceed 64 characters", nil)
	}
	return nil
}

func NewRole(name string, color uint32, priority uint16, allowMention bool, perm ServerPermissionBits, sid ServerId) *Role {
	return &Role{
		Id:           RoleId(uuid.New()),
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
		DeletedAt:    nil,
		Name:         name,
		Color:        color,
		Priority:     priority,
		AllowMention: allowMention,
		Permissions:  perm,
		ServerId:     sid,
	}
}
