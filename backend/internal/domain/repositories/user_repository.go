package repositories

import (
	e "backend/internal/domain/entities"
	"context"
)

type UserRepo interface {
	Find(ctx context.Context, id e.UserId) (*e.User, error)
	FindByIds(ctx context.Context, ids []e.UserId) ([]*e.User, error)
	FindFriends(ctx context.Context, userId e.UserId) ([]*e.User, error)
	FindByEmail(ctx context.Context, email string) (*e.User, error)
	FindByUsername(ctx context.Context, username string) (*e.User, error)
	FindManyByUsername(ctx context.Context, username string) ([]*e.User, error)

	FindSettings(ctx context.Context, userId e.UserId) (*e.UserSettings, error)

	FindFriendRequest(ctx context.Context, userId e.UserId) ([]*e.FriendRequest, error)

	Save(ctx context.Context, user *e.User) error
	SaveSettings(ctx context.Context, settings *e.UserSettings) (*e.UserSettings, error)
}
