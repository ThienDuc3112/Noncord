package query

import (
	"backend/internal/application/common"
	"time"

	"github.com/google/uuid"
)

type EnrichedMessage struct {
	common.Message
	Nickname  string
	AvatarUrl string
}

type GetMessage struct {
	MessageId uuid.UUID
	UserId    uuid.UUID
}

type GetMessageResult struct {
	Result EnrichedMessage
}

type GetMessagesByChannelId struct {
	ChannelId uuid.UUID
	UserId    uuid.UUID
	Before    time.Time
	Limit     int32
}

type GetMessagesByChannelIdResult struct {
	Result []EnrichedMessage
	More   bool
}

type GetMessagesByGroupId struct {
	GroupId uuid.UUID
	UserId  uuid.UUID
	Before  time.Time
	Limit   int32
}

type GetMessagesByGroupIdResult struct {
	Result []EnrichedMessage
	More   bool
}
