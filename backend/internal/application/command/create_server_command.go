package command

import (
	"backend/internal/application/common"

	"github.com/google/uuid"
)

type CreateServerCommand struct {
	UserId   uuid.UUID
	ServerId uuid.UUID
}

type CreateServerCommandResult struct {
	Result *common.Server
}
