package postgres

import (
	"backend/internal/domain/entities"
	e "backend/internal/domain/entities"
	"backend/internal/domain/repositories"
	"backend/internal/infra/db/postgres/gen"
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/google/uuid"
)

type PGUserRepo struct {
	q *gen.Queries
}

func NewPGUserRepo(conn gen.DBTX) repositories.UserRepo {
	return &PGUserRepo{
		q: gen.New(conn),
	}
}

func (r *PGUserRepo) Save(ctx context.Context, user *e.User) error {
	deletedAt := sql.NullTime{}
	if user.DeletedAt == nil {
		deletedAt.Valid = false
	} else {
		deletedAt.Time = *user.DeletedAt
		deletedAt.Valid = true
	}

	password := sql.NullString{}
	if len(user.Password) == 0 {
		password.Valid = false
	} else {
		password.String = user.Password
		password.Valid = true
	}

	_, err := r.q.CreateUser(ctx, gen.CreateUserParams{
		ID:          uuid.UUID(user.Id),
		CreatedAt:   user.CreatedAt,
		UpdatedAt:   user.UpdatedAt,
		DeletedAt:   deletedAt,
		Username:    user.Username,
		DisplayName: user.DisplayName,
		AboutMe:     user.AboutMe,
		Email:       user.Email,
		Password:    password,
		Disabled:    user.Disabled,
		AvatarUrl:   user.AvatarUrl,
		BannerUrl:   user.BannerUrl,
		Flags:       int16(user.Flags),
	})

	return err
}

func (r *PGUserRepo) Find(ctx context.Context, id e.UserId) (*e.User, error) {
	u, err := r.q.FindUserById(ctx, uuid.UUID(id))
	if errors.Is(err, sql.ErrNoRows) {
		return nil, entities.NewError(entities.ErrCodeNoObject, "no user found", err)
	} else if err != nil {
		return nil, err
	}

	return fromDbUser(&u), nil
}

func (r *PGUserRepo) FindByIds(ctx context.Context, ids []e.UserId) ([]*e.User, error) {
	return nil, fmt.Errorf("Not implemented")
}

func (r *PGUserRepo) FindFriends(ctx context.Context, userId e.UserId) ([]*e.User, error) {
	return nil, fmt.Errorf("Not implemented")
}

func (r *PGUserRepo) FindByEmail(ctx context.Context, email string) (*e.User, error) {
	u, err := r.q.FindUserByEmail(ctx, email)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, entities.NewError(entities.ErrCodeNoObject, "no user found", err)
	} else if err != nil {
		return nil, err
	}

	return fromDbUser(&u), nil
}

func (r *PGUserRepo) FindByUsername(ctx context.Context, username string) (*e.User, error) {
	u, err := r.q.FindUserByUsername(ctx, username)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, entities.NewError(entities.ErrCodeNoObject, "no user found", err)
	} else if err != nil {
		return nil, err
	}

	return fromDbUser(&u), nil
}

func (r *PGUserRepo) FindManyByUsername(ctx context.Context, username string) ([]*e.User, error) {
	return nil, fmt.Errorf("Not implemented")
}

func (r *PGUserRepo) FindSettings(ctx context.Context, userId e.UserId) (*e.UserSettings, error) {
	return nil, fmt.Errorf("Not implemented")
}

func (r *PGUserRepo) FindFriendRequest(ctx context.Context, userId e.UserId) ([]*e.FriendRequest, error) {
	return nil, fmt.Errorf("Not implemented")
}

func (r *PGUserRepo) SaveSettings(ctx context.Context, settings *e.UserSettings) (*e.UserSettings, error) {
	return nil, fmt.Errorf("Not implemented")
}

func (r *PGUserRepo) Delete(ctx context.Context, id e.UserId) error {
	return fmt.Errorf("Not implemented")
}

func (r *PGUserRepo) WithTx(tx repositories.DBTX) repositories.UserRepo {
	return &PGUserRepo{
		q: gen.New(tx),
	}
}

var _ repositories.UserRepo = &PGUserRepo{}
