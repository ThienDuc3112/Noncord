package postgres

import (
	e "backend/internal/domain/entities"
	"backend/internal/domain/repositories"
	"backend/internal/infra/db/postgres/gen"
	"context"

	"github.com/google/uuid"
)

type PGUserNotiRepo struct {
	db gen.DBTX
}

func (r *PGUserNotiRepo) Find(ctx context.Context, userId e.UserId, refId uuid.UUID) (*e.UserNotification, error)
func (r *PGUserNotiRepo) Save(ctx context.Context, preference *e.UserNotification) (*e.UserNotification, error)

var _ repositories.UserNotificationRepo = &PGUserNotiRepo{}
