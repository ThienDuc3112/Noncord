package postgres

import (
	"backend/internal/application/common"
	"backend/internal/application/interfaces"
	"backend/internal/domain/entities"
	"backend/internal/infra/db/postgres/gen"
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type PGUserQueries struct {
	q *gen.Queries
}

func NewPGUserQueries(pool *pgxpool.Pool) interfaces.UserQueries {
	return &PGUserQueries{gen.New(pool)}
}

func (q *PGUserQueries) GetBasic(ctx context.Context, id uuid.UUID) (common.UserResult, error) {
	u, err := q.q.FindUserById(ctx, id)
	if errors.Is(err, pgx.ErrNoRows) {
		return common.UserResult{}, entities.NewError(entities.ErrCodeNoObject, "user not found", nil)
	} else if err != nil {
		return common.UserResult{}, entities.NewError(entities.ErrCodeDepFail, "cannot get user", err)
	}

	return toCommonUser(u), nil
}
