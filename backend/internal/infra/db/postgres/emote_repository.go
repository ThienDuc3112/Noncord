package postgres

import (
	e "backend/internal/domain/entities"
	"backend/internal/domain/repositories"
	"backend/internal/infra/db/postgres/gen"
	"context"
)

type PGEmoteRepo struct {
	db gen.DBTX
}

func (r *PGEmoteRepo) Find(ctx context.Context, id e.EmoteId) (*e.Emote, error)
func (r *PGEmoteRepo) FindByServerId(ctx context.Context, serverId e.ServerId) ([]*e.Emote, error)
func (r *PGEmoteRepo) Save(ctx context.Context, emote *e.Emote) (*e.Emote, error)
func (r *PGEmoteRepo) Delete(ctx context.Context, id e.EmoteId) error

var _ repositories.EmoteRepo = &PGEmoteRepo{}
