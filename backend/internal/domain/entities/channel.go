package entities

import (
	"time"

	"github.com/google/uuid"
)

type ChannelId uuid.UUID

type Channel struct {
	Id          ChannelId
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   *time.Time
	Name        string
	Description string
	ServerId    ServerId
	Order       uint8
}

type ChannelRolePermissionOverride struct {
	UpdatedAt time.Time
	RoleId    RoleId
	ChannelId ChannelId
	Allow     ServerPermissionBits
	Deny      ServerPermissionBits
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

type DMGroupMember struct {
	DMGroupId DMGroupId
	Member    UserId
	CreatedAt time.Time
}
