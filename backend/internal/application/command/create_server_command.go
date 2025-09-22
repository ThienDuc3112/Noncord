package command

import (
	"backend/internal/application/common"

	"github.com/google/uuid"
)

type CreateServerCommand struct {
	UserId          uuid.UUID
	Name            string
	UserDisplayName string
}

type CreateServerCommandResult struct {
	Result *common.Server
}
