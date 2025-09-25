package postgres

import (
	e "backend/internal/domain/entities"
	"backend/internal/domain/repositories"
	"backend/internal/infra/db/postgres/gen"
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/gookit/goutil/arrutil"
	"github.com/jackc/pgx/v5"
)

type PGServerRepo struct {
	repo *gen.Queries
}

func NewPGServerRepo(db gen.DBTX) repositories.ServerRepo {
	return &PGServerRepo{
		repo: gen.New(db),
	}
}

func (r *PGServerRepo) Save(ctx context.Context, server *e.Server) (*e.Server, error) {
	s, err := r.repo.SaveServer(ctx, gen.SaveServerParams{
		ID:                  uuid.UUID(server.Id),
		CreatedAt:           server.CreatedAt,
		UpdatedAt:           server.UpdatedAt,
		Name:                server.Name,
		Description:         server.Description,
		IconUrl:             server.IconUrl,
		BannerUrl:           server.BannerUrl,
		NeedApproval:        server.NeedApproval,
		DefaultPermssion:    int64(server.DefaultPermission),
		AnnouncementChannel: (*uuid.UUID)(server.AnnouncementChannel),
		Owner:               uuid.UUID(server.Owner),
	})
	if err != nil {
		return nil, err
	}

	return fromDbServer(s), nil
}

func (r *PGServerRepo) Find(ctx context.Context, id e.ServerId) (*e.Server, error) {
	s, err := r.repo.FindServerById(ctx, uuid.UUID(id))
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, e.NewError(e.ErrCodeNoObject, "no server by this id", err)
	} else if err != nil {
		return nil, err
	}

	return fromDbServer(s), nil
}

func (r *PGServerRepo) FindByIds(ctx context.Context, ids []e.ServerId) ([]*e.Server, error) {
	var mapper arrutil.MapFn[e.ServerId, uuid.UUID] = func(input e.ServerId) (target uuid.UUID, find bool) {
		return uuid.UUID(input), true
	}
	servers, err := r.repo.FindServersByIds(ctx, arrutil.Map(ids, mapper))
	if err != nil {
		return nil, err
	}

	return arrutil.Map(servers, func(s gen.Server) (target *e.Server, find bool) {
		return fromDbServer(s), true
	}), nil
}

func (r *PGServerRepo) Delete(ctx context.Context, id e.ServerId) error {
	return r.repo.DeleteServer(ctx, uuid.UUID(id))
}

func (r *PGServerRepo) FindByInvitationId(ctx context.Context, id e.InvitationId) (*e.Server, error) {
	server, err := r.repo.FindServerFromInviteId(ctx, uuid.UUID(id))
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, e.NewError(e.ErrCodeNoObject, "no server by this id", err)
	} else if err != nil {
		return nil, err
	}

	return fromDbServer(server), nil
}

func (r *PGServerRepo) FindByUser(ctx context.Context, userId e.UserId) ([]*e.Server, error) {
	servers, err := r.repo.FindServersFromUserId(ctx, uuid.UUID(userId))
	if err != nil {
		return nil, err
	}

	return arrutil.Map(servers, func(s gen.Server) (target *e.Server, find bool) {
		return fromDbServer(s), true
	}), nil
}

var _ repositories.ServerRepo = &PGServerRepo{}
