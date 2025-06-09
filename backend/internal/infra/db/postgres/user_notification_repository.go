package postgres

import (
	e "backend/internal/domain/entities"
	"backend/internal/domain/repositories"
	"context"
	"database/sql"
)

type PGUserNotiRepo struct {
	db *sql.DB
}

func (r *PGUserNotiRepo) Find(ctx context.Context, userId e.UserId, serverId e.ServerId) (*e.UserNotification, error)
func (r *PGUserNotiRepo) Save(ctx context.Context, preference *e.UserNotification) (*e.UserNotification, error)

var _ repositories.UserNotificationRepo = &PGUserNotiRepo{}
