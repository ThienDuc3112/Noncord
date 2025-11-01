-- name: InsertEventToOutbox :one
INSERT INTO outbox( 
  id,
  aggregate_name,
  aggregate_id,
  event_type,
  schema_version,
  occurred_at,
  payload
) 
VALUES ($1, $2, $3, $4, $5, $6, $7) 
RETURNING *;

-- name: ClaimOutboxBatch :many
WITH candidates AS (
  SELECT id
  FROM outbox
  WHERE status IN ('pending', 'inflight')
    AND (status = 'pending' OR claimed_at < now() - INTERVAL '2 minutes')
  ORDER BY occurred_at
  FOR UPDATE SKIP LOCKED
  LIMIT $1
)
UPDATE outbox AS o
SET status     = 'inflight',
    claimed_at = now(),
    attempts   = attempts + 1
FROM candidates c
WHERE o.id = c.id
RETURNING o.*;

-- name: MarkedOutboxDispatched :exec
UPDATE outbox SET status = 'dispatched', published_at = NOW() WHERE id = $1;

-- name: MarkedOutboxFailed :exec
UPDATE outbox SET status = 'failed' WHERE id = $1;

-- name: RequeueOutbox :exec
UPDATE outbox SET status = 'pending', claimed_at = NULL WHERE id = $1;
