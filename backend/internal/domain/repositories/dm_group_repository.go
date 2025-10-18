package repositories

import (
	e "backend/internal/domain/entities"
	"context"
)

type DMGroupRepo interface {
	Find(ctx context.Context, id e.DMGroupId) (*e.DMGroup, error)
	FindByUserId(ctx context.Context, userId e.UserId) ([]*e.DMGroup, error)

	Save(ctx context.Context, group *e.DMGroup) (*e.DMGroup, error)
}
