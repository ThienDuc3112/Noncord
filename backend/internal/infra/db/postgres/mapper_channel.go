package postgres

import (
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
