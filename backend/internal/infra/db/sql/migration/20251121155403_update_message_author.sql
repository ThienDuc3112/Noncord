-- +goose Up
-- +goose StatementBegin
CREATE TYPE AUTHOR_TYPE AS ENUM('user', 'bot', 'system');
ALTER TABLE messages ADD COLUMN author_type AUTHOR_TYPE NOT NULL DEFAULT 'user';
ALTER TABLE messages ALTER COLUMN author_id DROP NOT NULL;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DELETE FROM messages WHERE author_id IS NULL;
ALTER TABLE messages ALTER COLUMN author_id SET NOT NULL;
ALTER TABLE messages DROP COLUMN author_type;
DROP TYPE AUTHOR_TYPE;
-- +goose StatementEnd
