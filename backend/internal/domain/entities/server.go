package entities

import (
	"time"

	"github.com/google/uuid"
)

type ServerId uuid.UUID

type Server struct {
	Id                  ServerId
	CreatedAt           time.Time
	UpdatedAt           time.Time
	DeletedAt           *time.Time
	Name                string
	Description         string
	IconUrl             string
	BannerUrl           string
	DefaultNotification bool
	NeedApproval        bool

	Categories []Category

	DefaultRole         RoleId
	AnnouncementChannel ChannelId
}

type CategoryId uuid.UUID

type Category struct {
	Id        CategoryId
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
	Name      string
	Order     uint8
}

type Membership struct {
	ServerId  ServerId
	UserId    UserId
	CreatedAt time.Time
}

type BanEntry struct {
	ServerId  ServerId
	UserId    UserId
	CreatedAt time.Time
}
