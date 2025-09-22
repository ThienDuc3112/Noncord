package query

import (
	"backend/internal/application/common"

	"github.com/google/uuid"
)

type GetServersUserIn struct {
	UserId uuid.UUID
}

type GetServersUserInResult struct {
	Result []*common.Server
}
