package postgres

import (
	e "backend/internal/domain/entities"
	"backend/internal/domain/repositories"
	"context"
	"database/sql"
)

type PGUserRepo struct {
	db *sql.DB
}

func (r *PGUserRepo) Find(ctx context.Context, id e.UserId) (*e.User, error)
func (r *PGUserRepo) FindByIds(ctx context.Context, ids []e.UserId) ([]*e.User, error)
func (r *PGUserRepo) FindFriends(ctx context.Context, userId e.UserId) ([]*e.User, error)
func (r *PGUserRepo) FindByUsername(ctx context.Context, username string) (*e.User, error)
func (r *PGUserRepo) FindManyByUsername(ctx context.Context, username string) ([]*e.User, error)
func (r *PGUserRepo) FindSettings(ctx context.Context, userId e.UserId) (*e.UserSettings, error)
func (r *PGUserRepo) FindFriendRequest(ctx context.Context, userId e.UserId) ([]*e.FriendRequest, error)
func (r *PGUserRepo) Save(ctx context.Context, user *e.User) (*e.User, error)
func (r *PGUserRepo) SaveSettings(ctx context.Context, settings *e.UserSettings) (*e.UserSettings, error)
func (r *PGUserRepo) Delete(ctx context.Context, id e.UserId) error

var _ repositories.UserRepo = &PGUserRepo{}
