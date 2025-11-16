package query

import (
	"backend/internal/domain/entities"

	"github.com/google/uuid"
)

type CheckChannelPerm struct {
	UserId     uuid.UUID
	ChannelId  uuid.UUID
	Permission entities.ServerPermissionBits
}
