-- +goose Up
-- +goose StatementBegin
ALTER TABLE memberships ADD CONSTRAINT unique_memberships_server_id_user_id UNIQUE (server_id,user_id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE memberships DROP CONSTRAINT unique_memberships_server_id_user_id;
-- +goose StatementEnd
