package query

import (
	"backend/internal/application/common"

	"github.com/google/uuid"
)

type GetInvitationsByServerId struct {
	ServerId uuid.UUID
}

type GetInvitationsByServerIdResult struct {
	Result []*common.Invitation
}
