package repositories

import (
	e "backend/internal/domain/entities"
	"context"

	"github.com/google/uuid"
)

type UserNotificationRepo interface {
	Find(ctx context.Context, userId e.UserId, refId uuid.UUID) (*e.UserNotification, error)
	Save(ctx context.Context, preference *e.UserNotification) (*e.UserNotification, error)
}
