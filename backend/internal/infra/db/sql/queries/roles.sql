-- name: SaveRole :one
INSERT INTO roles (
	id,
  created_at,
  updated_at,
  deleted_at,
	name,
	color,
	priority,
	allow_mention,
	permissions,
	server_id
) VALUES (
  $1,
  $2,
  $3,
  $4,
  $5,
  $6,
  $7,
  $8,
  $9,
  $10
) 
ON CONFLICT (id)
DO UPDATE SET
  updated_at = $3,
  deleted_at = $4,
	name = $5,
	color = $6,
	priority = $7,
	allow_mention = $8,
	permissions = $9
RETURNING *;

-- name: FindRolesByServerId :many
SELECT * FROM roles WHERE server_id = $1 AND deleted_at IS NULL;
