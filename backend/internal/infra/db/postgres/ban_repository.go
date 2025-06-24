package postgres

import (
	e "backend/internal/domain/entities"
	"backend/internal/domain/repositories"
	"backend/internal/infra/db/postgres/gen"
	"context"
	"fmt"
)

type PGBanRepo struct {
	db gen.DBTX
}

func (r *PGBanRepo) Find(ctx context.Context, serverId e.ServerId, userId e.UserId) (*e.BanEntry, error) {
	return nil, fmt.Errorf("Not implemented")
}
func (r *PGBanRepo) FindByServerId(ctx context.Context, serverId e.ServerId) ([]*e.BanEntry, error) {
	return nil, fmt.Errorf("Not implemented")
}
func (r *PGBanRepo) Save(ctx context.Context, ban *e.BanEntry) (*e.BanEntry, error) {
	return nil, fmt.Errorf("Not implemented")
}
func (r *PGBanRepo) Delete(ctx context.Context, userId e.UserId, serverId e.ServerId) error {
	return fmt.Errorf("Not implemented")
}

var _ repositories.BanRepo = &PGBanRepo{}
