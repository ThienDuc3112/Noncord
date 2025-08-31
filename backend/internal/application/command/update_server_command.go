package command

import (
	"backend/internal/application/common"

	"github.com/google/uuid"
)

type UpdateServerCommand struct {
	UserId   uuid.UUID
	ServerId uuid.UUID
	// Roles []common.Role

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

type UpdateServerCommandResult struct {
	Result *common.Server
}
