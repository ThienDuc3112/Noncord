package mapper

import (
	"backend/internal/application/common"
	"backend/internal/domain/entities"

	"github.com/google/uuid"
)

func MessageToResult(m *entities.Message) *common.Message {
	return &common.Message{
		Id:         uuid.UUID(m.Id),
		CreatedAt:  m.CreatedAt,
		UpdatedAt:  m.UpdatedAt,
		DeletedAt:  m.DeletedAt,
		ChannelId:  (*uuid.UUID)(m.ChannelId),
		GroupId:    (*uuid.UUID)(m.GroupId),
		Author:     (*uuid.UUID)(m.Author),
		AuthorType: string(m.AuthorType),
		Message:    m.Message,
	}
}
