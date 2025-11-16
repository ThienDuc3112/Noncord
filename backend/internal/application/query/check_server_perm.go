package query

import (
	"backend/internal/domain/entities"

	"github.com/google/uuid"
)

type CheckServerPerm struct {
	UserId     uuid.UUID
	ServerId   uuid.UUID
	Permission entities.ServerPermissionBits
}
