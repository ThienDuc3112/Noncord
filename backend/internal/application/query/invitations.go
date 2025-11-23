package query

import (
	"backend/internal/application/common"

	"github.com/google/uuid"
)

type GetInvitationsByServerId struct {
	ServerId uuid.UUID
	UserId   uuid.UUID
}

type GetInvitationsByServerIdResult struct {
	Result []*common.Invitation
}

type GetInvitation struct {
	InvitationId uuid.UUID
}

type GetInvitationResult struct {
	Result *common.Invitation
}
