package postgres

import (
	"backend/internal/domain/entities"
	"backend/internal/infra/db/postgres/gen"
)

func fromDbMessage(m gen.Message, attachments []entities.Attachment) *entities.Message {
	return &entities.Message{
		Id:          entities.MessageId(m.ID),
		CreatedAt:   m.CreatedAt,
		UpdatedAt:   m.UpdatedAt,
		DeletedAt:   m.DeletedAt,
		ChannelId:   (*entities.ChannelId)(m.ChannelID),
		GroupId:     (*entities.DMGroupId)(m.GroupID),
		Author:      entities.UserId(m.AuthorID),
		Message:     m.Message,
		Attachments: attachments,
	}
}
