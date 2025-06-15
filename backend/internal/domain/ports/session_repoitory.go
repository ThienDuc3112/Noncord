package ports

import (
	"backend/internal/domain/entities"
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
}

type SessionRepository interface {
	Save(*Session) error
	FindById(uuid.UUID) (*Session, error)
	FindByToken(string) (*Session, error)
	FindByUserId(entities.UserId) ([]*Session, error)
}
