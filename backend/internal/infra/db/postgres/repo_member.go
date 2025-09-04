package postgres

import (
	e "backend/internal/domain/entities"
	"backend/internal/domain/repositories"
	"backend/internal/infra/db/postgres/gen"
	"context"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/gookit/goutil/arrutil"
	"github.com/jackc/pgx/v5"
)

type PGMemberRepo struct {
	q *gen.Queries
}

func NewPGMemberRepo(conn gen.DBTX) repositories.MemberRepo {
	return &PGMemberRepo{
		q: gen.New(conn),
	}
}

func (r *PGMemberRepo) Find(ctx context.Context, userId e.UserId, serverId e.ServerId) (*e.Membership, error) {
	membership, err := r.q.FindMembership(ctx, gen.FindMembershipParams{
		UserID:   uuid.UUID(userId),
		ServerID: uuid.UUID(serverId),
	})
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, e.NewError(e.ErrCodeNoObject, "no user found", err)
	} else if err != nil {
		return nil, err
	}

	return fromDbMembership(membership), nil
}

func (r *PGMemberRepo) FindByUserId(ctx context.Context, userId e.UserId) ([]*e.Membership, error) {
	memberships, err := r.q.FindMembershipsByUserId(ctx, uuid.UUID(userId))
	if err != nil {
		return nil, err
	}

	return arrutil.Map(memberships, func(m gen.Membership) (target *e.Membership, find bool) {
		return fromDbMembership(m), true
	}), nil
}

func (r *PGMemberRepo) FindByServerId(ctx context.Context, serverId e.ServerId) ([]*e.Membership, error) {
	memberships, err := r.q.FindMembershipsByServerId(ctx, uuid.UUID(serverId))
	if err != nil {
		return nil, err
	}

	return arrutil.Map(memberships, func(m gen.Membership) (target *e.Membership, find bool) {
		return fromDbMembership(m), true
	}), nil
}

func (r *PGMemberRepo) FindRoleAssignments(ctx context.Context, userId e.UserId, serverId e.ServerId) ([]*e.RoleAssignment, error) {
	return nil, fmt.Errorf("Not implemented")
}

func (r *PGMemberRepo) Save(ctx context.Context, membership *e.Membership) (*e.Membership, error) {
	res, err := r.q.SaveMembership(ctx, gen.SaveMembershipParams{
		ServerID:  uuid.UUID(membership.ServerId),
		UserID:    uuid.UUID(membership.UserId),
		CreatedAt: membership.CreatedAt,
		Nickname:  membership.Nickname,
	})
	if err != nil {
		return nil, err
	}

	return fromDbMembership(res), nil
}

func (r *PGMemberRepo) SaveRoleAssignment(ctx context.Context, assignment *e.RoleAssignment) (*e.RoleAssignment, error) {
	return nil, fmt.Errorf("Not implemented")
}

func (r *PGMemberRepo) Delete(ctx context.Context, userId e.UserId, serverId e.ServerId) error {
	return r.q.DeleteMembership(ctx, gen.DeleteMembershipParams{
		UserID:   uuid.UUID(userId),
		ServerID: uuid.UUID(serverId),
	})
}

func (r *PGMemberRepo) DeleteRoleAssignment(ctx context.Context, userId e.UserId, serverId e.ServerId, roleId e.RoleId) error {
	return fmt.Errorf("Not implemented")
}

var _ repositories.MemberRepo = &PGMemberRepo{}
