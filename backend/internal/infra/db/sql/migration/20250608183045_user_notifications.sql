-- +goose Up
-- +goose StatementBegin
CREATE TABLE user_server_settings_overrides (
  server_id UUID NOT NULL REFERENCES servers(id) ON DELETE CASCADE,
  user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
  updated_at TIMESTAMP WITH TIME ZONE NOT NULL,
  notification_settings SMALLINT NOT NULL,
  PRIMARY KEY(server_id, user_id)
);
CREATE INDEX idx_usso_user_id_server_id ON user_server_settings_overrides(user_id, server_id);

CREATE TABLE user_channel_settings_overrides (
  channel_id UUID NOT NULL REFERENCES channels(id) ON DELETE CASCADE,
  user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
  updated_at TIMESTAMP WITH TIME ZONE NOT NULL,
  notification_settings SMALLINT NOT NULL,
  PRIMARY KEY(channel_id, user_id)
);
CREATE INDEX idx_ucso_user_id_server_id ON user_channel_settings_overrides(user_id, channel_id);

CREATE TABLE user_dm_group_settings_overrides (
  dm_group_id UUID NOT NULL REFERENCES dm_groups(id) ON DELETE CASCADE,
  user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
  updated_at TIMESTAMP WITH TIME ZONE NOT NULL,
  notification_settings SMALLINT NOT NULL,
  PRIMARY KEY(dm_group_id, user_id)
);
CREATE INDEX idx_udgso_user_id_server_id ON user_dm_group_settings_overrides(user_id, dm_group_id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE user_dm_group_settings_overrides;
DROP TABLE user_channel_settings_overrides;
DROP TABLE user_server_settings_overrides;
-- +goose StatementEnd
