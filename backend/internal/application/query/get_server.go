package query

import (
	"backend/internal/application/common"

	"github.com/google/uuid"
)

type GetServer struct {
	ServerId uuid.UUID
	UserId   uuid.UUID
}

type GetServerResult struct {
	Result *common.Server
}
