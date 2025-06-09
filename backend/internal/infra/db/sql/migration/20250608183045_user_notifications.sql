-- +goose Up
-- +goose StatementBegin
CREATE TYPE scope_type AS ENUM('SERVER', 'CHANNEL', 'DM');
CREATE TABLE user_notification_overrides (
  reference_id UUID NOT NULL,
  user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
  updated_at TIMESTAMP WITH TIME ZONE NOT NULL,
  notification_settings SMALLINT NOT NULL,
  scope scope_type NOT NULL,
  PRIMARY KEY(reference_id, user_id)
);
CREATE INDEX idx_uno_user_id_reference_id ON user_notification_overrides(user_id, reference_id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE user_notification_overrides;
DROP TYPE scope_type;
-- +goose StatementEnd
