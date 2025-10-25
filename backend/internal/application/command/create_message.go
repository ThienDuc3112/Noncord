package command

import (
	"backend/internal/application/common"

	"github.com/google/uuid"
)

type CreateMessageCommand struct {
	UserId    uuid.UUID
	ChannelId *uuid.UUID
	GroupId   *uuid.UUID
	Content   string
	// Attachments []Attachment
}

type CreateMessageCommandResult struct {
	Result *common.Message
}
