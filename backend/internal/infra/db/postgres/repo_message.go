package postgres

import (
	e "backend/internal/domain/entities"
	"backend/internal/domain/repositories"
	"backend/internal/infra/db/postgres/gen"
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/gookit/goutil/arrutil"
	"github.com/jackc/pgx/v5"
)

type PGMessageRepo struct {
	q *gen.Queries
}

func (r *PGMessageRepo) Find(ctx context.Context, id e.MessageId) (*e.Message, error) {
	m, err := r.q.FindMessageById(ctx, uuid.UUID(id))
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, e.NewError(e.ErrCodeNoObject, "no message of this id exist", err)
	} else if err != nil {
		return nil, err
	}

	return fromDbMessage(m, nil), nil
}

func (r *PGMessageRepo) msgMapper(m gen.Message) (target *e.Message, find bool) {
	return fromDbMessage(m, nil), true
}

func (r *PGMessageRepo) FindByChannelId(ctx context.Context, channelId e.ChannelId, before time.Time, limit int32) ([]*e.Message, error) {
	m, err := r.q.FindMessagesByChannelId(ctx, gen.FindMessagesByChannelIdParams{
		ChannelID: (*uuid.UUID)(&channelId),
		CreatedAt: before,
		Limit:     limit,
	})
	if err != nil {
		return nil, err
	}

	return arrutil.Map(m, r.msgMapper), nil
}

func (r *PGMessageRepo) FindByGroupId(ctx context.Context, groupId e.DMGroupId, before time.Time, limit int32) ([]*e.Message, error) {
	m, err := r.q.FindMessagesByGroupId(ctx, gen.FindMessagesByGroupIdParams{
		GroupID:   (*uuid.UUID)(&groupId),
		CreatedAt: before,
		Limit:     limit,
	})
	if err != nil {
		return nil, err
	}

	return arrutil.Map(m, r.msgMapper), nil
}

func (r *PGMessageRepo) Save(ctx context.Context, msg *e.Message) (*e.Message, error) {
	m, err := r.q.SaveMessage(ctx, gen.SaveMessageParams{
		ID:        uuid.UUID(msg.Id),
		CreatedAt: msg.CreatedAt,
		UpdatedAt: msg.UpdatedAt,
		DeletedAt: msg.DeletedAt,
		ChannelID: (*uuid.UUID)(msg.ChannelId),
		GroupID:   (*uuid.UUID)(msg.GroupId),
		AuthorID:  uuid.UUID(msg.Author),
		Message:   msg.Message,
	})
	if err != nil {
		return nil, err
	}

	if err = pullAndPushEvents(ctx, r.q, msg.PullsEvents()); err != nil {
		return nil, err
	}

	return fromDbMessage(m, nil), nil
}

var _ repositories.MessageRepo = &PGMessageRepo{}
