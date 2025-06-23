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
ON CONFLICT (id)
DO UPDATE SET 
  created_at = $2,
  updated_at = $3,
  deleted_at = $4,
  username = $5,
  display_name = $6,
  about_me = $7,
  email = $8,
  password = $9,
  disabled = $10,
  avatar_url = $11,
  banner_url = $12,
  flags = $13
RETURNING *;

-- name: CreateUserSetting :one
INSERT INTO user_settings (
  user_id,
  language,
	dm_allow_option,
	dm_filter_option,
	friend_request_permission,
	collect_analytics_permission,
	theme,
	show_emote,
	notification_settings,
	afk_timeout
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
  language = $2,
	dm_allow_option = $3,
	dm_filter_option = $4,
	friend_request_permission = $5,
	collect_analytics_permission = $6,
	theme = $7,
	show_emote = $8,
	notification_settings = $9,
	afk_timeout = $10
RETURNING *;

-- name: FindUserById :one
SELECT * FROM users WHERE id = $1 AND deleted_at IS NULL;

-- name: FindUserByUsername :one
SELECT * FROM users WHERE username = $1 AND deleted_at IS NULL;

-- name: FindUserByEmail :one
SELECT * FROM users WHERE email = $1 AND deleted_at IS NULL;
