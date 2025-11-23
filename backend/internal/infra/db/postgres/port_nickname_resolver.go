package postgres

import (
	"backend/internal/application/ports"
	"backend/internal/domain/entities"
	"backend/internal/infra/db/postgres/gen"
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type PGNicknameResolver struct {
	q *gen.Queries
}

func NewPGNicknameResolver(pool *pgxpool.Pool) ports.UserResolver {
	return &PGNicknameResolver{gen.New(pool)}
}

func (r *PGNicknameResolver) FromUserServer(ctx context.Context, userId, serverId uuid.UUID) (ports.UserEnrichment, error) {
	res, err := r.q.FindUserServerNickname(ctx, gen.FindUserServerNicknameParams{
		ID:       userId,
		ServerID: serverId,
	})
	if errors.Is(err, pgx.ErrNoRows) {
		return ports.UserEnrichment{}, entities.NewError(entities.ErrCodeNoObject, "user not found", nil)
	} else if err != nil {
		return ports.UserEnrichment{}, entities.NewError(entities.ErrCodeDepFail, "cannot get user", err)
	}

	if res.Nickname.Valid && res.Nickname.String != "" {
		return ports.UserEnrichment{
			AvatarUrl: res.AvatarUrl,
			Nickname:  res.Nickname.String,
		}, nil
	}
	return ports.UserEnrichment{
		AvatarUrl: res.AvatarUrl,
		Nickname:  res.DisplayName,
	}, nil
}

func (r *PGNicknameResolver) FromUserChannel(ctx context.Context, userId, channelId uuid.UUID) (ports.UserEnrichment, error) {
	res, err := r.q.FindUserChannelNickname(ctx, gen.FindUserChannelNicknameParams{
		ChannelID: channelId,
		UserID:    userId,
	})
	if errors.Is(err, pgx.ErrNoRows) {
		return ports.UserEnrichment{}, entities.NewError(entities.ErrCodeNoObject, "user not found", nil)
	} else if err != nil {
		return ports.UserEnrichment{}, entities.NewError(entities.ErrCodeDepFail, "cannot get user", err)
	}

	if res.Nickname.Valid && res.Nickname.String != "" {
		return ports.UserEnrichment{
			AvatarUrl: res.AvatarUrl,
			Nickname:  res.Nickname.String,
		}, nil
	}
	return ports.UserEnrichment{
		AvatarUrl: res.AvatarUrl,
		Nickname:  res.DisplayName,
	}, nil
}

func (r *PGNicknameResolver) FromUser(ctx context.Context, userId uuid.UUID) (ports.UserEnrichment, error) {
	u, err := r.q.FindUserById(ctx, userId)
	if errors.Is(err, pgx.ErrNoRows) {
		return ports.UserEnrichment{}, entities.NewError(entities.ErrCodeNoObject, "user not found", nil)
	} else if err != nil {
		return ports.UserEnrichment{}, entities.NewError(entities.ErrCodeDepFail, "cannot get user", err)
	}

	return ports.UserEnrichment{
		AvatarUrl: u.AvatarUrl,
		Nickname:  u.DisplayName,
	}, nil
}
