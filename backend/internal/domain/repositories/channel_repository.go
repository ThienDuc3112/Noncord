package repositories

import (
	e "backend/internal/domain/entities"
	"context"
)

type ChannelRepo interface {
	Create(ctx context.Context, channel *e.Channel) (*e.Channel, error)
	Find(ctx context.Context, id e.ChannelId) (*e.Channel, error)
	FindByServerId(ctx context.Context, serverId e.ServerId) ([]*e.Channel, error)
	FindVisibleToUser(ctx context.Context, serverId e.ServerId, userId e.UserId) ([]*e.Channel, error)
	Update(ctx context.Context, channel *e.Channel) (*e.Channel, error)
	Delete(ctx context.Context, id e.ChannelId) error
}
