-- +goose Up
-- +goose StatementBegin
CREATE TABLE emotes (
	id UUID NOT NULL PRIMARY KEY,
  created_at TIMESTAMP WITH TIME ZONE NOT NULL,
  updated_at TIMESTAMP WITH TIME ZONE NOT NULL,
  deleted_at TIMESTAMP WITH TIME ZONE,
  server_id UUID NOT NULL REFERENCES servers(id) ON DELETE CASCADE,
	name VARCHAR(64) NOT NULL,
	icon_url VARCHAR(2048) NOT NULL
);
CREATE INDEX idx_emotes_server_id ON emotes(server_id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE emotes;
-- +goose StatementEnd
