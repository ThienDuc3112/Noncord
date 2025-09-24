package query

import (
	"backend/internal/application/common"

	"github.com/google/uuid"
)

type GetChannelsByServer struct {
	ServerId uuid.UUID
}

type GetChannelsByServerResult struct {
	Result []*common.Channel
}
