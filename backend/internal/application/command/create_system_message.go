package command

import (
	"github.com/google/uuid"
)

type CreateSystemMessageCommand struct {
	ServerId uuid.UUID
	Content  string
}
