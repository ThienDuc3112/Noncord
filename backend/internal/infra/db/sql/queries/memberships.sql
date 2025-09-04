-- name: SaveMembership :one
INSERT INTO memberships (
  server_id,
  user_id,
  created_at,
  nickname
) VALUES (
  $1,
  $2,
  $3,
  $4
)
ON CONFLICT (server_id, user_id)
DO UPDATE SET
  nickname = $4
RETURNING *;

-- name: FindMembership :one
SELECT * FROM memberships WHERE server_id = $1 AND user_id = $2;

-- name: FindMembershipsByUserId :many
SELECT * FROM memberships WHERE user_id = $1;

-- name: FindMembershipsByServerId :many
SELECT * FROM memberships WHERE server_id = $1;

-- name: DeleteMembership :exec
DELETE FROM memberships WHERE user_id = $1 AND server_id = $2;
