-- +goose Up
-- +goose StatementBegin
CREATE TABLE outbox (
  id UUID NOT NULL PRIMARY KEY,
  aggregate_name TEXT NOT NULL,
  aggregate_id UUID NOT NULL,

  event_type TEXT NOT NULL,
  schema_version INT NOT NULL,
  occurred_at TIMESTAMP WITH TIME ZONE NOT NULL,

  payload JSONB NOT NULL,

  status TEXT NOT NULL DEFAULT 'pending', -- pending|inflight|dispatched|failed
  attempts INT NOT NULL DEFAULT 0,
  claimed_at TIMESTAMP WITH TIME ZONE,
  published_at TIMESTAMP WITH TIME ZONE
);

DROP TABLE channel_user_permission_override;
DROP TABLE channel_role_permission_override;

CREATE TYPE overwrite_target AS ENUM ('role','user');
CREATE TABLE channel_permission_overwrite (
	channel_id UUID NOT NULL REFERENCES channels(id) ON DELETE CASCADE,
	role_id UUID REFERENCES roles(id) ON DELETE CASCADE,
	user_id UUID REFERENCES roles(id) ON DELETE CASCADE,
  target_type overwrite_target NOT NULL,
  updated_at TIMESTAMP WITH TIME ZONE NOT NULL,
	allow BIGINT NOT NULL,
	deny BIGINT NOT NULL,
  CHECK (
    (target_type='role' AND role_id IS NOT NULL AND user_id IS NULL) OR 
    (target_type='user' AND user_id IS NOT NULL AND role_id IS NULL)
  )
);

CREATE UNIQUE INDEX ux_channel_permission_overwrite_role 
  ON channel_permission_overwrite(channel_id, role_id)
  WHERE role_id IS NOT NULL;
CREATE UNIQUE INDEX ux_channel_permission_overwrite_user 
  ON channel_permission_overwrite(channel_id, user_id)
  WHERE user_id IS NOT NULL;
CREATE INDEX ix_channel_permission_overwrite_channel ON channel_permission_overwrite(channel_id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE channel_permission_overwrite;
DROP TYPE overwrite_target;

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

DROP TABLE outbox;
-- +goose StatementEnd
