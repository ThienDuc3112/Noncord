package postgres

import (
	e "backend/internal/domain/entities"
	"backend/internal/domain/repositories"
	"context"
	"database/sql"
)

type PGUserRepository struct {
	db *sql.DB
}

func (r PGUserRepository) Find(ctx context.Context, id e.UserId) (*e.User, error)
func (r PGUserRepository) FindByIds(ctx context.Context, ids []e.UserId) ([]*e.User, error)
func (r PGUserRepository) FindFriends(ctx context.Context, userId e.UserId) ([]*e.User, error)
func (r PGUserRepository) FindByUsername(ctx context.Context, username string) (*e.User, error)
func (r PGUserRepository) FindManyByUsername(ctx context.Context, username string) ([]*e.User, error)
func (r PGUserRepository) FindSettings(ctx context.Context, userId e.UserId) (*e.UserSettings, error)
func (r PGUserRepository) FindFriendRequest(ctx context.Context, userId e.UserId) ([]*e.FriendRequest, error)
func (r PGUserRepository) Save(ctx context.Context, user *e.User) (*e.User, error)
func (r PGUserRepository) SaveSettings(ctx context.Context, settings *e.UserSettings) (*e.UserSettings, error)
func (r PGUserRepository) Delete(ctx context.Context, id e.UserId) error

var _ repositories.UserRepo = &PGUserRepository{}
