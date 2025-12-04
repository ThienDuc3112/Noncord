package postgres

import (
	"backend/internal/application/common"
	"backend/internal/domain/entities"
	"backend/internal/infra/db/postgres/gen"
)

func fromDbChannel(channel gen.Channel) *entities.Channel {
	return &entities.Channel{
		Id:             entities.ChannelId(channel.ID),
		CreatedAt:      channel.CreatedAt,
		UpdatedAt:      channel.UpdatedAt,
		DeletedAt:      channel.DeletedAt,
		Name:           channel.Name,
		Description:    channel.Description,
		ServerId:       entities.ServerId(channel.ServerID),
		Order:          uint16(channel.Ordering),
		ParentCategory: (*entities.CategoryId)(channel.ParentCategory),
	}
}

func toCommonChannel(c gen.Channel) common.Channel {
	return common.Channel{
		Id:             c.ID,
		CreatedAt:      c.CreatedAt,
		UpdatedAt:      c.UpdatedAt,
		DeletedAt:      c.DeletedAt,
		Name:           c.Name,
		Description:    c.Description,
		ServerId:       c.ServerID,
		Order:          uint16(c.Ordering),
		ParentCategory: c.ParentCategory,
	}
}
