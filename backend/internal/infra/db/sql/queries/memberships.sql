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

-- name: FindRoleAssignmentsByMembershipId :many
SELECT role_id FROM role_assignment WHERE membership_id = $1;

-- name: FindMembershipsByUserId :many
SELECT * FROM memberships WHERE user_id = $1;

-- name: FindRoleAssignmentsByUserId :many
SELECT ra.* FROM role_assignment ra, memberships mb WHERE ra.membership_id = mb.id AND mb.user_id = $1;

-- name: FindMembershipsByServerId :many
SELECT * FROM memberships WHERE server_id = $1;

-- name: FindRoleAssignmentsByServerId :many
SELECT ra.* FROM role_assignment ra, memberships mb WHERE ra.membership_id = mb.id AND mb.server_id = $1;

-- name: DeleteMembership :exec
DELETE FROM memberships WHERE user_id = $1 AND server_id = $2;

-- name: FindMembershipWithChannelId :one
SELECT mb.* FROM memberships mb, channels c WHERE c.id = $1 AND mb.user_id = $2 AND mb.server_id = c.server_id;

-- name: SetMembershipRoles :many
WITH deleted AS (
  DELETE FROM role_assignment WHERE membership_id = @membership_id
)
INSERT INTO role_assignment (membership_id, role_id) 
SELECT @membership_id, UNNEST(COALESCE(@role_ids::uuid[], '{}'::uuid[]))
ON CONFLICT DO NOTHING
RETURNING role_id;

