-- +goose Up
-- +goose StatementBegin
CREATE TABLE roles (
	id UUID NOT NULL PRIMARY KEY,
  created_at TIMESTAMP WITH TIME ZONE NOT NULL,
  updated_at TIMESTAMP WITH TIME ZONE NOT NULL,
  deleted_at TIMESTAMP WITH TIME ZONE,
	name VARCHAR(64) NOT NULL,
	color INT NOT NULL,
	priority SMALLINT NOT NULL,
	allow_mention BOOLEAN NOT NULL,
	permissions BIGINT NOT NULL,
	server_id UUID NOT NULL REFERENCES servers(id) ON DELETE CASCADE
);
CREATE INDEX idx_roles_server_id ON roles(server_id);
CREATE INDEX idx_roles_deleted_at ON roles(deleted_at);
ALTER TABLE servers ADD CONSTRAINT fk_servers_default_role FOREIGN KEY(default_role) REFERENCES roles(id) ON DELETE SET NULL;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE servers DROP CONSTRAINT fk_servers_default_role;
DROP TABLE roles;
-- +goose StatementEnd
