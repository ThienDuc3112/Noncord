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

func NewAttachment(filetype Filetype, url, name string, mid MessageId, uid UserId, size uint32) *Attachment {
	return &Attachment{
		Id:        AttachmentId(uuid.New()),
		Filetype:  filetype,
		Url:       url,
		Filename:  name,
		MessageId: mid,
		UserId:    uid,
		Size:      size,
	}
}

type Reaction struct {
	MessageId MessageId
	UserId    UserId
	EmoteId   EmoteId
}

func NewReaction(mid MessageId, uid UserId, eid EmoteId) *Reaction {
	return &Reaction{
		MessageId: mid,
		UserId:    uid,
		EmoteId:   eid,
	}
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

func NewMessage(channelId *ChannelId, groupId *DMGroupId, authId UserId, msg string, attachments []Attachment) *Message {
	return &Message{
		Id:          MessageId(uuid.New()),
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
		DeletedAt:   nil,
		ChannelId:   channelId,
		GroupId:     groupId,
		Author:      authId,
		Message:     msg,
		Attachments: attachments,
	}
}
