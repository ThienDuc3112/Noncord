-- name: CreateSession :one
INSERT INTO sessions (
  id,
  rotation_count,
  created_at,
  updated_at,
  expires_at,
  user_id,
  user_agent,
  refresh_token
) VALUES (
  $1,
  $2,
  $3,
  $4,
  $5,
  $6,
  $7,
  $8
) 
ON CONFLICT (id)
DO UPDATE SET 
  rotation_count = $2,
  created_at = $3,
  updated_at = $4,
  expires_at = $5,
  user_id = $6,
  user_agent = $7,
  refresh_token = $8
RETURNING *;

-- name: FindSessionById :one
SELECT * FROM sessions WHERE id = $1;
