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
	EventServerDefaultRoleChanged         = "server.default_role_changed"
	EventServerDeleted                    = "server.deleted"
	// EventServerDefaultRoleChanged   = "server.default_permission_changed"

	ServerCreatedSchemaVersion                    = 1
	ServerNameUpdatedSchemaVersion                = 1
	ServerDescriptionUpdatedSchemaVersion         = 1
	ServerIconURLUpdatedSchemaVersion             = 1
	ServerBannerURLUpdatedSchemaVersion           = 1
	ServerNeedApprovalChangedSchemaVersion        = 1
	ServerAnnouncementChannelChangedSchemaVersion = 1
	ServerDefaultRoleChangedSchemaVersion         = 1
	ServerDeletedSchemaVersion                    = 1
	// ServerDefaultRoleChangedSchemaVersion   = 1
)

// ----------------- Event payloads + constructors -----------------

type ServerCreated struct {
	events.Base
	Name         string    `json:"name"`
	Description  string    `json:"description"`
	OwnerID      uuid.UUID `json:"owner_id"`
	IconURL      string    `json:"icon_url,omitempty"`
	BannerURL    string    `json:"banner_url,omitempty"`
	NeedApproval bool      `json:"need_approval"`
	// DefaultPermission uint64    `json:"default_permission"`
	DefaultRole *uuid.UUID `json:"default_role"`
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
		DefaultRole:  (*uuid.UUID)(s.DefaultRole),
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

// type ServerDefaultPermissionChanged struct {
// 	events.Base
// 	Old uint64 `json:"old"`
// 	New uint64 `json:"new"`
// }

// func NewServerDefaultPermissionChanged(s *Server, old ServerPermissionBits) ServerDefaultPermissionChanged {
// 	return ServerDefaultPermissionChanged{
// 		Base: events.NewBase("server", uuid.UUID(s.Id), EventServerDefaultPermissionChanged, ServerDefaultPermissionChangedSchemaVersion),
// 		Old:  uint64(old),
// 		New:  uint64(s.DefaultPermission),
// 	}
// }

type ServerDefaultRoleChanged struct {
	events.Base
	Old *uuid.UUID `json:"old"`
	New *uuid.UUID `json:"new"`
}

func NewServerDefaultRoleChanged(s *Server, old *RoleId) ServerDefaultRoleChanged {
	return ServerDefaultRoleChanged{
		Base: events.NewBase("server", uuid.UUID(s.Id), EventServerDefaultRoleChanged, ServerDefaultRoleChangedSchemaVersion),
		Old:  (*uuid.UUID)(old),
		New:  (*uuid.UUID)(s.DefaultRole),
	}
}

type ServerDeleted struct {
	events.Base
	DeletedAt time.Time `json:"deleted_at"`
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

func init() {
	events.Register(EventServerCreated, ServerCreatedSchemaVersion, func() events.DomainEvent { return ServerCreated{} })
	events.Register(EventServerNameUpdated, ServerNameUpdatedSchemaVersion, func() events.DomainEvent { return ServerNameUpdated{} })
	events.Register(EventServerDescriptionUpdated, ServerDescriptionUpdatedSchemaVersion, func() events.DomainEvent { return ServerDescriptionUpdated{} })
	events.Register(EventServerIconURLUpdated, ServerIconURLUpdatedSchemaVersion, func() events.DomainEvent { return ServerIconURLUpdated{} })
	events.Register(EventServerBannerURLUpdated, ServerBannerURLUpdatedSchemaVersion, func() events.DomainEvent { return ServerBannerURLUpdated{} })
	events.Register(EventServerNeedApprovalChanged, ServerNeedApprovalChangedSchemaVersion, func() events.DomainEvent { return ServerNeedApprovalChanged{} })
	events.Register(EventServerAnnouncementChannelChanged, ServerAnnouncementChannelChangedSchemaVersion, func() events.DomainEvent { return ServerAnnouncementChannelChanged{} })
	events.Register(EventServerDefaultRoleChanged, ServerDefaultRoleChangedSchemaVersion, func() events.DomainEvent { return ServerDefaultRoleChanged{} })
	events.Register(EventServerDeleted, ServerDeletedSchemaVersion, func() events.DomainEvent { return ServerDeleted{} })
}
