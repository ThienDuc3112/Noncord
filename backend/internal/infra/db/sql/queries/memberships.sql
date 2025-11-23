-- name: SaveMembership :one
INSERT INTO memberships (
  id,
  server_id,
  user_id,
  created_at,
  nickname
) VALUES (
  $1,
  $2,
  $3,
  $4,
  $5
)
ON CONFLICT (id)
DO UPDATE SET
  nickname = $5
RETURNING *;

-- name: FindMembership :one
SELECT * FROM memberships WHERE server_id = $1 AND user_id = $2;

-- name: FindMembershipsByUserId :many
SELECT * FROM memberships WHERE user_id = $1;

-- name: FindMembershipsByServerId :many
SELECT * FROM memberships WHERE server_id = $1;

-- name: DeleteMembership :exec
DELETE FROM memberships WHERE user_id = $1 AND server_id = $2;

-- name: FindMembershipWithChannelId :one
SELECT mb.* FROM memberships mb, channels c WHERE c.id = $1 AND mb.user_id = $2 AND mb.server_id = c.server_id;
