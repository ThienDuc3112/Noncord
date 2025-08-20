package command

import (
	"github.com/google/uuid"
)

type UpdateServerCommand struct {
	UserId   uuid.UUID
	ServerId uuid.UUID

	Updates UpdateServerOption
}

type UpdateServerOption struct {
	Name         *string
	Description  *string
	IconUrl      *string
	BannerUrl    *string
	NeedApproval *bool

	AnnouncementChannel uuid.NullUUID
}
