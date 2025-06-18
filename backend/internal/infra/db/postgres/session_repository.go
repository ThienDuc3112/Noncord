package postgres

import (
	"backend/internal/domain/entities"
	"backend/internal/domain/ports"
	"backend/internal/infra/db/postgres/gen"
	"context"
	"database/sql"

	"github.com/google/uuid"
	"github.com/gookit/goutil/arrutil"
)

type PGSessionRepo struct {
	repo *gen.Queries
}

func NewPGSessionRepo(db gen.DBTX) *PGSessionRepo {
	return &PGSessionRepo{
		repo: gen.New(db),
	}
}

func (r *PGSessionRepo) Save(ctx context.Context, session *ports.Session) error {
	_, err := r.repo.CreateSession(ctx, gen.CreateSessionParams{
		ID:            session.Id,
		RotationCount: session.RotationCount,
		CreatedAt:     session.CreatedAt,
		UpdatedAt:     session.UpdatedAt,
		ExpiresAt:     session.ExpiresAt,
		UserID:        uuid.UUID(session.UserId),
		UserAgent:     session.UserAgent,
		RefreshToken:  session.Token,
	})
	return err
}

func (r *PGSessionRepo) FindById(ctx context.Context, id uuid.UUID) (*ports.Session, error) {
	session, err := r.repo.FindSessionById(ctx, id)
	if err == sql.ErrNoRows {
		return nil, entities.NewError(entities.ErrCodeNoObject, "no session by this id", err)
	} else if err != nil {
		return nil, err
	}
	return &ports.Session{
		Id:            session.ID,
		RotationCount: session.RotationCount,
		CreatedAt:     session.CreatedAt,
		UpdatedAt:     session.UpdatedAt,
		ExpiresAt:     session.ExpiresAt,
		UserId:        entities.UserId(session.UserID),
		UserAgent:     session.UserAgent,
		Token:         session.RefreshToken,
	}, nil
}

func (r *PGSessionRepo) FindByToken(ctx context.Context, token string) (*ports.Session, error) {
	session, err := r.repo.FindSessionByToken(ctx, token)
	if err == sql.ErrNoRows {
		return nil, entities.NewError(entities.ErrCodeNoObject, "no session by this id", err)
	} else if err != nil {
		return nil, err
	}
	return &ports.Session{
		Id:            session.ID,
		RotationCount: session.RotationCount,
		CreatedAt:     session.CreatedAt,
		UpdatedAt:     session.UpdatedAt,
		ExpiresAt:     session.ExpiresAt,
		UserId:        entities.UserId(session.UserID),
		UserAgent:     session.UserAgent,
		Token:         session.RefreshToken,
	}, nil
}

func (r *PGSessionRepo) FindByUserId(ctx context.Context, id entities.UserId) ([]*ports.Session, error) {
	sessions, err := r.repo.FindSessionsByUserId(ctx, uuid.UUID(id))
	if err != nil {
		return nil, err
	}
	return arrutil.Map(sessions, func(session gen.Session) (target *ports.Session, find bool) {
		return &ports.Session{
			Id:            session.ID,
			RotationCount: session.RotationCount,
			CreatedAt:     session.CreatedAt,
			UpdatedAt:     session.UpdatedAt,
			ExpiresAt:     session.ExpiresAt,
			UserId:        entities.UserId(session.UserID),
			UserAgent:     session.UserAgent,
			Token:         session.RefreshToken,
		}, true
	}), nil
}

var _ ports.SessionRepository = &PGSessionRepo{}
