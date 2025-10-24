package postgres

import (
	e "backend/internal/domain/entities"
	"backend/internal/domain/repositories"
	"backend/internal/infra/db/postgres/gen"
	"context"
	"fmt"
)

type PGDMGroupRepo struct {
	q *gen.Queries
}

func (r *PGDMGroupRepo) Find(ctx context.Context, id e.DMGroupId) (*e.DMGroup, error) {
	return nil, fmt.Errorf("Not implemented")
}

func (r *PGDMGroupRepo) FindByUserId(ctx context.Context, userId e.UserId) ([]*e.DMGroup, error) {
	return nil, fmt.Errorf("Not implemented")
}

func (r *PGDMGroupRepo) Save(ctx context.Context, group *e.DMGroup) (*e.DMGroup, error) {
	return nil, fmt.Errorf("Not implemented")
}

func (r *PGDMGroupRepo) Delete(ctx context.Context, id e.DMGroupId) error {
	return fmt.Errorf("Not implemented")
}

var _ repositories.DMGroupRepo = &PGDMGroupRepo{}
