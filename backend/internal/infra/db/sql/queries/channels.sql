-- name: SaveChannel :one
INSERT INTO channels (
	id,
  created_at,
  updated_at,
  deleted_at,
	name,
	description,
	server_id,
	ordering,
	parent_category
) VALUES (
  $1,
  $2,
  $3,
  $4,
  $5,
  $6,
  $7,
  $8,
  $9
) 
ON CONFLICT (id)
DO UPDATE SET
  updated_at = $3,
  deleted_at = $4,
	name = $5,
	description = $6,
	server_id = $7,
	ordering = $8,
	parent_category = $9
RETURNING *;

-- name: FindChannelById :one
SELECT * FROM channels WHERE id = $1 AND deleted_at IS NULL;

-- name: FindChannelsByServerId :many
SELECT * FROM channels WHERE server_id = $1 AND deleted_at IS NULL ORDER BY ordering;

-- name: FindChannelsByIds :many
SELECT * FROM channels WHERE id = ANY(@ids::UUID[]) AND deleted_at IS NULL;

-- name: DeleteChannel :exec
UPDATE channels SET deleted_at = NOW() WHERE id = $1;

-- name: GetServerMaxOrdering :one
SELECT COALESCE(MAX(ordering), 0)::int AS max_order FROM channels WHERE server_id = $1 AND deleted_at IS NULL;
