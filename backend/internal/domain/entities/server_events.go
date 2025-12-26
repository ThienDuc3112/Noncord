package entities

import (
	"backend/internal/domain/events"
	"time"

	"github.com/google/uuid"
)

const (
	EventServerCreated                    = "server.created"
	EventServerNameUpdated                = "server.name_updated"
	EventServerDescriptionUpdated         = "server.description_updated"
	EventServerIconURLUpdated             = "server.icon_url_updated"
	EventServerBannerURLUpdated           = "server.banner_url_updated"
	EventServerNeedApprovalChanged        = "server.need_approval_changed"
	EventServerAnnouncementChannelChanged = "server.announcement_channel_changed"
	EventServerDeleted                    = "server.deleted"
	EventRoleCreated                      = "server.role.created"
	EventRoleDeleted                      = "server.role.deleted"
	EventRoleNameUpdated                  = "server.role.name_updated"
	EventRoleColorUpdated                 = "server.role.color_updated"
	EventRoleAllowMentionChanged          = "server.role.allow_mention_changed"
	EventRolePermissionsUpdated           = "server.role.permissions_updated"

	ServerCreatedSchemaVersion                    = 1
	ServerNameUpdatedSchemaVersion                = 1
	ServerDescriptionUpdatedSchemaVersion         = 1
	ServerIconURLUpdatedSchemaVersion             = 1
	ServerBannerURLUpdatedSchemaVersion           = 1
	ServerNeedApprovalChangedSchemaVersion        = 1
	ServerAnnouncementChannelChangedSchemaVersion = 1
	ServerDeletedSchemaVersion                    = 1
	ServerRoleCreatedSchemaVersion                = 1
	ServerRoleDeletedSchemaVersion                = 1
	ServerRoleNameUpdatedSchemaVersion            = 1
	ServerRoleColorUpdatedSchemaVersion           = 1
	ServerRoleAllowMentionChangedSchemaVersion    = 1
	ServerRolePermissionsUpdatedSchemaVersion     = 1
)

// ----------------- Event payloads + constructors -----------------

