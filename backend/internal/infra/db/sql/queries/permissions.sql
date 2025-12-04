-- name: FindAllChannelInUserServers :many
SELECT c.id, c.server_id 
FROM channels c, memberships m
WHERE c.server_id = m.server_id AND m.user_id = $1;

-- name: FindAllRolesUserHave :many
SELECT mb.user_id, mb.server_id, ra.role_id, r.permissions FROM memberships mb
INNER JOIN role_assignment ra ON ra.membership_id = mb.id
INNER JOIN roles r ON ra.role_id = r.id
WHERE mb.user_id = $1;

-- name: FindAllUserRolePermission :many
SELECT mb.user_id, mb.server_id, ra.role_id, r.permissions FROM memberships mb
INNER JOIN role_assignment ra ON ra.membership_id = mb.id
INNER JOIN roles r ON ra.role_id = r.id
WHERE mb.user_id = $1
ORDER BY r.priority ASC;

-- name: FindAllUserServerRolePermission :many
SELECT mb.user_id, mb.server_id, ra.role_id, r.permissions FROM memberships mb
INNER JOIN role_assignment ra ON ra.membership_id = mb.id
INNER JOIN roles r ON ra.role_id = r.id
WHERE mb.user_id = $1 AND mb.server_id = $2
ORDER BY r.priority ASC;

-- name: FindAllChannelUserOverwrite :many
SELECT cpo.channel_id, cpo.allow, cpo.deny FROM 
  channels c,
  memberships mb,
  channel_permission_overwrite cpo
WHERE
  mb.user_id = $1
  AND mb.server_id = c.server_id
  AND cpo.channel_id = c.id 
  AND cpo.user_id = mb.user_id;

-- name: FindAllChannelUserRoleOverwrite :many
SELECT ra.role_id, mb.server_id, cpo.channel_id, cpo.allow, cpo.deny FROM 
  channels c,
  memberships mb,
  channel_permission_overwrite cpo,
  role_assignment ra
WHERE
  mb.user_id = $1
  AND mb.server_id = c.server_id
  AND cpo.channel_id = c.id 
  AND ra.membership_id = mb.id
  AND ra.role_id = cpo.role_id;

-- name: FindAllChannelServerUserOverwrite :many
SELECT cpo.channel_id, cpo.allow, cpo.deny FROM 
  channels c,
  memberships mb,
  channel_permission_overwrite cpo
WHERE
  mb.user_id = $1
  AND mb.server_id = $2
  AND mb.server_id = c.server_id
  AND cpo.channel_id = c.id 
  AND cpo.user_id = mb.user_id;

-- name: FindAllChannelUserServerRoleOverwrite :many
SELECT ra.role_id, cpo.channel_id, cpo.allow, cpo.deny FROM 
  channels c,
  memberships mb,
  channel_permission_overwrite cpo,
  role_assignment ra
WHERE
  mb.user_id = $1
  AND mb.server_id = $2
  AND mb.server_id = c.server_id
  AND cpo.channel_id = c.id 
  AND ra.membership_id = mb.id
  AND ra.role_id = cpo.role_id;
