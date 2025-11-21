package response

import (
	"time"

	"github.com/google/uuid"
)

type Message struct {
	Id         uuid.UUID  `json:"id"`
	CreatedAt  time.Time  `json:"createdAt"`
	UpdatedAt  time.Time  `json:"updatedAt"`
	ChannelId  *uuid.UUID `json:"channelId"`
	GroupId    *uuid.UUID `json:"groupId"`
	Author     *uuid.UUID `json:"author"`
	AuthorType string     `json:"authorType"`
	Message    string     `json:"message"`
}

type GetMessagesResponse struct {
	Result []Message `json:"result"`
	Next   *string   `json:"next"`
}
