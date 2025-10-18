package repositories

import (
	e "backend/internal/domain/entities"
	"context"
)

type BanRepo interface {
	Find(ctx context.Context, serverId e.ServerId, userId e.UserId) (*e.BanEntry, error)
	FindByServerId(ctx context.Context, serverId e.ServerId) ([]*e.BanEntry, error)
	Save(ctx context.Context, ban *e.BanEntry) (*e.BanEntry, error)
}
