-- name: FindMessageById :one
SELECT * FROM messages WHERE id = $1 AND deleted_at IS NULL;

-- name: FindMessagesByChannelId :many
SELECT * FROM messages WHERE channel_id = $1 AND created_at < $2 AND deleted_at IS NULL ORDER BY created_at DESC LIMIT $3;

-- name: FindMessagesByGroupId :many
SELECT * FROM messages WHERE group_id = $1 AND created_at < $2 AND deleted_at IS NULL ORDER BY created_at DESC LIMIT $3;

-- name: SaveMessage :one
INSERT INTO messages (
	id,
  created_at,
  updated_at,
  deleted_at,
  channel_id,
  group_id,
  author_id,
  message
) VALUES ($1, $2, $3, $4, $5, $6, $7, $8) 
RETURNING *;
