package command

import (
	"backend/internal/application/common"

	"github.com/google/uuid"
)

type CreateMessageCommand struct {
	UserId          *uuid.UUID
	AuthorType      string
	TargetId        uuid.UUID
	Content         string
	IsTargetChannel bool
	// Attachments []Attachment
}

type CreateMessageCommandResult struct {
	Result *common.Message
}
