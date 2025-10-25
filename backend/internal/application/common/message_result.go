package common

import (
	"time"

	"github.com/google/uuid"
)

type Message struct {
	Id        uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
	ChannelId *uuid.UUID
	GroupId   *uuid.UUID
	Author    uuid.UUID
	Message   string
	// Attachments []Attachment
}
