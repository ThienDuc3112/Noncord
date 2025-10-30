package entities

import (
	"backend/internal/domain/events"
	"time"

	"github.com/google/uuid"
)

type ChannelId uuid.UUID

type Channel struct {
	events.Recorder

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

func (c *Channel) UpdateName(newName string) error {
	if len(newName) == 0 {
		return NewError(ErrCodeValidationError, "channel name cannot be empty", nil)
	}
	if len(newName) > 64 {
		return NewError(ErrCodeValidationError, "channel name cannot exceed 64 characters", nil)
	}
	if c.Name != newName {
		old := c.Name
		c.Name = newName
		c.UpdatedAt = time.Now()
		c.Record(NewChannelNameUpdated(c, old))
	}
	return nil
}

func (c *Channel) UpdateDescription(newDesc string) error {
	if len(newDesc) > 256 {
		return NewError(ErrCodeValidationError, "channel description cannot exceed 256 characters", nil)
	}
	if c.Description != newDesc {
		old := c.Description
		c.Description = newDesc
		c.UpdatedAt = time.Now()
		c.Record(NewChannelDescriptionUpdated(c, old))
	}
	return nil
}

func (c *Channel) Delete() error {
	now := time.Now()
	c.DeletedAt = &now
	c.Record(NewChannelDeleted(c))
	return nil
}

func NewChannel(name, desc string, serverId ServerId, order uint16, parent *CategoryId) *Channel {
	c := &Channel{
		Id:             ChannelId(uuid.New()),
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
		DeletedAt:      nil,
		Name:           name,
		Description:    desc,
		ServerId:       serverId,
		Order:          order,
		ParentCategory: parent,
	}

	c.Record(NewChannelCreated(c))

	return c
}

type OverwriteTarget string

const (
	ChannelUserTarget OverwriteTarget = "user"
	ChannelRoleTarget OverwriteTarget = "role"
)

type ChannelPermOverwrite struct {
	ChannelId       ChannelId
	RoleId          *RoleId
	UserId          *UserId
	OverwriteTarget OverwriteTarget
	UpdatedAt       time.Time
	Allow           ServerPermissionBits
	Deny            ServerPermissionBits
}

func (p *ChannelPermOverwrite) Validate() error {
	if (p.Allow & p.Deny) != 0 {
		return NewError(ErrCodeValidationError, "cannot allow and deny the same permission", nil)
	}
	if p.OverwriteTarget != ChannelRoleTarget && p.OverwriteTarget != ChannelUserTarget {
		return NewError(ErrCodeValidationError, "invalid overwrite target", nil)
	}
	if p.OverwriteTarget == ChannelRoleTarget && (p.RoleId == nil || p.UserId != nil) {
		return NewError(ErrCodeValidationError, "overwrite target role cannot have null role id or set user id", nil)
	}
	if p.OverwriteTarget == ChannelUserTarget && (p.UserId == nil || p.RoleId != nil) {
		return NewError(ErrCodeValidationError, "overwrite target user cannot have null user id or set role id", nil)
	}

	return nil
}

func NewChannelPermOverwrite(channelId ChannelId, target OverwriteTarget, targetId uuid.UUID, allow, deny ServerPermissionBits) (*ChannelPermOverwrite, error) {
	channel := &ChannelPermOverwrite{
		ChannelId:       channelId,
		OverwriteTarget: target,
		UpdatedAt:       time.Now(),
		Allow:           allow,
		Deny:            deny,
	}
	if target == ChannelUserTarget {
		channel.UserId = (*UserId)(&targetId)
	}
	if target == ChannelRoleTarget {
		channel.RoleId = (*RoleId)(&targetId)
	}
	if err := channel.Validate(); err != nil {
		return nil, err
	}
	return channel, nil
}
