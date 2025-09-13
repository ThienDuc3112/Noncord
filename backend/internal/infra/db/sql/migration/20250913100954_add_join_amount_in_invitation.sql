-- +goose Up
-- +goose StatementBegin
ALTER TABLE invitations ADD COLUMN join_count INT NOT NULL DEFAULT 0;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE invitations DROP COLUMN join_count;
-- +goose StatementEnd
