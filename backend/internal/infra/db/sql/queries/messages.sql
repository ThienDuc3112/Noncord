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
  author_type,
  message
) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9) 
RETURNING *;

-- name: GetEnrichedMessageById :one
SELECT m.*, u.display_name, u.avatar_url, mb.nickname FROM messages m
LEFT JOIN users u ON m.author_id = u.id
LEFT JOIN channels c ON m.channel_id = c.id
LEFT JOIN servers s ON c.server_id = s.id
LEFT JOIN memberships mb ON mb.user_id = u.id AND mb.server_id = s.id
WHERE m.id = $1 AND m.deleted_at IS NULL;

-- name: GetEnrichedMessageByChannelId :many
SELECT m.*, u.display_name, u.avatar_url, mb.nickname FROM messages m
LEFT JOIN users u ON m.author_id = u.id
LEFT JOIN channels c ON m.channel_id = c.id
LEFT JOIN servers s ON c.server_id = s.id
LEFT JOIN memberships mb ON mb.user_id = u.id AND mb.server_id = s.id
WHERE m.channel_id = $1 AND m.created_at < $2 AND m.deleted_at IS NULL ORDER BY m.created_at DESC LIMIT $3;

-- name: GetEnrichedMessageByGroupId :many
SELECT m.*, u.display_name, u.avatar_url FROM messages m
JOIN users u ON m.author_id = u.id
WHERE m.group_id = $1 AND m.created_at < $2 AND m.deleted_at IS NULL ORDER BY m.created_at DESC LIMIT $3;
