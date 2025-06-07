package repositories

import (
	e "backend/internal/domain/entities"
	"context"
)

type ServerRepo interface {
	Create(ctx context.Context, server *e.Server) (*e.Server, error)
	Find(ctx context.Context, id e.ServerId) (*e.Server, error)
	FindByUserId(ctx context.Context, userId e.UserId) ([]*e.Server, error)
	Update(ctx context.Context, server *e.Server) (*e.Server, error)
	Delete(ctx context.Context, id e.ServerId) error
}
