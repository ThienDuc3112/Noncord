package postgres

import (
	"backend/internal/domain/entities"
	"backend/internal/domain/repositories"
	"backend/internal/infra/db/postgres/gen"
	"context"
	"fmt"
)

type PGPermissionRepo struct {
	q *gen.Queries
}

func (r *PGPermissionRepo) GetUserChannelPermission(context.Context, entities.ChannelId, entities.UserId) (repositories.UserChannelPermissionResult, error) {
	return repositories.UserChannelPermissionResult{}, fmt.Errorf("Not implemented")
}

var _ repositories.PermissionRepo = &PGPermissionRepo{}
