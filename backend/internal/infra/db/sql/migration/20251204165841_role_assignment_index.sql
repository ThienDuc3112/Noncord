-- +goose Up
-- +goose StatementBegin
CREATE INDEX idx_role_assignment_membership_id ON role_assignment(membership_id);
CREATE INDEX idx_role_assignment_role_id ON role_assignment(role_id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP INDEX idx_role_assignment_role_id;
DROP INDEX idx_role_assignment_membership_id;
-- +goose StatementEnd
