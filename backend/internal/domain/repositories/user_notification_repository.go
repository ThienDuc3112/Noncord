package repositories

import (
	e "backend/internal/domain/entities"
	"context"
)

type UserNotificationRepo interface {
	FindServerPreferences(ctx context.Context, userId e.UserId, serverId e.ServerId) (*e.UserServerSettingsOverride, error)
	FindChannelPreferences(ctx context.Context, userId e.UserId, channelId e.ChannelId) (*e.UserChannelSettingsOverride, error)
	FindDMGroupPreferences(ctx context.Context, userId e.UserId, groupId e.DMGroupId) (*e.UserDMSettingsOverride, error)

	SaveServerPreferences(ctx context.Context, preference *e.UserServerSettingsOverride) (*e.UserServerSettingsOverride, error)
	SaveChannelPreferences(ctx context.Context, preference *e.UserChannelSettingsOverride) (*e.UserChannelSettingsOverride, error)
	SaveDMGroupPreferences(ctx context.Context, preference *e.UserDMSettingsOverride) (*e.UserDMSettingsOverride, error)
}
