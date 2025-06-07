package repositories

import (
	e "backend/internal/domain/entities"
	"context"
)

type DMGroupRepo interface {
	Create(ctx context.Context, group *e.DMGroup) (*e.DMGroup, error)
	Find(ctx context.Context, id e.DMGroupId) (*e.DMGroup, error)
	FindByUserId(ctx context.Context, userId e.UserId) ([]*e.DMGroup, error)
	Update(ctx context.Context, group *e.DMGroup) (*e.DMGroup, error)
	Delete(ctx context.Context, id e.DMGroupId) error
}
