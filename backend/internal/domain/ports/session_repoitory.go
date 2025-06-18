package ports

import (
	"backend/internal/domain/entities"
	"context"
	"time"

	"github.com/google/uuid"
)

type Session struct {
	Id            uuid.UUID
	RotationCount int32
	CreatedAt     time.Time
	UpdatedAt     time.Time
	ExpiresAt     time.Time
	UserId        entities.UserId
	UserAgent     string
	Token         string
}

type SessionRepository interface {
	Save(context.Context, *Session) error
	FindById(context.Context, uuid.UUID) (*Session, error)
	FindByToken(context.Context, string) (*Session, error)
	FindByUserId(context.Context, entities.UserId) ([]*Session, error)
}
