-- +goose Up
-- +goose StatementBegin
CREATE TABLE servers (
	id UUID NOT NULL PRIMARY KEY,
  created_at TIMESTAMP WITH TIME ZONE NOT NULL,
  updated_at TIMESTAMP WITH TIME ZONE NOT NULL,
  deleted_at TIMESTAMP WITH TIME ZONE,
	name VARCHAR(256) NOT NULL,
	description VARCHAR(512) NOT NULL,
	icon_url VARCHAR(2048) NOT NULL,
  banner_url VARCHAR(2048) NOT NULL,
	need_approval BOOLEAN NOT NULL,
	-- Categories []Category
	default_role UUID,
	announcement_channel UUID
);

CREATE TABLE categories (
	id UUID NOT NULL PRIMARY KEY,
  created_at TIMESTAMP WITH TIME ZONE NOT NULL,
  updated_at TIMESTAMP WITH TIME ZONE NOT NULL,
  deleted_at TIMESTAMP WITH TIME ZONE,
	server_id UUID NOT NULL REFERENCES servers(id) ON DELETE CASCADE,
	name VARCHAR(256) NOT NULL,
  ordering SMALLINT NOT NULL
);
CREATE INDEX idx_categories_server_id ON categories(server_id);

CREATE TABLE invitations (
	id UUID NOT NULL PRIMARY KEY,
	server_id UUID NOT NULL REFERENCES servers(id) ON DELETE CASCADE,
  created_at TIMESTAMP WITH TIME ZONE NOT NULL,
  expired_at TIMESTAMP WITH TIME ZONE,
	bypass_approval BOOLEAN NOT NULL,
	join_limit INT NOT NULL
);
CREATE INDEX idx_invitations_server_id ON invitations(server_id);

CREATE TABLE join_requests (
	id UUID NOT NULL PRIMARY KEY,
  created_at TIMESTAMP WITH TIME ZONE NOT NULL,
  approved_at TIMESTAMP WITH TIME ZONE,
  revoked_at TIMESTAMP WITH TIME ZONE,
	server_id UUID NOT NULL REFERENCES servers(id) ON DELETE CASCADE,
  requester UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
  approver UUID REFERENCES users(id) ON DELETE CASCADE,
  approved_state BOOLEAN NOT NULL
);
CREATE INDEX idx_join_requests_server_id ON join_requests(server_id);
CREATE INDEX idx_join_requests_requester ON join_requests(requester);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE join_requests;
DROP TABLE invitations;
DROP TABLE categories;
DROP TABLE servers;
-- +goose StatementEnd
