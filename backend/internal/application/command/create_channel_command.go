package command

import (
	"backend/internal/application/common"

	"github.com/google/uuid"
)

type CreateChannelCommand struct {
	Name           string
	Description    string
	ServerId       uuid.UUID
	ParentCategory *uuid.UUID
	UserId         uuid.UUID
}

type CreateChannelCommandResult struct {
	Result *common.Channel
}
