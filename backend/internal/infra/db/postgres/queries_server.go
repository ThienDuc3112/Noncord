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

type PGServerQueries struct {
	q *gen.Queries
}

func NewPGServerQueries(pool *pgxpool.Pool) interfaces.ServerQueries {
	return &PGServerQueries{gen.New(pool)}
}

func (q *PGServerQueries) Get(ctx context.Context, p query.GetServer) (query.GetServerResult, error) {
	s, err := q.q.FindServerById(ctx, p.ServerId)
	if errors.Is(err, pgx.ErrNoRows) {
		return query.GetServerResult{}, entities.NewError(entities.ErrCodeNoObject, "server not found", nil)
	} else if err != nil {
		return query.GetServerResult{}, entities.NewError(entities.ErrCodeDepFail, "cannot get server", err)
	}

	// TODO: use the common logic to filter out which channel can and cannot be show
	cs, err := q.q.FindChannelsByServerId(ctx, p.ServerId)
	if err != nil {
		return query.GetServerResult{}, entities.NewError(entities.ErrCodeDepFail, "cannot get channels", err)
	}

	roles, err := q.q.FindRolesByServerId(ctx, p.ServerId)
	if err != nil {
		return query.GetServerResult{}, entities.NewError(entities.ErrCodeDepFail, "cannot get roles", err)
	}

	rs := toCommonServer(s)

	return query.GetServerResult{
		Preview: common.ServerPreview{
			Id:          s.ID,
			Name:        s.Name,
			IconUrl:     s.IconUrl,
			BannerUrl:   s.BannerUrl,
			Description: s.Description,
		},
		Full: &rs,
		Channel: arrutil.Map(cs, func(c gen.Channel) (target common.Channel, find bool) {
			return toCommonChannel(c), true
		}),
		Roles: arrutil.Map(roles, func(r gen.Role) (target common.Role, find bool) {
			return toCommonRole(r), true
		}),
	}, nil
}

func (q *PGServerQueries) GetServers(ctx context.Context, p query.GetServers) (query.GetServersResult, error) {
	servers, err := q.q.FindServersByIds(ctx, p.ServerIds)
	if err != nil {
		return query.GetServersResult{}, entities.NewError(entities.ErrCodeDepFail, "cannot get servers", err)
	}
	return query.GetServersResult{
		Result: arrutil.Map(servers, func(s gen.Server) (common.Server, bool) {
			return toCommonServer(s), true
		}),
	}, nil
}

func (q *PGServerQueries) GetServersUserIn(ctx context.Context, p query.GetServersUserIn) (query.GetServersUserInResult, error) {
	servers, err := q.q.FindServersFromUserId(ctx, p.UserId)
	if err != nil {
		return query.GetServersUserInResult{}, entities.NewError(entities.ErrCodeDepFail, "cannot get servers", err)
	}
	return query.GetServersUserInResult{
		Result: arrutil.Map(servers, func(s gen.Server) (common.Server, bool) {
			return toCommonServer(s), true
		}),
	}, nil
}

func (q *PGServerQueries) GetServerIdsUserIn(ctx context.Context, p query.GetServersUserIn) (uuid.UUIDs, error) {
	servers, err := q.q.FindServersFromUserId(ctx, p.UserId)
	if err != nil {
		return nil, entities.NewError(entities.ErrCodeDepFail, "cannot get servers", err)
	}
	return arrutil.Map(servers, func(s gen.Server) (uuid.UUID, bool) { return s.ID, true }), nil
}
