-- +goose Up
-- +goose StatementBegin
CREATE TABLE memberships (
  server_id UUID NOT NULL REFERENCES servers(id) ON DELETE CASCADE,
  user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
  created_at TIMESTAMP WITH TIME ZONE NOT NULL,
  PRIMARY KEY(server_id, user_id)
);
CREATE INDEX idx_memberships_user_id_server_id ON memberships(user_id, server_id);

CREATE TABLE role_assignment (
  server_id UUID NOT NULL REFERENCES servers(id) ON DELETE CASCADE,
  user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
  role_id UUID NOT NULL REFERENCES roles(id) ON DELETE CASCADE,
  created_at TIMESTAMP WITH TIME ZONE NOT NULL,
  FOREIGN KEY(server_id, user_id) REFERENCES memberships(server_id, user_id) ON DELETE CASCADE,
  PRIMARY KEY(server_id, user_id, role_id)
);
CREATE INDEX idx_role_assignment_server_id_user_id ON role_assignment(server_id, user_id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE role_assignment;
DROP TABLE memberships;
-- +goose StatementEnd
