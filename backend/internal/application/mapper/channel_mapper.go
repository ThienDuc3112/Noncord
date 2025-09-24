package mapper

import (
	"backend/internal/application/common"
	"backend/internal/domain/entities"

	"github.com/google/uuid"
)

func ChannelToResult(c *entities.Channel) *common.Channel {
	return &common.Channel{
		Id:             uuid.UUID(c.Id),
		CreatedAt:      c.CreatedAt,
		ServerId:       uuid.UUID(c.ServerId),
		UpdatedAt:      c.UpdatedAt,
		DeletedAt:      c.DeletedAt,
		Name:           c.Name,
		Description:    c.Description,
		Order:          c.Order,
		ParentCategory: (*uuid.UUID)(c.ParentCategory),
	}
}
