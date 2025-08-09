package command

import (
	"backend/internal/application/common"

	"github.com/google/uuid"
)

type UpdateServerCommand struct {
	UserId uuid.UUID
	Server common.Server
}
