-- +goose Up
-- +goose StatementBegin
CREATE TABLE messages (
	id UUID NOT NULL PRIMARY KEY,
  created_at TIMESTAMP WITH TIME ZONE NOT NULL,
  updated_at TIMESTAMP WITH TIME ZONE NOT NULL,
  deleted_at TIMESTAMP WITH TIME ZONE,
  channel_id UUID REFERENCES channels(id) ON DELETE CASCADE,
  group_id UUID REFERENCES dm_groups(id) ON DELETE CASCADE,
  author_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
  message VARCHAR(4096) NOT NULL
);
CREATE INDEX idx_messages_channel_id ON messages(channel_id, created_at DESC);
CREATE INDEX idx_messages_group_id ON messages(group_id, created_at DESC);
CREATE INDEX idx_messages_deleted_at ON messages(deleted_at);

CREATE TABLE attachments (
	id UUID NOT NULL PRIMARY KEY,
	filetype VARCHAR(16) NOT NULL,
	url VARCHAR(2048) NOT NULL,
	filename VARCHAR(128) NOT NULL,
	message_id UUID REFERENCES messages(id) ON DELETE CASCADE,
  user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
	size INT NOT NULL
);

CREATE TABLE reactions (
  message_id UUID NOT NULL REFERENCES messages(id) ON DELETE CASCADE,
  user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
  emote_id UUID NOT NULL,
  PRIMARY KEY(message_id, user_id, emote_id)
);
CREATE INDEX idx_reactions_message_id ON reactions(message_id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE reactions;
DROP TABLE attachments;
DROP TABLE messages;
-- +goose StatementEnd
