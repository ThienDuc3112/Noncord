package repositories

import (
	e "backend/internal/domain/entities"
	"context"
)

type InvitationRepo interface {
	Find(ctx context.Context, id e.InvititationId) (*e.Invititation, error)
	FindByServerId(ctx context.Context, serverId e.ServerId) ([]*e.Invititation, error)
	Save(ctx context.Context, invite *e.Invititation) (*e.Invititation, error)
}
