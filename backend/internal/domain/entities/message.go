package entities

import (
	"backend/internal/domain/events"
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
	Id       AttachmentId
	Filetype Filetype
	Url      string
	Filename string
	UserId   UserId
	Size     uint32
}

func (a *Attachment) Validate() error {
	if a.Url != "" && !IsValidUrl(a.Url) {
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

func NewAttachment(filetype Filetype, url, name string, uid UserId, size uint32) *Attachment {
	return &Attachment{
		Id:       AttachmentId(uuid.New()),
		Filetype: filetype,
		Url:      url,
		Filename: name,
		UserId:   uid,
		Size:     size,
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
	events.Recorder

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
	if !noGroup && !noChannel {
		return NewError(ErrCodeValidationError, "cannot have message in both dm group and channel", nil)
	}

	return nil
}

func NewMessage(channelId *ChannelId, groupId *DMGroupId, authId UserId, msg string, attachments []Attachment) (*Message, error) {
	now := time.Now()
	message := &Message{
		Id:          MessageId(uuid.New()),
		CreatedAt:   now,
		UpdatedAt:   now,
		DeletedAt:   nil,
		ChannelId:   channelId,
		GroupId:     groupId,
		Author:      authId,
		Message:     msg,
		Attachments: attachments,
	}

	if err := message.Validate(); err != nil {
		return nil, err
	}
	message.Record(NewMessageCreated(message))
	return message, nil
}

func (m *Message) UpdateContent(newContent string) error {
	if newContent == m.Message {
		return nil
	}
	if len(newContent) > 4096 {
		return NewError(ErrCodeValidationError, "message cannot be longer than 4096", nil)
	}
	// Cannot be empty if there are no attachments
	if newContent == "" && len(m.Attachments) == 0 {
		return NewError(ErrCodeValidationError, "cannot send empty message", nil)
	}

	old := m.Message
	m.Message = newContent
	m.UpdatedAt = time.Now()
	m.Record(NewMessageEdited(m, old))
	return nil
}

func (m *Message) RemoveAttachment(attID AttachmentId) error {
	idx := -1
	var removed Attachment
	for i := range m.Attachments {
		if m.Attachments[i].Id == attID {
			idx = i
			removed = m.Attachments[i]
			break
		}
	}
	if idx < 0 {
		return NewError(ErrCodeValidationError, "attachment not found", nil)
	}

	// Remove while preserving order
	copy(m.Attachments[idx:], m.Attachments[idx+1:])
	m.Attachments = m.Attachments[:len(m.Attachments)-1]

	// Ensure message isn't empty after removal
	if m.Message == "" && len(m.Attachments) == 0 {
		return NewError(ErrCodeValidationError, "cannot remove last attachment from an empty message", nil)
	}

	m.UpdatedAt = time.Now()
	m.Record(NewMessageAttachmentRemoved(m, removed.Id))
	return nil
}

func (m *Message) AddReaction(userID UserId, emoteID EmoteId) {
	m.Record(NewMessageReactionAdded(m, userID, emoteID))
}

func (m *Message) RemoveReaction(userID UserId, emoteID EmoteId) {
	m.Record(NewMessageReactionRemoved(m, userID, emoteID))
}

func (m *Message) Delete() error {
	now := time.Now()
	m.DeletedAt = &now
	m.Record(NewMessageDeleted(m))
	return nil
}
