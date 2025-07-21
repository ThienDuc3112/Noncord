-- name: SaveServer :one
INSERT INTO servers (
	id,
  created_at,
  updated_at,
	name,
	description,
	icon_url,
  banner_url,
	need_approval,
	default_role,
	announcement_channel
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
  created_at = $2,
  updated_at = $3,
	name = $4,
	description = $5,
	icon_url = $6,
  banner_url = $7,
	need_approval = $8,
	default_role = $9,
	announcement_channel = $10
RETURNING *;

-- name: FindServerById :one
SELECT * FROM servers WHERE id = $1 AND deleted_at IS NOT NULL;

-- name: FindServersByIds :many
SELECT * FROM servers WHERE id = ANY(@ids::UUID[]) AND deleted_at IS NOT NULL;

-- name: DeleteServer :exec
UPDATE servers SET deleted_at = NOW() WHERE id = $1;
