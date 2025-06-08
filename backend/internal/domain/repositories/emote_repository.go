package repositories

import (
	e "backend/internal/domain/entities"
	"context"
)

type EmoteRepo interface {
	Find(ctx context.Context, id e.EmoteId) (*e.Emote, error)
	FindByServerId(ctx context.Context, serverId e.ServerId) ([]*e.Emote, error)

	Save(ctx context.Context, emote *e.Emote) (*e.Emote, error)

	Delete(ctx context.Context, id e.EmoteId) error
}
