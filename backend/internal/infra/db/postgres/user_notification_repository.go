package postgres

import (
	e "backend/internal/domain/entities"
	"backend/internal/domain/repositories"
	"context"
	"database/sql"
)

type PGUserNotiRepo struct {
	db *sql.DB
}

func (r *PGUserNotiRepo) FindServerPreferences(ctx context.Context, userId e.UserId, serverId e.ServerId) (*e.UserServerSettingsOverride, error)
func (r *PGUserNotiRepo) FindChannelPreferences(ctx context.Context, userId e.UserId, channelId e.ChannelId) (*e.UserChannelSettingsOverride, error)
func (r *PGUserNotiRepo) FindDMGroupPreferences(ctx context.Context, userId e.UserId, groupId e.DMGroupId) (*e.UserDMSettingsOverride, error)
func (r *PGUserNotiRepo) SaveServerPreferences(ctx context.Context, preference *e.UserServerSettingsOverride) (*e.UserServerSettingsOverride, error)
func (r *PGUserNotiRepo) SaveChannelPreferences(ctx context.Context, preference *e.UserChannelSettingsOverride) (*e.UserChannelSettingsOverride, error)
func (r *PGUserNotiRepo) SaveDMGroupPreferences(ctx context.Context, preference *e.UserDMSettingsOverride) (*e.UserDMSettingsOverride, error)

var _ repositories.UserNotificationRepo = &PGUserNotiRepo{}
