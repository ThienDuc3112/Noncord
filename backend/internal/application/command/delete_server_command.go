package command

import "github.com/google/uuid"

type DeleteServerCommand struct {
	UserId   uuid.UUID
	ServerId uuid.UUID
}
