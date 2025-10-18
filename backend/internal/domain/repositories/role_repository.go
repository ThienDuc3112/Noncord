package repositories

import (
	e "backend/internal/domain/entities"
	"context"
)

type RoleRepo interface {
	Find(ctx context.Context, id e.RoleId) (*e.Role, error)
	FindByServerId(ctx context.Context, serverId e.ServerId) ([]*e.Role, error)
	FindByMember(ctx context.Context, serverId e.ServerId, userId e.UserId) ([]*e.Role, error)

	Save(ctx context.Context, role *e.Role) (*e.Role, error)
}
