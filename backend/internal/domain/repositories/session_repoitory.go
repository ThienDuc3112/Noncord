package repositories

import (
	"backend/internal/domain/entities"
	"context"

	"github.com/google/uuid"
)

type SessionRepo interface {
	Save(context.Context, *entities.Session) error
	FindById(context.Context, uuid.UUID) (*entities.Session, error)
	FindByToken(context.Context, string) (*entities.Session, error)
	FindByUserId(context.Context, entities.UserId) ([]*entities.Session, error)
}
