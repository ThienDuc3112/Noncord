package command

import (
	"backend/internal/application/common"

	"github.com/google/uuid"
)

type UpdateServerCommand struct {
	UserId   uuid.UUID
	ServerId uuid.UUID

	Updates UpdateServerOption
}

type UpdateServerOption struct {
	Name                *string
	Description         *string
	IconUrl             *string
	BannerUrl           *string
	NeedApproval        *bool
	AnnouncementChannel uuid.NullUUID
}

type UpdateServerCommandResult struct {
	Result *common.Server
}

type UpsertRoleCommand struct{}

type UpsertRoleCommandResult struct{}

type DeleteRoleCommand struct{}
