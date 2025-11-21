package query

import (
	"backend/internal/application/common"
	"time"

	"github.com/google/uuid"
)

type GetMessagesByGroupId struct {
	GroupId uuid.UUID
	UserId  uuid.UUID
	Before  *time.Time
	Limit   *int32
}

type GetMessagesByGroupIdResult struct {
	Result []*common.Message
	More   bool
}
