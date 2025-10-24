package postgres

import (
	"backend/internal/domain/entities"
	e "backend/internal/domain/entities"
	"backend/internal/domain/repositories"
	"backend/internal/infra/db/postgres/gen"
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgtype"
)

type PGUserRepo struct {
	q *gen.Queries
}

func (r *PGUserRepo) Save(ctx context.Context, user *e.User) error {
	password := pgtype.Text{}
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
		DeletedAt:   user.DeletedAt,
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
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == "23505" {
			log.Printf("[ERROR] PGSessionRepo.Save pg error: %v\n", pgErr.Detail)
			return entities.NewError(entities.ErrCodeValidationError, "username or email already in used", pgErr)
		}
	}

	return pullAndPushEvents(ctx, r.q, user.PullsEvents())
}

func (r *PGUserRepo) Find(ctx context.Context, id e.UserId) (*e.User, error) {
	u, err := r.q.FindUserById(ctx, uuid.UUID(id))
	if errors.Is(err, pgx.ErrNoRows) {
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
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, entities.NewError(entities.ErrCodeNoObject, "no user found", err)
	} else if err != nil {
		return nil, err
	}

	return fromDbUser(&u), nil
}

func (r *PGUserRepo) FindByUsername(ctx context.Context, username string) (*e.User, error) {
	u, err := r.q.FindUserByUsername(ctx, username)
	if errors.Is(err, pgx.ErrNoRows) {
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

var _ repositories.UserRepo = &PGUserRepo{}
