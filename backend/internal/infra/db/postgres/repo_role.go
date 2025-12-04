package postgres

import (
	e "backend/internal/domain/entities"
	"backend/internal/domain/repositories"
	"backend/internal/infra/db/postgres/gen"
	"context"
	"fmt"

	"github.com/google/uuid"
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
	nr, err := r.q.SaveRole(ctx, gen.SaveRoleParams{
		ID:           uuid.UUID(role.Id),
		CreatedAt:    role.CreatedAt,
		UpdatedAt:    role.UpdatedAt,
		DeletedAt:    role.DeletedAt,
		Name:         role.Name,
		Color:        int32(role.Color),
		Priority:     int16(role.Priority),
		AllowMention: role.AllowMention,
		Permissions:  int64(role.Permissions),
		ServerID:     uuid.UUID(role.ServerId),
	})
	if err != nil {
		return nil, err
	}

	if err = pullAndPushEvents(ctx, r.q, role.PullsEvents()); err != nil {
		return nil, err
	}

	return fromDbRole(nr), nil
}

func (r *PGRoleRepo) Delete(ctx context.Context, id e.RoleId) error {
	return fmt.Errorf("Not implemented")
}

var _ repositories.RoleRepo = &PGRoleRepo{}
