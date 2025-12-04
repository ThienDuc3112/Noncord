package repositories

import (
	e "backend/internal/domain/entities"
	"context"
)

type MemberRepo interface {
	Find(ctx context.Context, userId e.UserId, serverId e.ServerId) (*e.Membership, error)
	FindByUserId(ctx context.Context, userId e.UserId) ([]*e.Membership, error)
	FindByServerId(ctx context.Context, serverId e.ServerId) ([]*e.Membership, error)

	Save(ctx context.Context, membership *e.Membership) (*e.Membership, error)
}
