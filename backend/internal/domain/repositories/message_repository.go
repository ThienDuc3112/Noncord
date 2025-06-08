package repositories

import (
	e "backend/internal/domain/entities"
	"context"
	"time"
)

type MessageRepo interface {
	Find(ctx context.Context, id e.MessageId) (*e.Message, error)
	FindByChannelId(ctx context.Context, channelId e.ChannelId, before time.Time, limit int32) ([]*e.Message, error)
	FindByGroupId(ctx context.Context, groupId e.DMGroupId, before time.Time, limit int32) ([]*e.Message, error)

	Save(ctx context.Context, msg *e.Message) (*e.Message, error)

	Delete(ctx context.Context, id e.MessageId) error
}
