// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.29.0
// source: sessions.sql

package gen

import (
	"context"
	"time"

	"github.com/google/uuid"
)

const createSession = `-- name: CreateSession :one
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
RETURNING id, rotation_count, created_at, updated_at, expires_at, user_id, user_agent, refresh_token
`

type CreateSessionParams struct {
	ID            uuid.UUID
	RotationCount int32
	CreatedAt     time.Time
	UpdatedAt     time.Time
	ExpiresAt     time.Time
	UserID        uuid.UUID
	UserAgent     string
	RefreshToken  string
}

func (q *Queries) CreateSession(ctx context.Context, arg CreateSessionParams) (Session, error) {
	row := q.db.QueryRowContext(ctx, createSession,
		arg.ID,
		arg.RotationCount,
		arg.CreatedAt,
		arg.UpdatedAt,
		arg.ExpiresAt,
		arg.UserID,
		arg.UserAgent,
		arg.RefreshToken,
	)
	var i Session
	err := row.Scan(
		&i.ID,
		&i.RotationCount,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.ExpiresAt,
		&i.UserID,
		&i.UserAgent,
		&i.RefreshToken,
	)
	return i, err
}

const findSessionById = `-- name: FindSessionById :one
SELECT id, rotation_count, created_at, updated_at, expires_at, user_id, user_agent, refresh_token FROM sessions WHERE id = $1 AND expires_at > NOW()
`

func (q *Queries) FindSessionById(ctx context.Context, id uuid.UUID) (Session, error) {
	row := q.db.QueryRowContext(ctx, findSessionById, id)
	var i Session
	err := row.Scan(
		&i.ID,
		&i.RotationCount,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.ExpiresAt,
		&i.UserID,
		&i.UserAgent,
		&i.RefreshToken,
	)
	return i, err
}

const findSessionByToken = `-- name: FindSessionByToken :one
SELECT id, rotation_count, created_at, updated_at, expires_at, user_id, user_agent, refresh_token FROM sessions WHERE refresh_token = $1 AND expires_at > NOW()
`

func (q *Queries) FindSessionByToken(ctx context.Context, refreshToken string) (Session, error) {
	row := q.db.QueryRowContext(ctx, findSessionByToken, refreshToken)
	var i Session
	err := row.Scan(
		&i.ID,
		&i.RotationCount,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.ExpiresAt,
		&i.UserID,
		&i.UserAgent,
		&i.RefreshToken,
	)
	return i, err
}

const findSessionsByUserId = `-- name: FindSessionsByUserId :many
SELECT id, rotation_count, created_at, updated_at, expires_at, user_id, user_agent, refresh_token FROM sessions WHERE user_id = $1 AND expires_at > NOW()
`

func (q *Queries) FindSessionsByUserId(ctx context.Context, userID uuid.UUID) ([]Session, error) {
	rows, err := q.db.QueryContext(ctx, findSessionsByUserId, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Session
	for rows.Next() {
		var i Session
		if err := rows.Scan(
			&i.ID,
			&i.RotationCount,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.ExpiresAt,
			&i.UserID,
			&i.UserAgent,
			&i.RefreshToken,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
