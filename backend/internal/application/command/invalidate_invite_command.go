package command

import (
	"github.com/google/uuid"
)

type InvalidateInvitationCommand struct {
	InvitationId uuid.UUID
	UserId       uuid.UUID
}