type RoleSnapshot struct {
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

type ServerCreated struct {
	events.Base
	Name         string    `json:"name"`
	Description  string    `json:"description"`
	OwnerID      uuid.UUID `json:"owner_id"`
	IconURL      string    `json:"icon_url,omitempty"`
	BannerURL    string    `json:"banner_url,omitempty"`
	NeedApproval bool      `json:"need_approval"`
	DefaultRole  uuid.UUID `json:"default_role"`
}

func NewServerCreated(s *Server) ServerCreated {
	return ServerCreated{
		Base:         events.NewBase("server", uuid.UUID(s.Id), EventServerCreated, ServerCreatedSchemaVersion),
		Name:         s.Name,
		Description:  s.Description,
		OwnerID:      uuid.UUID(s.Owner),
		IconURL:      s.IconUrl,
		BannerURL:    s.BannerUrl,
		NeedApproval: s.NeedApproval,
		DefaultRole:  uuid.UUID(s.DefaultRole),
	}
}

type ServerNameUpdated struct {
	events.Base
	Old string `json:"old"`
	New string `json:"new"`
}

func NewServerNameUpdated(s *Server, old string) ServerNameUpdated {
	return ServerNameUpdated{
		Base: events.NewBase("server", uuid.UUID(s.Id), EventServerNameUpdated, ServerNameUpdatedSchemaVersion),
		Old:  old,
		New:  s.Name,
	}
}

type ServerDescriptionUpdated struct {
	events.Base
	Old string `json:"old"`
	New string `json:"new"`
}

func NewServerDescriptionUpdated(s *Server, old string) ServerDescriptionUpdated {
	return ServerDescriptionUpdated{
		Base: events.NewBase("server", uuid.UUID(s.Id), EventServerDescriptionUpdated, ServerDescriptionUpdatedSchemaVersion),
		Old:  old,
		New:  s.Description,
	}
}

type ServerIconURLUpdated struct {
	events.Base
	Old string `json:"old"`
	New string `json:"new"`
}

func NewServerIconURLUpdated(s *Server, old string) ServerIconURLUpdated {
	return ServerIconURLUpdated{
		Base: events.NewBase("server", uuid.UUID(s.Id), EventServerIconURLUpdated, ServerIconURLUpdatedSchemaVersion),
		Old:  old,
		New:  s.IconUrl,
	}
}

type ServerBannerURLUpdated struct {
	events.Base
	Old string `json:"old"`
	New string `json:"new"`
}

func NewServerBannerURLUpdated(s *Server, old string) ServerBannerURLUpdated {
	return ServerBannerURLUpdated{
		Base: events.NewBase("server", uuid.UUID(s.Id), EventServerBannerURLUpdated, ServerBannerURLUpdatedSchemaVersion),
		Old:  old,
		New:  s.BannerUrl,
	}
}

type ServerNeedApprovalChanged struct {
	events.Base
	Old bool `json:"old"`
	New bool `json:"new"`
}

func NewServerNeedApprovalChanged(s *Server, old bool) ServerNeedApprovalChanged {
	return ServerNeedApprovalChanged{
		Base: events.NewBase("server", uuid.UUID(s.Id), EventServerNeedApprovalChanged, ServerNeedApprovalChangedSchemaVersion),
		Old:  old,
		New:  s.NeedApproval,
	}
}

type ServerAnnouncementChannelChanged struct {
	events.Base
	OldChannelID *uuid.UUID `json:"old_channel_id,omitempty"`
	NewChannelID *uuid.UUID `json:"new_channel_id,omitempty"`
}

func NewServerAnnouncementChannelChanged(s *Server, old *ChannelId) ServerAnnouncementChannelChanged {
	return ServerAnnouncementChannelChanged{
		Base:         events.NewBase("server", uuid.UUID(s.Id), EventServerAnnouncementChannelChanged, ServerAnnouncementChannelChangedSchemaVersion),
		OldChannelID: (*uuid.UUID)(old),
		NewChannelID: (*uuid.UUID)(s.AnnouncementChannel),
	}
}

type ServerDeleted struct {
	events.Base
	DeletedAt time.Time `json:"deletedAt"`
}

func NewServerDeleted(s *Server) ServerDeleted {
	var deletedAt time.Time
	if s.DeletedAt != nil {
		deletedAt = *s.DeletedAt
	}
	return ServerDeleted{
		Base:      events.NewBase("server", uuid.UUID(s.Id), EventServerDeleted, ServerDeletedSchemaVersion),
		DeletedAt: deletedAt,
	}
}

type RoleCreated struct {
	events.Base

	Name         string               `json:"name"`
	Color        uint32               `json:"color"`
	Priority     uint16               `json:"priority"`
	AllowMention bool                 `json:"allowMention"`
	Permissions  ServerPermissionBits `json:"permissions"`
	ServerId     ServerId             `json:"serverId"`
}

func NewRoleCreated(r *Role) RoleCreated {
	return RoleCreated{
		Base: events.NewBase("role", uuid.UUID(r.Id), EventRoleCreated, ServerRoleCreatedSchemaVersion),

		Name:         r.Name,
		Color:        r.Color,
		Priority:     r.Priority,
		AllowMention: r.AllowMention,
		Permissions:  r.Permissions,
		ServerId:     r.ServerId,
	}
}

type RoleDeleted struct {
	events.Base
	DeletedAt time.Time `json:"deletedAt"`
}

func NewRoleDeleted(r *Role) RoleDeleted {
	var deletedAt time.Time
	if r.DeletedAt != nil {
		deletedAt = *r.DeletedAt
	}

	return RoleDeleted{
		Base: events.NewBase("role", uuid.UUID(r.Id), EventRoleDeleted, ServerRoleDeletedSchemaVersion),

		DeletedAt: deletedAt,
	}
}

type RoleNameUpdated struct {
	events.Base
	Old  string `json:"old"`
	Name string `json:"name"`
}

func NewRoleNameUpdated(r *Role, old string) RoleNameUpdated {
	return RoleNameUpdated{
		Base: events.NewBase("role", uuid.UUID(r.Id), EventRoleNameUpdated, ServerRoleNameUpdatedSchemaVersion),
		Old:  old,
		Name: r.Name,
	}
}

type RoleColorUpdated struct {
	events.Base
	Old   uint32 `json:"old"`
	Color uint32 `json:"color"`
}

func NewRoleColorUpdated(r *Role, old uint32) RoleColorUpdated {
	return RoleColorUpdated{
		Base:  events.NewBase("role", uuid.UUID(r.Id), EventRoleColorUpdated, ServerRoleColorUpdatedSchemaVersion),
		Old:   old,
		Color: r.Color,
	}
}

type RoleAllowMentionUpdated struct {
	events.Base
	Old          bool `json:"old"`
	AllowMention bool `json:"allowMention"`
}

func NewRoleAllowMentionUpdated(r *Role, old bool) RoleAllowMentionUpdated {
	return RoleAllowMentionUpdated{
		Base:         events.NewBase("role", uuid.UUID(r.Id), EventRoleAllowMentionChanged, ServerRoleAllowMentionChangedSchemaVersion),
		Old:          old,
		AllowMention: r.AllowMention,
	}
}

type RolePermissionsUpdated struct {
	events.Base
	Old         uint64 `json:"old"`
	Permissions uint64 `json:"permissions"`
}

func NewRolePermissionsUpdated(r *Role, old ServerPermissionBits) RolePermissionsUpdated {
	return RolePermissionsUpdated{
		Base:        events.NewBase("role", uuid.UUID(r.Id), EventRolePermissionsUpdated, ServerRolePermissionsUpdatedSchemaVersion),
		Old:         uint64(old),
		Permissions: uint64(r.Permissions),
	}
}

func init() {
	events.Register(EventServerCreated, ServerCreatedSchemaVersion, func() events.DomainEvent { return ServerCreated{} })
	events.Register(EventServerNameUpdated, ServerNameUpdatedSchemaVersion, func() events.DomainEvent { return ServerNameUpdated{} })
	events.Register(EventServerDescriptionUpdated, ServerDescriptionUpdatedSchemaVersion, func() events.DomainEvent { return ServerDescriptionUpdated{} })
	events.Register(EventServerIconURLUpdated, ServerIconURLUpdatedSchemaVersion, func() events.DomainEvent { return ServerIconURLUpdated{} })
	events.Register(EventServerBannerURLUpdated, ServerBannerURLUpdatedSchemaVersion, func() events.DomainEvent { return ServerBannerURLUpdated{} })
	events.Register(EventServerNeedApprovalChanged, ServerNeedApprovalChangedSchemaVersion, func() events.DomainEvent { return ServerNeedApprovalChanged{} })
	events.Register(EventServerAnnouncementChannelChanged, ServerAnnouncementChannelChangedSchemaVersion, func() events.DomainEvent { return ServerAnnouncementChannelChanged{} })
	events.Register(EventServerDeleted, ServerDeletedSchemaVersion, func() events.DomainEvent { return ServerDeleted{} })
	events.Register(EventRoleCreated, ServerRoleCreatedSchemaVersion, func() events.DomainEvent { return RoleCreated{} })
	events.Register(EventRoleDeleted, ServerRoleDeletedSchemaVersion, func() events.DomainEvent { return RoleDeleted{} })
	events.Register(EventRoleNameUpdated, ServerRoleNameUpdatedSchemaVersion, func() events.DomainEvent { return RoleNameUpdated{} })
	events.Register(EventRoleColorUpdated, ServerRoleColorUpdatedSchemaVersion, func() events.DomainEvent { return RoleColorUpdated{} })
	events.Register(EventRoleAllowMentionChanged, ServerRoleAllowMentionChangedSchemaVersion, func() events.DomainEvent { return RoleAllowMentionUpdated{} })
	events.Register(EventRolePermissionsUpdated, ServerRolePermissionsUpdatedSchemaVersion, func() events.DomainEvent { return RolePermissionsUpdated{} })
}
