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

func (r *PGServerRepo) Save(ctx context.Context, server *e.Server) (*e.Server, error) {
	if err := r.repo.DeferAllConstraint(ctx); err != nil {
		return nil, err
	}

	s, err := r.repo.SaveServer(ctx, gen.SaveServerParams{
		ID:                  uuid.UUID(server.Id),
		CreatedAt:           server.CreatedAt,
		UpdatedAt:           server.UpdatedAt,
		Name:                server.Name,
		Description:         server.Description,
		IconUrl:             server.IconUrl,
		BannerUrl:           server.BannerUrl,
		NeedApproval:        server.NeedApproval,
		DefaultRole:         (uuid.UUID)(server.DefaultRole),
		AnnouncementChannel: (*uuid.UUID)(server.AnnouncementChannel),
		Owner:               uuid.UUID(server.Owner),
	})
	if err != nil {
		return nil, err
	}

	newRoles := server.Roles
	if server.IsRoleDirty() {
		for rid, role := range server.Roles {
			if role.IsDirty() {
				newRole, err := r.repo.SaveRole(ctx, gen.SaveRoleParams{
					ID:           uuid.UUID(role.Id),
					CreatedAt:    role.CreatedAt,
					UpdatedAt:    role.UpdatedAt,
					DeletedAt:    role.DeletedAt,
					Name:         role.Name,
					Color:        int32(role.Color),
					Priority:     int16(role.Priority),
					AllowMention: role.AllowMention,
					Permissions:  int64(role.Permissions),
					ServerID:     uuid.UUID(role.ServerId),
				})
				if err != nil {
					return nil, err
				}

				newRoles[rid] = fromDbRole(newRole)
			}
		}
	}

	if err = pullAndPushEvents(ctx, r.repo, server.PullsEvents()); err != nil {
		return nil, err
	}

	return fromDbServer(s, newRoles), nil
}

func (r *PGServerRepo) Find(ctx context.Context, id e.ServerId) (*e.Server, error) {
	s, err := r.repo.FindServerById(ctx, uuid.UUID(id))
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, e.NewError(e.ErrCodeNoObject, "no server by this id", err)
	} else if err != nil {
		return nil, err
	}

	roles, err := r.repo.FindRolesByServerId(ctx, s.ID)
	if err != nil {
		return nil, err
	}
	rolesMap := make(map[e.RoleId]*e.Role)

	for _, role := range roles {
		rolesMap[e.RoleId(role.ID)] = fromDbRole(role)
	}

	return fromDbServer(s, rolesMap), nil
}

func (r *PGServerRepo) FindByIds(ctx context.Context, ids []e.ServerId) ([]*e.Server, error) {
	var mapper arrutil.MapFn[e.ServerId, uuid.UUID] = func(input e.ServerId) (target uuid.UUID, find bool) {
		return uuid.UUID(input), true
	}
	rawIds := arrutil.Map(ids, mapper)
	servers, err := r.repo.FindServersByIds(ctx, rawIds)
	if err != nil {
		return nil, err
	}

	roles, err := r.repo.FindRolesByServerIds(ctx, rawIds)
	if err != nil {
		return nil, err
	}

	rolesMap := make(map[uuid.UUID]map[e.RoleId]*e.Role)
	for _, role := range roles {
		if _, ok := rolesMap[role.ServerID]; !ok {
			rolesMap[role.ServerID] = make(map[e.RoleId]*e.Role)
		}
		rolesMap[role.ServerID][e.RoleId(role.ID)] = fromDbRole(role)
	}

	return arrutil.Map(servers, func(s gen.Server) (target *e.Server, find bool) {
		return fromDbServer(s, rolesMap[s.ID]), true
	}), nil
}

func (r *PGServerRepo) FindByInvitationId(ctx context.Context, id e.InvitationId) (*e.Server, error) {
	server, err := r.repo.FindServerFromInviteId(ctx, uuid.UUID(id))
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, e.NewError(e.ErrCodeNoObject, "no server by this id", err)
	} else if err != nil {
		return nil, err
	}

	roles, err := r.repo.FindRolesByServerId(ctx, server.ID)
	if err != nil {
		return nil, err
	}
	rolesMap := make(map[e.RoleId]*e.Role)

	for _, role := range roles {
		rolesMap[e.RoleId(role.ID)] = fromDbRole(role)
	}

	return fromDbServer(server, rolesMap), nil
}

var _ repositories.ServerRepo = &PGServerRepo{}
