package entities

import (
	"time"

	"github.com/google/uuid"
)

type ChannelId uuid.UUID

type Channel struct {
	Id             ChannelId
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      *time.Time
	Name           string
	Description    string
	ServerId       ServerId
	Order          uint16
	ParentCategory *CategoryId
}

func (c *Channel) Validate() error {
	if len(c.Name) > 64 {
		return NewError(ErrCodeValidationError, "name cannot exceed 64 characters", nil)
	}
	if len(c.Name) == 0 {
		return NewError(ErrCodeValidationError, "name cannot be empty", nil)
	}
	if len(c.Description) > 256 {
		return NewError(ErrCodeValidationError, "description cannot exceed 256 characters", nil)
	}
	return nil
}

type ChannelRolePermissionOverride struct {
	ChannelId ChannelId
	RoleId    RoleId
	UpdatedAt time.Time
	Allow     ServerPermissionBits
	Deny      ServerPermissionBits
}

func (p *ChannelRolePermissionOverride) Validate() error {
	if (p.Allow & p.Deny) != 0 {
		return NewError(ErrCodeValidationError, "cannot allow and deny the same permission", nil)
	}
	return nil
}

type ChannelUserPermissionOverride struct {
	ChannelId ChannelId
	UserId    UserId
	CreatedAt time.Time
	Allow     ServerPermissionBits
	Deny      ServerPermissionBits
}

func (p *ChannelUserPermissionOverride) Validate() error {
	if (p.Allow & p.Deny) != 0 {
		return NewError(ErrCodeValidationError, "cannot allow and deny the same permission", nil)
	}
	return nil
}
