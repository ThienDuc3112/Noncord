package postgres

import (
	e "backend/internal/domain/entities"
	"backend/internal/domain/repositories"
	"backend/internal/infra/db/postgres/gen"
	"context"
	"fmt"

	"github.com/google/uuid"
)

type PGUserNotiRepo struct {
	q *gen.Queries
}

func (r *PGUserNotiRepo) Find(ctx context.Context, userId e.UserId, refId uuid.UUID) (*e.UserNotification, error) {
	return nil, fmt.Errorf("Not implemented")
}

func (r *PGUserNotiRepo) Save(ctx context.Context, preference *e.UserNotification) (*e.UserNotification, error) {
	return nil, fmt.Errorf("Not implemented")
}

var _ repositories.UserNotificationRepo = &PGUserNotiRepo{}
