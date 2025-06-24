-- +goose Up
-- +goose StatementBegin
CREATE TABLE channels (
	id UUID NOT NULL PRIMARY KEY,
  created_at TIMESTAMP WITH TIME ZONE NOT NULL,
  updated_at TIMESTAMP WITH TIME ZONE NOT NULL,
  deleted_at TIMESTAMP WITH TIME ZONE,
	name VARCHAR(64) NOT NULL,
	description VARCHAR(256) NOT NULL,
	server_id UUID NOT NULL REFERENCES servers(id) ON DELETE CASCADE,
	ordering SMALLINT NOT NULL,
	parent_category UUID REFERENCES categories(id) ON DELETE SET NULL
);
CREATE INDEX idx_channels_server_id ON channels(server_id);
CREATE INDEX idx_channels_deleted_at ON channels(deleted_at);
ALTER TABLE servers 
  ADD CONSTRAINT fk_servers_announcement_channel 
  FOREIGN KEY(announcement_channel) REFERENCES channels(id) ON DELETE SET NULL;

CREATE TABLE channel_role_permission_override (
	channel_id UUID NOT NULL REFERENCES channels(id) ON DELETE CASCADE,
	role_id UUID NOT NULL REFERENCES roles(id) ON DELETE CASCADE,
  updated_at TIMESTAMP WITH TIME ZONE NOT NULL,
	allow BIGINT NOT NULL,
	deny BIGINT NOT NULL
);
CREATE INDEX idx_crpo_channel_id ON channel_role_permission_override(channel_id);
CREATE INDEX idx_crpo_role_id ON channel_role_permission_override(role_id);

CREATE TABLE channel_user_permission_override (
	channel_id UUID NOT NULL REFERENCES channels(id) ON DELETE CASCADE,
	user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
  updated_at TIMESTAMP WITH TIME ZONE NOT NULL,
	allow BIGINT NOT NULL,
	deny BIGINT NOT NULL
);
CREATE INDEX idx_cupo_channel_id ON channel_user_permission_override(channel_id);
CREATE INDEX idx_cupo_user_id ON channel_user_permission_override(user_id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE channel_user_permission_override;
DROP TABLE channel_role_permission_override;
ALTER TABLE servers DROP CONSTRAINT fk_servers_announcement_channel;
DROP TABLE channels;
-- +goose StatementEnd
