package entities

import (
	"backend/internal/domain/events"
	"github.com/google/uuid"
	"time"
)

const (
	EventServerCreated                    = "server.created"
	EventServerNameUpdated                = "server.name_updated"
	EventServerDescriptionUpdated         = "server.description_updated"
	EventServerIconURLUpdated             = "server.icon_url_updated"
	EventServerBannerURLUpdated           = "server.banner_url_updated"
	EventServerNeedApprovalChanged        = "server.need_approval_changed"
	EventServerAnnouncementChannelChanged = "server.announcement_channel_changed"
	EventServerDefaultPermissionChanged   = "server.default_permission_changed"
	EventServerDeleted                    = "server.deleted"

	ServerCreatedSchemaVersion                    = 1
	ServerNameUpdatedSchemaVersion                = 1
	ServerDescriptionUpdatedSchemaVersion         = 1
	ServerIconURLUpdatedSchemaVersion             = 1
	ServerBannerURLUpdatedSchemaVersion           = 1
	ServerNeedApprovalChangedSchemaVersion        = 1
	ServerAnnouncementChannelChangedSchemaVersion = 1
	ServerDefaultPermissionChangedSchemaVersion   = 1
	ServerDeletedSchemaVersion                    = 1
)

// ----------------- Event payloads + constructors -----------------

type ServerCreated struct {
	events.Base
	Name              string    `json:"name"`
	Description       string    `json:"description"`
	OwnerID           uuid.UUID `json:"owner_id"`
	IconURL           string    `json:"icon_url,omitempty"`
	BannerURL         string    `json:"banner_url,omitempty"`
	NeedApproval      bool      `json:"need_approval"`
	DefaultPermission uint64    `json:"default_permission"`
}

func NewServerCreated(s *Server) ServerCreated {
	return ServerCreated{
		Base:              events.NewBase("server", uuid.UUID(s.Id), EventServerCreated, ServerCreatedSchemaVersion, nil, nil),
		Name:              s.Name,
		Description:       s.Description,
		OwnerID:           uuid.UUID(s.Owner),
		IconURL:           s.IconUrl,
		BannerURL:         s.BannerUrl,
		NeedApproval:      s.NeedApproval,
		DefaultPermission: uint64(s.DefaultPermission),
	}
}

type ServerNameUpdated struct {
	events.Base
	Old string `json:"old"`
	New string `json:"new"`
}

func NewServerNameUpdated(s *Server, old string) ServerNameUpdated {
	return ServerNameUpdated{
		Base: events.NewBase("server", uuid.UUID(s.Id), EventServerNameUpdated, ServerNameUpdatedSchemaVersion, nil, nil),
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
		Base: events.NewBase("server", uuid.UUID(s.Id), EventServerDescriptionUpdated, ServerDescriptionUpdatedSchemaVersion, nil, nil),
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
		Base: events.NewBase("server", uuid.UUID(s.Id), EventServerIconURLUpdated, ServerIconURLUpdatedSchemaVersion, nil, nil),
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
		Base: events.NewBase("server", uuid.UUID(s.Id), EventServerBannerURLUpdated, ServerBannerURLUpdatedSchemaVersion, nil, nil),
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
		Base: events.NewBase("server", uuid.UUID(s.Id), EventServerNeedApprovalChanged, ServerNeedApprovalChangedSchemaVersion, nil, nil),
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
		Base:         events.NewBase("server", uuid.UUID(s.Id), EventServerAnnouncementChannelChanged, ServerAnnouncementChannelChangedSchemaVersion, nil, nil),
		OldChannelID: (*uuid.UUID)(old),
		NewChannelID: (*uuid.UUID)(s.AnnouncementChannel),
	}
}

type ServerDefaultPermissionChanged struct {
	events.Base
	Old uint64 `json:"old"`
	New uint64 `json:"new"`
}

func NewServerDefaultPermissionChanged(s *Server, old ServerPermissionBits) ServerDefaultPermissionChanged {
	return ServerDefaultPermissionChanged{
		Base: events.NewBase("server", uuid.UUID(s.Id), EventServerDefaultPermissionChanged, ServerDefaultPermissionChangedSchemaVersion, nil, nil),
		Old:  uint64(old),
		New:  uint64(s.DefaultPermission),
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
		Base:      events.NewBase("server", uuid.UUID(s.Id), EventServerDeleted, ServerDeletedSchemaVersion, nil, nil),
		DeletedAt: deletedAt,
	}
}

func init() {
	events.Register(EventServerCreated, ServerCreatedSchemaVersion, func() events.DomainEvent { return &ServerCreated{} })
	events.Register(EventServerNameUpdated, ServerNameUpdatedSchemaVersion, func() events.DomainEvent { return &ServerNameUpdated{} })
	events.Register(EventServerDescriptionUpdated, ServerDescriptionUpdatedSchemaVersion, func() events.DomainEvent { return &ServerDescriptionUpdated{} })
	events.Register(EventServerIconURLUpdated, ServerIconURLUpdatedSchemaVersion, func() events.DomainEvent { return &ServerIconURLUpdated{} })
	events.Register(EventServerBannerURLUpdated, ServerBannerURLUpdatedSchemaVersion, func() events.DomainEvent { return &ServerBannerURLUpdated{} })
	events.Register(EventServerNeedApprovalChanged, ServerNeedApprovalChangedSchemaVersion, func() events.DomainEvent { return &ServerNeedApprovalChanged{} })
	events.Register(EventServerAnnouncementChannelChanged, ServerAnnouncementChannelChangedSchemaVersion, func() events.DomainEvent { return &ServerAnnouncementChannelChanged{} })
	events.Register(EventServerDefaultPermissionChanged, ServerDefaultPermissionChangedSchemaVersion, func() events.DomainEvent { return &ServerDefaultPermissionChanged{} })
	events.Register(EventServerDeleted, ServerDeletedSchemaVersion, func() events.DomainEvent { return &ServerDeleted{} })
}

