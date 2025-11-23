package ports

import (
	"context"

	"github.com/google/uuid"
)

type UserEnrichment struct {
	AvatarUrl string
	Nickname  string
}

type UserResolver interface {
	FromUserServer(ctx context.Context, userId, serverId uuid.UUID) (UserEnrichment, error)
	FromUserChannel(ctx context.Context, userId, channelId uuid.UUID) (UserEnrichment, error)
	FromUser(ctx context.Context, userId uuid.UUID) (UserEnrichment, error)
}
