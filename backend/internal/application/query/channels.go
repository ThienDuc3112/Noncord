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

type GetChannel struct {
	ChannelId uuid.UUID
	UserId    uuid.UUID
}

type GetChannelResult struct {
	Result *common.Channel
}
