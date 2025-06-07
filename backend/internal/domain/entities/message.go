package entities

import (
	"time"

	"github.com/google/uuid"
)

type AttachmentId uuid.UUID

type Filetype string

const (
	FiletypePNG   Filetype = "image/png"
	FiletypeJPG   Filetype = "image/jpg"
	FiletypeMP4   Filetype = "video/mp4"
	FiletypeMPEG  Filetype = "video/mpeg"
	FiletypeTXT   Filetype = "text"
	FiletypeOTHER Filetype = "other"
)

type MessageId uuid.UUID

type Attachment struct {
	Id        AttachmentId
	Filetype  Filetype
	Url       string
	MessageId MessageId
}

type EmoteId uuid.UUID

type Emote struct {
	Id       EmoteId
	Name     string
	ServerId ServerId
	IconUrl  string
}

type Reaction struct {
	MessageId MessageId
	UserId    UserId
	Emote     EmoteId
}

type Message struct {
	Id          MessageId
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   *time.Time
	ChannelId   ChannelId
	Author      UserId
	Message     string
	Attachments []Attachment
	Reactions   []Reaction
}
