package postgres

import (
	"backend/internal/application/common"
	"backend/internal/application/interfaces"
	"backend/internal/application/query"
	"backend/internal/domain/entities"
	"backend/internal/infra/db/postgres/gen"
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/gookit/goutil/arrutil"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type PGMessageQueries struct {
	q *gen.Queries
}

func NewPGMessageQueries(pool *pgxpool.Pool) interfaces.MessageQueries {
	return &PGMessageQueries{gen.New(pool)}
}

func (q *PGMessageQueries) Get(ctx context.Context, params query.GetMessage) (query.GetMessageResult, error) {
	msg, err := q.q.GetEnrichedMessageById(ctx, params.MessageId)
	if errors.Is(err, pgx.ErrNoRows) {
		return query.GetMessageResult{}, entities.NewError(entities.ErrCodeNoObject, "message not found", nil)
	} else if err != nil {
		return query.GetMessageResult{}, entities.NewError(entities.ErrCodeDepFail, "cannot get message", err)
	}

	// TODO: Check permission with roles, channel overwrite and stuff
	nickname := msg.DisplayName.String
	if msg.Nickname.Valid && msg.Nickname.String != "" {
		nickname = msg.Nickname.String
	}

	return query.GetMessageResult{Result: query.EnrichedMessage{
		Message: common.Message{
			Id:         msg.ID,
			CreatedAt:  msg.CreatedAt,
			UpdatedAt:  msg.UpdatedAt,
			DeletedAt:  msg.DeletedAt,
			ChannelId:  msg.ChannelID,
			GroupId:    msg.GroupID,
			Author:     msg.AuthorID,
			AuthorType: string(msg.AuthorType),
			Message:    msg.Message,
		},
		Nickname:  nickname,
		AvatarUrl: msg.AvatarUrl.String,
	}}, nil
}

func (q *PGMessageQueries) GetByGroupId(ctx context.Context, params query.GetMessagesByGroupId) (query.GetMessagesByGroupIdResult, error) {
	// TODO: Check permission here

	limit := int32(100)
	if params.Limit <= 500 && params.Limit >= 1 {
		limit = params.Limit
	}

	msgs, err := q.q.GetEnrichedMessageByGroupId(ctx, gen.GetEnrichedMessageByGroupIdParams{
		GroupID:   (*uuid.UUID)(&params.GroupId),
		CreatedAt: params.Before,
		Limit:     limit + 1,
	})
	if err != nil {
		return query.GetMessagesByGroupIdResult{}, entities.NewError(entities.ErrCodeDepFail, "cannot get messages", err)
	}

	parsedMsgs := arrutil.Map(msgs, func(m gen.GetEnrichedMessageByGroupIdRow) (target query.EnrichedMessage, find bool) {
		return query.EnrichedMessage{
			Message: common.Message{
				Id:         m.ID,
				CreatedAt:  m.CreatedAt,
				UpdatedAt:  m.UpdatedAt,
				DeletedAt:  m.DeletedAt,
				ChannelId:  m.ChannelID,
				GroupId:    m.GroupID,
				Author:     m.AuthorID,
				AuthorType: string(m.AuthorType),
				Message:    m.Message,
			},
			Nickname:  m.DisplayName,
			AvatarUrl: m.AvatarUrl,
		}, true
	})
	more := false
	if len(parsedMsgs) > int(limit) {
		parsedMsgs = parsedMsgs[:limit]
		more = true
	}

	return query.GetMessagesByGroupIdResult{
		Result: parsedMsgs,
		More:   more,
	}, entities.NewError(entities.ErrCodeForbidden, "dm group not implemented", nil)
}

func (q *PGMessageQueries) GetByChannelId(ctx context.Context, params query.GetMessagesByChannelId) (query.GetMessagesByChannelIdResult, error) {
	_, err := q.q.FindMembershipWithChannelId(ctx, gen.FindMembershipWithChannelIdParams{
		ID:     params.ChannelId,
		UserID: params.UserId,
	})
	if errors.Is(err, pgx.ErrNoRows) {
		return query.GetMessagesByChannelIdResult{}, entities.NewError(entities.ErrCodeForbidden, "user not in server", nil)
	} else if err != nil {
		return query.GetMessagesByChannelIdResult{}, entities.NewError(entities.ErrCodeDepFail, "cannot check user membership status", err)
	}
	// TODO: Perform checking with roles, permission overwrites, etc...

	limit := int32(100)
	if params.Limit <= 500 && params.Limit >= 1 {
		limit = params.Limit
	}

	msgs, err := q.q.GetEnrichedMessageByChannelId(ctx, gen.GetEnrichedMessageByChannelIdParams{
		ChannelID: (*uuid.UUID)(&params.ChannelId),
		CreatedAt: params.Before,
		Limit:     limit + 1,
	})
	if err != nil {
		return query.GetMessagesByChannelIdResult{}, entities.NewError(entities.ErrCodeDepFail, "cannot get messages", err)
	}

	parsedMsgs := arrutil.Map(msgs, func(m gen.GetEnrichedMessageByChannelIdRow) (target query.EnrichedMessage, find bool) {
		nickname := m.DisplayName.String
		if m.Nickname.Valid && m.Nickname.String != "" {
			nickname = m.Nickname.String
		}
		return query.EnrichedMessage{
			Message: common.Message{
				Id:         m.ID,
				CreatedAt:  m.CreatedAt,
				UpdatedAt:  m.UpdatedAt,
				DeletedAt:  m.DeletedAt,
				ChannelId:  m.ChannelID,
				GroupId:    m.GroupID,
				Author:     m.AuthorID,
				AuthorType: string(m.AuthorType),
				Message:    m.Message,
			},
			Nickname:  nickname,
			AvatarUrl: m.AvatarUrl.String,
		}, true
	})
	more := false
	if len(parsedMsgs) > int(limit) {
		parsedMsgs = parsedMsgs[:limit]
		more = true
	}

	return query.GetMessagesByChannelIdResult{
		Result: parsedMsgs,
		More:   more,
	}, nil
}
