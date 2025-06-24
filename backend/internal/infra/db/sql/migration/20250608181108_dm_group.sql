-- +goose Up
-- +goose StatementBegin
CREATE TABLE dm_groups (
	id UUID NOT NULL PRIMARY KEY,
  created_at TIMESTAMP WITH TIME ZONE NOT NULL,
  updated_at TIMESTAMP WITH TIME ZONE NOT NULL,
  deleted_at TIMESTAMP WITH TIME ZONE,
	name VARCHAR(64) NOT NULL,
	icon_url VARCHAR(2048) NOT NULL,
  is_group BOOLEAN NOT NULL
);
CREATE INDEX idx_dm_groups_deleted_at ON dm_groups(deleted_at);

CREATE TABLE dm_groups_member (
  member_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
  group_id UUID NOT NULL REFERENCES dm_groups(id) ON DELETE CASCADE,
  joined_at TIMESTAMP WITH TIME ZONE NOT NULL,
  PRIMARY KEY(member_id, group_id)
);
CREATE INDEX idx_dm_groups_member_member_id ON dm_groups_member(member_id);
CREATE INDEX idx_dm_groups_member_group_id ON dm_groups_member(group_id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE dm_groups_member;
DROP TABLE dm_groups;
-- +goose StatementEnd
