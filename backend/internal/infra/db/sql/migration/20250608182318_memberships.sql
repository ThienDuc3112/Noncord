-- +goose Up
-- +goose StatementBegin
CREATE TABLE memberships (
  id UUID NOT NULL PRIMARY KEY, 
  server_id UUID NOT NULL REFERENCES servers(id) ON DELETE CASCADE,
  user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
  created_at TIMESTAMP WITH TIME ZONE NOT NULL,
  nickname VARCHAR(128) NOT NULL
);
CREATE INDEX idx_memberships_user_id_server_id ON memberships(user_id, server_id);

CREATE TABLE role_assignment (
  membership_id UUID NOT NULL REFERENCES memberships(id) ON DELETE CASCADE,
  role_id UUID NOT NULL REFERENCES roles(id) ON DELETE CASCADE,
  PRIMARY KEY(membership_id, role_id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE role_assignment;
DROP TABLE memberships;
-- +goose StatementEnd
