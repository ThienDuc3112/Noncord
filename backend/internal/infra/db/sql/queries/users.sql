-- name: CreateUser :one
INSERT INTO users (
  id,
  created_at,
  updated_at,
  deleted_at,
  username,
  display_name,
  about_me,
  email,
  password,
  disabled,
  avatar_url,
  banner_url,
  flags
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
  $11,
  $12,
  $13
)
RETURNING *;
