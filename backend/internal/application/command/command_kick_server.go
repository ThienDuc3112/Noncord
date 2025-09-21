package command

import "github.com/google/uuid"

type KickCommand struct {
	UserId     uuid.UUID
	UserToKick uuid.UUID
}
