package command

import (
	"backend/internal/application/common"

	"github.com/google/uuid"
)

type JoinServerCommand struct {
	UserId       uuid.UUID
	InvitationId uuid.UUID
}

type JoinServerCommandResult struct {
	Result *common.Membership
}
