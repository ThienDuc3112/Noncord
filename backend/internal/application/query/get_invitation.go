package query

import (
	"backend/internal/application/common"

	"github.com/google/uuid"
)

type GetInvitation struct {
	InvitationId uuid.UUID
}

type GetInvitationResult struct {
	Result *common.Invitation
}
