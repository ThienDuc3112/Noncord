package interfaces

import (
	"backend/internal/application/common"
	"context"

	"github.com/google/uuid"
)

type UserService interface {
}

type UserQueries interface {
	GetBasic(context.Context, uuid.UUID) (common.UserResult, error)
}
