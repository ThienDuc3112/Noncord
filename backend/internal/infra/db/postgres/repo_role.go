package postgres

import (
	e "backend/internal/domain/entities"
	"backend/internal/domain/repositories"
	"backend/internal/infra/db/postgres/gen"
	"context"
	"fmt"
)

type PGRoleRepo struct {
	q *gen.Queries
}

func (r *PGRoleRepo) Find(ctx context.Context, id e.RoleId) (*e.Role, error) {
	return nil, fmt.Errorf("Not implemented")
}

func (r *PGRoleRepo) FindByServerId(ctx context.Context, serverId e.ServerId) ([]*e.Role, error) {
	return nil, fmt.Errorf("Not implemented")
}

func (r *PGRoleRepo) FindByMember(ctx context.Context, serverId e.ServerId, userId e.UserId) ([]*e.Role, error) {
	return nil, fmt.Errorf("Not implemented")
}

func (r *PGRoleRepo) Save(ctx context.Context, role *e.Role) (*e.Role, error) {
	return nil, fmt.Errorf("Not implemented")
}

func (r *PGRoleRepo) Delete(ctx context.Context, id e.RoleId) error {
	return fmt.Errorf("Not implemented")
}

var _ repositories.RoleRepo = &PGRoleRepo{}
