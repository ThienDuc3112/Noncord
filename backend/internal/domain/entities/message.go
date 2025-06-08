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
	FiletypeGIF   Filetype = "image/gif"
	FiletypeWEBP  Filetype = "image/webp"
	FiletypeMP4   Filetype = "video/mp4"
	FiletypeTXT   Filetype = "text"
	FiletypeOTHER Filetype = "other"
)

type MessageId uuid.UUID

type Attachment struct {
	Id        AttachmentId
	Filetype  Filetype
	Url       string
	Filename  string
	MessageId MessageId
	UserId    UserId
	Size      uint32
}

func (a *Attachment) Validate() error {
	if a.Url != "" && !emailReg.MatchString(a.Url) {
		return NewError(ErrCodeValidationError, "invalid attachment url", nil)
	}
	if len(a.Url) > 2048 {
		return NewError(ErrCodeValidationError, "url too long", nil)
	}
	if len(a.Filename) > 128 {
		return NewError(ErrCodeValidationError, "file name too long", nil)
	}

	return nil
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
	ChannelId   *ChannelId
	GroupId     *DMGroupId
	Author      UserId
	Message     string
	Attachments []Attachment
}

func (m *Message) Validate() error {
	if m.Message == "" && len(m.Attachments) == 0 {
		return NewError(ErrCodeValidationError, "cannot send empty message", nil)
	}
	if len(m.Message) > 4096 {
		return NewError(ErrCodeValidationError, "message cannot be longer than 4096", nil)
	}
	if len(m.Attachments) > 10 {
		return NewError(ErrCodeValidationError, "attachments limit exceed", nil)
	}
	noChannel := m.ChannelId == nil
	noGroup := m.GroupId == nil
	if noGroup && noChannel {
		return NewError(ErrCodeValidationError, "cannot have orphan message", nil)
	}

	return nil
}
