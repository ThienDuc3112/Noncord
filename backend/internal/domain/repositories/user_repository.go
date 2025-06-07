package repositories

import (
	e "backend/internal/domain/entities"
	"context"
)

type UserRepo interface {
	Create(ctx context.Context, user *e.User) (*e.User, error)
	Find(ctx context.Context, id e.UserId) (*e.User, error)
	FindFriends(ctx context.Context, userId e.UserId) ([]*e.User, error)
	Update(ctx context.Context, user *e.User) (*e.User, error)
	Delete(ctx context.Context, id e.UserId) error
}
