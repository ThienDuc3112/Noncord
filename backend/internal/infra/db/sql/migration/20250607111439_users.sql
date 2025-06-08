-- +goose Up
-- +goose StatementBegin
CREATE TABLE users (
  id UUID NOT NULL PRIMARY KEY,
  created_at TIMESTAMP WITH TIME ZONE NOT NULL,
  updated_at TIMESTAMP WITH TIME ZONE NOT NULL,
  deleted_at TIMESTAMP WITH TIME ZONE,
  username VARCHAR(32) UNIQUE NOT NULL,
  display_name VARCHAR(128) NOT NULL,
  about_me VARCHAR(1024) NOT NULL,
  email VARCHAR(256) UNIQUE NOT NULL,
  password CHAR(60),
  disabled BOOLEAN NOT NULL,
  avatar_url VARCHAR(2048) NOT NULL,
  banner_url VARCHAR(2048) NOT NULL,
  flags SMALLINT NOT NULL
);


CREATE TABLE user_settings (
  user_id UUID REFERENCES users(id) ON DELETE CASCADE PRIMARY KEY,
  language VARCHAR(16) NOT NULL,

	-- Privacy settings
	dm_allow_option SMALLINT NOT NULL,
	dm_filter_option SMALLINT NOT NULL,
	friend_request_permission SMALLINT NOT NULL,
	collect_analytics_permission BOOLEAN NOT NULL,

	-- UI settings
	theme VARCHAR(8) NOT NULL,

	-- Chat settings
	show_emote BOOLEAN NOT NULL,

	-- Default Notifications
	notification_settings SMALLINT not NULL,
	afk_timeout BIGINT NOT NULL
);

CREATE TABLE friendships (
  created_at TIMESTAMP WITH TIME ZONE NOT NULL,
  user_id_1 UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
  user_id_2 UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
  PRIMARY KEY(user_id_1, user_id_2)
);
CREATE INDEX idx_friendships_user_id_2 ON friendships(user_id_2);

CREATE TABLE friend_request (
  created_at TIMESTAMP WITH TIME ZONE NOT NULL,
  requester UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
  target UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
  message VARCHAR(2048) NOT NULL,
  PRIMARY KEY(requester, target)
);
CREATE INDEX idx_friend_request_target ON friend_request(target);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE friend_request;
DROP TABLE friendships;
DROP TABLE user_settings;
DROP TABLE users;
-- +goose StatementEnd
