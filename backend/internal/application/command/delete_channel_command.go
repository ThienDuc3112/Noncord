package command

import "github.com/google/uuid"

type DeleteChannelCommand struct {
	ChannelId uuid.UUID
	UserId    uuid.UUID
}
