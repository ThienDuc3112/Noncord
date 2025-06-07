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
	Order          uint8
	ParentCategory *CategoryId
}

func (c *Channel) Validate() error {
	if len(c.Name) > 64 {
		return NewError(ErrCodeValidationError, "name exceed 64 character", nil)
	}
	if len(c.Name) == 0 {
		return NewError(ErrCodeValidationError, "name cannot be empty", nil)
	}
	return nil
}

type ChannelRolePermissionOverride struct {
	UpdatedAt time.Time
	RoleId    RoleId
	ChannelId ChannelId
	Allow     ServerPermissionBits
	Deny      ServerPermissionBits
}

func (p *ChannelRolePermissionOverride) Validate() error {
	if (p.Allow & p.Deny) != 0 {
		return NewError(ErrCodeValidationError, "cannot allow and deny the same permission", nil)
	}
	return nil
}

type ChannelUserWhitelist struct {
	ChannelId ChannelId
	UserId    UserId
	CreatedAt time.Time
}

type ChannelRoleWhitelist struct {
	ChannelId ChannelId
	RoleId    RoleId
	CreatedAt time.Time
}

type DMGroupId uuid.UUID

type DMGroup struct {
	Id        DMGroupId
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
	Name      string
	IconUrl   string
	IsGroup   bool
}

func (g *DMGroup) Validate() error {
	if !g.IsGroup && g.IconUrl != "" {
		return NewError(ErrCodeValidationError, "direct message cannot set icon url", nil)
	}
	if g.IconUrl != "" && !emailReg.MatchString(g.IconUrl) {
		return NewError(ErrCodeValidationError, "invalid icon url", nil)
	}
	return nil
}

type DMGroupMember struct {
	DMGroupId DMGroupId
	Member    UserId
	CreatedAt time.Time
}
