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
	announcement_channel,
  owner
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
  $10,
  $11
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
	announcement_channel = $10,
  owner = $11
RETURNING *;

-- name: FindServerById :one
SELECT * FROM servers WHERE id = $1 AND deleted_at IS NULL;

-- name: FindServersByIds :many
SELECT * FROM servers WHERE id = ANY(@ids::UUID[]) AND deleted_at IS NULL;

-- name: DeleteServer :exec
UPDATE servers SET deleted_at = NOW() WHERE id = $1;

-- name: FindInvitationById :one
SELECT * FROM invitations WHERE id = $1 AND (expired_at IS NULL OR expired_at > NOW()) AND (join_limit <= 0 OR join_limit > join_count);

-- name: FindInvitationsByServerId :many
SELECT * FROM invitations WHERE server_id = $1 AND (expired_at IS NULL OR expired_at > NOW()) AND (join_limit <= 0 OR join_limit > join_count);

-- name: SaveInvitation :one
INSERT INTO invitations (
	id,
	server_id,
  created_at,
  expired_at,
	bypass_approval,
	join_limit,
  join_count
) VALUES (
  $1,
  $2,
  $3,
  $4,
  $5,
  $6,
  $7
) 
ON CONFLICT (id)
DO UPDATE SET 
	server_id = $2,
  created_at = $3,
  expired_at = $4,
	bypass_approval = $5,
	join_limit = $6,
  join_count = $7
RETURNING *;

-- name: FindServerFromInviteId :one
SELECT s.* 
FROM servers s 
JOIN invitations i ON s.id = i.server_id
WHERE i.id = $1 AND (i.expired_at IS NULL OR i.expired_at > NOW()) AND (i.join_limit <= 0 OR i.join_limit > i.join_count);

-- name: FindServersFromUserId :many
SELECT s.* 
FROM servers s
JOIN memberships m ON m.server_id = s.id
WHERE m.user_id = $1 AND s.deleted_at IS NULL;
