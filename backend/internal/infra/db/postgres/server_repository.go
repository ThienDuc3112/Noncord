package postgres

import (
	e "backend/internal/domain/entities"
	"backend/internal/domain/repositories"
	"backend/internal/infra/db/postgres/gen"
	"context"
)

type PGServerRepo struct {
	db gen.DBTX
}

func (r *PGServerRepo) Find(ctx context.Context, id e.ServerId) (*e.Server, error)
func (r *PGServerRepo) FindByIds(ctx context.Context, ids []e.ServerId) ([]*e.Server, error)
func (r *PGServerRepo) Save(ctx context.Context, server *e.Server) (*e.Server, error)
func (r *PGServerRepo) Delete(ctx context.Context, id e.ServerId) error

var _ repositories.ServerRepo = &PGServerRepo{}
