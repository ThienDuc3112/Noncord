package query

import (
	"backend/internal/application/common"

	"github.com/google/uuid"
)

type GetServers struct {
	ServerIds []uuid.UUID
}

type GetServersResult struct {
	Result []*common.Server
}
