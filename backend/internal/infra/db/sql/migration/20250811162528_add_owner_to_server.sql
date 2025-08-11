-- +goose Up
-- +goose StatementBegin
ALTER TABLE servers ADD COLUMN owner UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE servers DROP COLUMN owner;
-- +goose StatementEnd
