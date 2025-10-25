package entities

import (
	"backend/internal/domain/events"
	"time"

	"github.com/google/uuid"
)

// ---------------------- Event names & schema versions ----------------------

const (
	EventMessageCreated           = "message.created"
	EventMessageEdited            = "message.edited"
	EventMessageAttachmentRemoved = "message.attachment_removed"
	EventMessageDeleted           = "message.deleted"
	EventMessageReactionAdded     = "message.reaction_added"
	EventMessageReactionRemoved   = "message.reaction_removed"

	MessageCreatedSchemaVersion           = 1
	MessageEditedSchemaVersion            = 1
	MessageAttachmentRemovedSchemaVersion = 1
	MessageDeletedSchemaVersion           = 1
	MessageReactionAddedSchemaVersion     = 1
	MessageReactionRemovedSchemaVersion   = 1
)

// ---------------------- Event payloads ----------------------

type AttachmentSnapshot struct {
	ID       uuid.UUID `json:"id"`
	Filetype Filetype  `json:"filetype"`
	URL      string    `json:"url,omitempty"`
	Filename string    `json:"filename"`
	Size     uint32    `json:"size"`
	UserID   uuid.UUID `json:"user_id"`
}

type MessageCreated struct {
	events.Base
	AuthorID    uuid.UUID            `json:"author_id"`
	ChannelID   *uuid.UUID           `json:"channel_id,omitempty"`
	GroupID     *uuid.UUID           `json:"group_id,omitempty"`
	Content     string               `json:"content,omitempty"`
	Attachments []AttachmentSnapshot `json:"attachments,omitempty"`
}

func NewMessageCreated(m *Message) MessageCreated {
	var snaps []AttachmentSnapshot
	if len(m.Attachments) > 0 {
		snaps = make([]AttachmentSnapshot, 0, len(m.Attachments))
		for _, a := range m.Attachments {
			snaps = append(snaps, AttachmentSnapshot{
				ID:       uuid.UUID(a.Id),
				Filetype: a.Filetype,
				URL:      a.Url,
				Filename: a.Filename,
				Size:     a.Size,
				UserID:   uuid.UUID(a.UserId),
			})
		}
	}
	return MessageCreated{
		Base:        events.NewBase("message", uuid.UUID(m.Id), EventMessageCreated, MessageCreatedSchemaVersion),
		AuthorID:    uuid.UUID(m.Author),
		ChannelID:   (*uuid.UUID)(m.ChannelId),
		GroupID:     (*uuid.UUID)(m.GroupId),
		Content:     m.Message,
		Attachments: snaps,
	}
}

type MessageEdited struct {
	events.Base
	Old string `json:"old"`
	New string `json:"new"`
}

func NewMessageEdited(m *Message, old string) MessageEdited {
	return MessageEdited{
		Base: events.NewBase("message", uuid.UUID(m.Id), EventMessageEdited, MessageEditedSchemaVersion),
		Old:  old,
		New:  m.Message,
	}
}

type MessageAttachmentRemoved struct {
	events.Base
	AttachmentID uuid.UUID `json:"attachment_id"`
}

func NewMessageAttachmentRemoved(m *Message, attID AttachmentId) MessageAttachmentRemoved {
	return MessageAttachmentRemoved{
		Base:         events.NewBase("message", uuid.UUID(m.Id), EventMessageAttachmentRemoved, MessageAttachmentRemovedSchemaVersion),
		AttachmentID: uuid.UUID(attID),
	}
}

type MessageDeleted struct {
	events.Base
	DeletedAt time.Time `json:"deleted_at"`
}

func NewMessageDeleted(m *Message) MessageDeleted {
	var deletedAt time.Time
	if m.DeletedAt != nil {
		deletedAt = *m.DeletedAt
	}
	return MessageDeleted{
		Base:      events.NewBase("message", uuid.UUID(m.Id), EventMessageDeleted, MessageDeletedSchemaVersion),
		DeletedAt: deletedAt,
	}
}

type MessageReactionAdded struct {
	events.Base
	UserID  uuid.UUID `json:"user_id"`
	EmoteID uuid.UUID `json:"emote_id"`
}

func NewMessageReactionAdded(m *Message, userID UserId, emoteID EmoteId) MessageReactionAdded {
	return MessageReactionAdded{
		Base:    events.NewBase("message", uuid.UUID(m.Id), EventMessageReactionAdded, MessageReactionAddedSchemaVersion),
		UserID:  uuid.UUID(userID),
		EmoteID: uuid.UUID(emoteID),
	}
}

type MessageReactionRemoved struct {
	events.Base
	UserID  uuid.UUID `json:"user_id"`
	EmoteID uuid.UUID `json:"emote_id"`
}

func NewMessageReactionRemoved(m *Message, userID UserId, emoteID EmoteId) MessageReactionRemoved {
	return MessageReactionRemoved{
		Base:    events.NewBase("message", uuid.UUID(m.Id), EventMessageReactionRemoved, MessageReactionRemovedSchemaVersion),
		UserID:  uuid.UUID(userID),
		EmoteID: uuid.UUID(emoteID),
	}
}

// ---------------------- Registration ----------------------

func init() {
	events.Register(EventMessageCreated, MessageCreatedSchemaVersion, func() events.DomainEvent { return &MessageCreated{} })
	events.Register(EventMessageEdited, MessageEditedSchemaVersion, func() events.DomainEvent { return &MessageEdited{} })
	events.Register(EventMessageAttachmentRemoved, MessageAttachmentRemovedSchemaVersion, func() events.DomainEvent { return &MessageAttachmentRemoved{} })
	events.Register(EventMessageDeleted, MessageDeletedSchemaVersion, func() events.DomainEvent { return &MessageDeleted{} })
	events.Register(EventMessageReactionAdded, MessageReactionAddedSchemaVersion, func() events.DomainEvent { return &MessageReactionAdded{} })
	events.Register(EventMessageReactionRemoved, MessageReactionRemovedSchemaVersion, func() events.DomainEvent { return &MessageReactionRemoved{} })
}
