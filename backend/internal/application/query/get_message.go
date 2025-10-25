package query

import (
	"backend/internal/application/common"

	"github.com/google/uuid"
)

type GetMessage struct {
	MessageId uuid.UUID
	UserId    uuid.UUID
}

type GetMessageResult struct {
	Result *common.Message
}
