package postgres

import (
	e "backend/internal/domain/entities"
	"backend/internal/domain/repositories"
	"backend/internal/infra/db/postgres/gen"
	"context"
)

type PGRoleRepo struct {
	db gen.DBTX
}

func (r *PGRoleRepo) Find(ctx context.Context, id e.RoleId) (*e.Role, error)
func (r *PGRoleRepo) FindByServerId(ctx context.Context, serverId e.ServerId) ([]*e.Role, error)
func (r *PGRoleRepo) FindByMember(ctx context.Context, serverId e.ServerId, userId e.UserId) ([]*e.Role, error)
func (r *PGRoleRepo) Save(ctx context.Context, role *e.Role) (*e.Role, error)
func (r *PGRoleRepo) Delete(ctx context.Context, id e.RoleId) error

var _ repositories.RoleRepo = &PGRoleRepo{}
