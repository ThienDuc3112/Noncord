package repositories

import (
	e "backend/internal/domain/entities"
	"context"
)

type ChannelRepo interface {
	Find(context.Context, e.ChannelId) (*e.Channel, error)
	FindIds(context.Context, []e.ChannelId) ([]*e.Channel, error)
	FindByServerId(context.Context, e.ServerId) ([]*e.Channel, error)

	GetServerMaxChannelOrder(context.Context, e.ServerId) (int32, error)

	Save(context.Context, *e.Channel) (*e.Channel, error)
}
