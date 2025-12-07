-- +goose Up
-- +goose StatementBegin
DELETE FROM servers WHERE default_role IS NULL;
ALTER TABLE servers 
  ALTER COLUMN default_role SET NOT NULL;
ALTER TABLE servers
  DROP CONSTRAINT fk_servers_default_role;
ALTER TABLE servers
  ADD CONSTRAINT fk_servers_default_role FOREIGN KEY(default_role) REFERENCES roles(id) ON DELETE NO ACTION DEFERRABLE INITIALLY IMMEDIATE;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE servers 
  ALTER COLUMN default_role DROP NOT NULL;
ALTER TABLE servers
  DROP CONSTRAINT fk_servers_default_role;
ALTER TABLE servers
  ADD CONSTRAINT fk_servers_default_role FOREIGN KEY(default_role) REFERENCES roles(id) ON DELETE NO ACTION;
-- +goose StatementEnd
