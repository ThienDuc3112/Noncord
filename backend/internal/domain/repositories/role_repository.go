package repositories

import (
	e "backend/internal/domain/entities"
	"context"
)

type RoleRepo interface {
	Create(ctx context.Context, role *e.Role) (*e.Role, error)
	Find(ctx context.Context, id e.RoleId) (*e.Role, error)
	FindByServerId(ctx context.Context, serverId e.ServerId) ([]*e.Role, error)
	Update(ctx context.Context, role *e.Role) (*e.Role, error)
	Delete(ctx context.Context, id e.RoleId) error
}
