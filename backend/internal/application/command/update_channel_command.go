package command

import (
	"backend/internal/application/common"

	"github.com/google/uuid"
)

type UpdateChannelCommand struct {
	UserId    uuid.UUID
	ChannelId uuid.UUID
}

type UpdateChannelCommandResult struct {
	Result *common.Channel
}
