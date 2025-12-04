package interfaces

import (
	"backend/internal/application/query"
	"context"

	"github.com/google/uuid"
)

type VisibilityQueries interface {
	// ChannelHasAll(ctx context.Context, query query.CheckChannelPerm) (bool, error)
	// ChannelHasAny(ctx context.Context, query query.CheckChannelPerm) (bool, error)
	// ServerHasAll(ctx context.Context, query query.CheckServerPerm) (bool, error)
	// ServerHasAny(ctx context.Context, query query.CheckServerPerm) (bool, error)

	GetVisibleChannels(ctx context.Context, userId uuid.UUID) (uuid.UUIDs, error)
	GetVisibleChannelsInServer(ctx context.Context, params query.GetVisibleChannelsInServer) (uuid.UUIDs, error)
	GetVisibleServers(ctx context.Context, userId uuid.UUID) (uuid.UUIDs, error)
}
