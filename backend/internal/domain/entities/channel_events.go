package entities

import (
	"backend/internal/domain/events"
	"github.com/google/uuid"
	"time"
)

const (
	// Event names
	EventChannelCreated               = "channel.created"
	EventChannelNameUpdated           = "channel.name_updated"
	EventChannelDescriptionUpdated    = "channel.description_updated"
	EventChannelParentCategoryChanged = "channel.parent_category_changed"
	EventChannelOrderChanged          = "channel.order_changed"
	EventChannelDeleted               = "channel.deleted"

	// Overwrite events
	EventChannelOverwriteUpserted = "channel.overwrite_upserted"
	EventChannelOverwriteDeleted  = "channel.overwrite_deleted"

	// Schema versions
	ChannelCreatedSchemaVersion               = 1
	ChannelNameUpdatedSchemaVersion           = 1
	ChannelDescriptionUpdatedSchemaVersion    = 1
	ChannelParentCategoryChangedSchemaVersion = 1
	ChannelOrderChangedSchemaVersion          = 1
	ChannelDeletedSchemaVersion               = 1

	ChannelOverwriteUpsertedSchemaVersion = 1
	ChannelOverwriteDeletedSchemaVersion  = 1
)

// ----------------- Event payloads + constructors -----------------

type ChannelCreated struct {
	events.Base
	ServerID       uuid.UUID  `json:"server_id"`
	Name           string     `json:"name"`
	Description    string     `json:"description"`
	Order          uint16     `json:"order"`
	ParentCategory *uuid.UUID `json:"parent_category_id,omitempty"`
	CreatedAt      time.Time  `json:"created_at"`
}

func NewChannelCreated(c *Channel) ChannelCreated {
	return ChannelCreated{
		Base:           events.NewBase("channel", uuid.UUID(c.Id), EventChannelCreated, ChannelCreatedSchemaVersion),
		ServerID:       uuid.UUID(c.ServerId),
		Name:           c.Name,
		Description:    c.Description,
		Order:          c.Order,
		ParentCategory: (*uuid.UUID)(c.ParentCategory),
		CreatedAt:      c.CreatedAt,
	}
}

type ChannelNameUpdated struct {
	events.Base
	Old string `json:"old"`
	New string `json:"new"`
}

func NewChannelNameUpdated(c *Channel, old string) ChannelNameUpdated {
	return ChannelNameUpdated{
		Base: events.NewBase("channel", uuid.UUID(c.Id), EventChannelNameUpdated, ChannelNameUpdatedSchemaVersion),
		Old:  old,
		New:  c.Name,
	}
}

type ChannelDescriptionUpdated struct {
	events.Base
	Old string `json:"old"`
	New string `json:"new"`
}

func NewChannelDescriptionUpdated(c *Channel, old string) ChannelDescriptionUpdated {
	return ChannelDescriptionUpdated{
		Base: events.NewBase("channel", uuid.UUID(c.Id), EventChannelDescriptionUpdated, ChannelDescriptionUpdatedSchemaVersion),
		Old:  old,
		New:  c.Description,
	}
}

type ChannelParentCategoryChanged struct {
	events.Base
	OldParentCategoryID *uuid.UUID `json:"old_parent_category_id,omitempty"`
	NewParentCategoryID *uuid.UUID `json:"new_parent_category_id,omitempty"`
}

func NewChannelParentCategoryChanged(c *Channel, old *CategoryId) ChannelParentCategoryChanged {
	return ChannelParentCategoryChanged{
		Base:                events.NewBase("channel", uuid.UUID(c.Id), EventChannelParentCategoryChanged, ChannelParentCategoryChangedSchemaVersion),
		OldParentCategoryID: (*uuid.UUID)(old),
		NewParentCategoryID: (*uuid.UUID)(c.ParentCategory),
	}
}

type ChannelOrderChanged struct {
	events.Base
	Old uint16 `json:"old"`
	New uint16 `json:"new"`
}

func NewChannelOrderChanged(c *Channel, old uint16) ChannelOrderChanged {
	return ChannelOrderChanged{
		Base: events.NewBase("channel", uuid.UUID(c.Id), EventChannelOrderChanged, ChannelOrderChangedSchemaVersion),
		Old:  old,
		New:  c.Order,
	}
}

