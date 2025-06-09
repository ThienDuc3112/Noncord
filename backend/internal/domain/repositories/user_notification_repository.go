package repositories

import (
	e "backend/internal/domain/entities"
	"context"
)

type UserNotificationRepo interface {
	Find(ctx context.Context, userId e.UserId, serverId e.ServerId) (*e.UserNotification, error)
	Save(ctx context.Context, preference *e.UserNotification) (*e.UserNotification, error)
}
