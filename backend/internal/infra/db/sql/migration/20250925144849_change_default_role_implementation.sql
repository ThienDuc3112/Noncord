-- +goose Up
-- +goose StatementBegin
ALTER TABLE servers DROP COLUMN default_role;
ALTER TABLE servers ADD COLUMN default_permssion BIGINT NOT NULL DEFAULT 1302913;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE servers DROP COLUMN default_permssion;
ALTER TABLE servers ADD COLUMN default_role UUID REFERENCES roles(id) ON DELETE SET NULL;
-- +goose StatementEnd
