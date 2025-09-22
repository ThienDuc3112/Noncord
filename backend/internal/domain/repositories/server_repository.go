package repositories

import (
	e "backend/internal/domain/entities"
	"context"
)

type ServerRepo interface {
	Find(ctx context.Context, id e.ServerId) (*e.Server, error)
	FindByIds(ctx context.Context, ids []e.ServerId) ([]*e.Server, error)
	Save(ctx context.Context, server *e.Server) (*e.Server, error)
	Delete(ctx context.Context, id e.ServerId) error
	FindByInvitationId(context.Context, e.InvitationId) (*e.Server, error)
	FindByUser(context.Context, e.UserId) ([]*e.Server, error)
}
