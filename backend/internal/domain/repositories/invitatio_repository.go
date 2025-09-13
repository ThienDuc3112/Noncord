package repositories

import (
	e "backend/internal/domain/entities"
	"context"
)

type InvitationRepo interface {
	Find(ctx context.Context, id e.InvitationId) (*e.Invitation, error)
	FindByServerId(ctx context.Context, serverId e.ServerId) ([]*e.Invitation, error)
	Save(ctx context.Context, invite *e.Invitation) (*e.Invitation, error)
}
