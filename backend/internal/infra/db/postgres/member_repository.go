package postgres

import (
	e "backend/internal/domain/entities"
	"backend/internal/domain/repositories"
	"backend/internal/infra/db/postgres/gen"
	"context"
)

type PGMemberRepo struct {
	db gen.DBTX
}

func (r *PGMemberRepo) Find(ctx context.Context, userId e.UserId, serverId e.ServerId) (*e.Membership, error)
func (r *PGMemberRepo) FindByUserId(ctx context.Context, userId e.UserId) ([]*e.Membership, error)
func (r *PGMemberRepo) FindByServerId(ctx context.Context, serverId e.ServerId) ([]*e.Membership, error)
func (r *PGMemberRepo) FindRoleAssignments(ctx context.Context, userId e.UserId, serverId e.ServerId) ([]*e.RoleAssignment, error)
func (r *PGMemberRepo) Save(ctx context.Context, membership *e.Membership) (*e.Membership, error)
func (r *PGMemberRepo) SaveRoleAssignment(ctx context.Context, assignment *e.RoleAssignment) (*e.RoleAssignment, error)
func (r *PGMemberRepo) Delete(ctx context.Context, userId e.UserId, serverId e.ServerId) error
func (r *PGMemberRepo) DeleteRoleAssignment(ctx context.Context, userId e.UserId, serverId e.ServerId, roleId e.RoleId) error

var _ repositories.MemberRepo = &PGMemberRepo{}
