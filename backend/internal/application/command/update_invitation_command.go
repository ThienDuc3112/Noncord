package command

import (
	"backend/internal/application/common"

	"github.com/google/uuid"
)

type UpdateInvitationCommand struct {
	InvitationId uuid.UUID
	UserId       uuid.UUID

	Updates UpdateInvitationOption
}

type UpdateInvitationOption struct {
	BypassApproval *bool
	JoinLimit      *int32
}

type UpdateInvitationCommandResult struct {
	Result *common.Invitation
}
