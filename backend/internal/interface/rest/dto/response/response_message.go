package response

import (
	"time"

	"github.com/google/uuid"
)

type Message struct {
	Id        uuid.UUID  `json:"id"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt time.Time  `json:"updatedAt"`
	ChannelId *uuid.UUID `json:"channelId"`
	GroupId   *uuid.UUID `json:"groupId"`
	Author    uuid.UUID  `json:"author"`
	Message   string     `json:"message"`
}
