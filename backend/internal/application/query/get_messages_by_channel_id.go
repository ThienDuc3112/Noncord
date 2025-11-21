package query

import (
	"backend/internal/application/common"
	"time"

	"github.com/google/uuid"
)

type GetMessagesByChannelId struct {
	ChannelId uuid.UUID
	UserId    uuid.UUID
	Before    time.Time
	Limit     int32
}

type GetMessagesByChannelIdResult struct {
	Result []*common.Message
	More   bool
}
