package command

import "github.com/google/uuid"

type LeaveServerCommand struct {
	UserId   uuid.UUID
	ServerId uuid.UUID
}
