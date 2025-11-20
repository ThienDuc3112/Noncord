-- name: FindAllChannelInUserServers :many
SELECT c.id 
FROM channels c, servers s, memberships m
WHERE c.server_id = s.id AND m.server_id = s.id AND m.user_id = $1;
