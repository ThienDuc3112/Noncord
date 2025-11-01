-- +goose Up
-- +goose StatementBegin
CREATE INDEX idx_outbox_claimable
ON outbox (status, claimed_at, occurred_at)
WHERE status IN ('pending','inflight');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP INDEX idx_outbox_claimable;
-- +goose StatementEnd
