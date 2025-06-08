-- +goose Up
-- +goose StatementBegin
CREATE TABLE ban_entries (
  server_id UUID NOT NULL REFERENCES servers(id) ON DELETE CASCADE,
  user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
  created_at TIMESTAMP WITH TIME ZONE NOT NULL,
  PRIMARY KEY(server_id, user_id)
);
CREATE INDEX idx_ban_entries_user_id ON ban_entries(user_id);
CREATE INDEX idx_ban_entries_server_id ON ban_entries(server_id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE ban_entries;
-- +goose StatementEnd
