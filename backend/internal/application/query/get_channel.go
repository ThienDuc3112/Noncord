package query

import (
	"backend/internal/application/common"

	"github.com/google/uuid"
)

type GetChannel struct {
	ChannelId uuid.UUID
}

type GetChannelResult struct {
	Result *common.Channel
}
