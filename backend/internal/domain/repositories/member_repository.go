package repositories

import (
	e "backend/internal/domain/entities"
	"context"
)

type MemberRepo interface {
	Find(ctx context.Context, userId e.UserId, serverId e.ServerId) (*e.Membership, error)
	FindByUserId(ctx context.Context, userId e.UserId) ([]*e.Membership, error)
	FindByServerId(ctx context.Context, serverId e.ServerId) ([]*e.Membership, error)

	FindRoleAssignments(ctx context.Context, userId e.UserId, serverId e.ServerId) ([]*e.RoleAssignment, error)

	Save(ctx context.Context, membership *e.Membership) (*e.Membership, error)
	SaveRoleAssignment(ctx context.Context, assignment *e.RoleAssignment) (*e.RoleAssignment, error)

	Delete(ctx context.Context, userId e.UserId, serverId e.ServerId) error
	DeleteRoleAssignment(ctx context.Context, userId e.UserId, serverId e.ServerId, roleId e.RoleId) error
}