type ChannelDeleted struct {
	events.Base
	DeletedAt time.Time `json:"deleted_at"`
}

func NewChannelDeleted(c *Channel) ChannelDeleted {
	var deletedAt time.Time
	if c.DeletedAt != nil {
		deletedAt = *c.DeletedAt
	}
	return ChannelDeleted{
		Base:      events.NewBase("channel", uuid.UUID(c.Id), EventChannelDeleted, ChannelDeletedSchemaVersion),
		DeletedAt: deletedAt,
	}
}

// ----------------- Permission overwrite events -----------------

type ChannelOverwriteUpserted struct {
	events.Base
	ChannelID       uuid.UUID       `json:"channel_id"`
	OverwriteTarget OverwriteTarget `json:"overwrite_target"` // "user" | "role"
	UserID          *uuid.UUID      `json:"user_id,omitempty"`
	RoleID          *uuid.UUID      `json:"role_id,omitempty"`
	Allow           uint64          `json:"allow"`
	Deny            uint64          `json:"deny"`
	UpdatedAt       time.Time       `json:"updated_at"`
}

func NewChannelOverwriteUpserted(po *ChannelPermOverwrite) ChannelOverwriteUpserted {
	return ChannelOverwriteUpserted{
		Base:            events.NewBase("channel", uuid.UUID(po.ChannelId), EventChannelOverwriteUpserted, ChannelOverwriteUpsertedSchemaVersion),
		ChannelID:       uuid.UUID(po.ChannelId),
		OverwriteTarget: po.OverwriteTarget,
		UserID:          (*uuid.UUID)(po.UserId),
		RoleID:          (*uuid.UUID)(po.RoleId),
		Allow:           uint64(po.Allow),
		Deny:            uint64(po.Deny),
		UpdatedAt:       po.UpdatedAt,
	}
}

type ChannelOverwriteDeleted struct {
	events.Base
	ChannelID       uuid.UUID       `json:"channel_id"`
	OverwriteTarget OverwriteTarget `json:"overwrite_target"` // "user" | "role"
	UserID          *uuid.UUID      `json:"user_id,omitempty"`
	RoleID          *uuid.UUID      `json:"role_id,omitempty"`
	DeletedAt       time.Time       `json:"deleted_at"`
}

func NewChannelOverwriteDeleted(channelId ChannelId, target OverwriteTarget, userId *UserId, roleId *RoleId) ChannelOverwriteDeleted {
	return ChannelOverwriteDeleted{
		Base:            events.NewBase("channel", uuid.UUID(channelId), EventChannelOverwriteDeleted, ChannelOverwriteDeletedSchemaVersion),
		ChannelID:       uuid.UUID(channelId),
		OverwriteTarget: target,
		UserID:          (*uuid.UUID)(userId),
		RoleID:          (*uuid.UUID)(roleId),
		DeletedAt:       time.Now(),
	}
}

// ----------------- Registration -----------------

func init() {
	events.Register(EventChannelCreated, ChannelCreatedSchemaVersion, func() events.DomainEvent { return &ChannelCreated{} })
	events.Register(EventChannelNameUpdated, ChannelNameUpdatedSchemaVersion, func() events.DomainEvent { return &ChannelNameUpdated{} })
	events.Register(EventChannelDescriptionUpdated, ChannelDescriptionUpdatedSchemaVersion, func() events.DomainEvent { return &ChannelDescriptionUpdated{} })
	events.Register(EventChannelParentCategoryChanged, ChannelParentCategoryChangedSchemaVersion, func() events.DomainEvent { return &ChannelParentCategoryChanged{} })
	events.Register(EventChannelOrderChanged, ChannelOrderChangedSchemaVersion, func() events.DomainEvent { return &ChannelOrderChanged{} })
	events.Register(EventChannelDeleted, ChannelDeletedSchemaVersion, func() events.DomainEvent { return &ChannelDeleted{} })

	events.Register(EventChannelOverwriteUpserted, ChannelOverwriteUpsertedSchemaVersion, func() events.DomainEvent { return &ChannelOverwriteUpserted{} })
	events.Register(EventChannelOverwriteDeleted, ChannelOverwriteDeletedSchemaVersion, func() events.DomainEvent { return &ChannelOverwriteDeleted{} })
}
