package command

import (
	"backend/internal/application/common"
	"time"

	"github.com/google/uuid"
)

type CreateInvitationCommand struct {
	ServerId       uuid.UUID
	UserId         uuid.UUID
	ExpiresAt      *time.Time
	BypassApproval bool
	JoinLimit      int32
}

type CreateInvitationCommandResult struct {
	Result *common.Invitation
}
