package postgres

import (
	e "backend/internal/domain/entities"
	"backend/internal/domain/repositories"
	"context"
	"database/sql"
)

type PGDMGroupRepo struct {
	db *sql.DB
}

func (r *PGDMGroupRepo) Find(ctx context.Context, id e.DMGroupId) (*e.DMGroup, error)
func (r *PGDMGroupRepo) FindByUserId(ctx context.Context, userId e.UserId) ([]*e.DMGroup, error)
func (r *PGDMGroupRepo) Save(ctx context.Context, group *e.DMGroup) (*e.DMGroup, error)
func (r *PGDMGroupRepo) Delete(ctx context.Context, id e.DMGroupId) error

var _ repositories.DMGroupRepo = &PGDMGroupRepo{}
