package response

import (
	"time"

	"github.com/google/uuid"
)

type Membership struct {
	ServerId      uuid.UUID   `json:"serverId"`
	UserId        uuid.UUID   `json:"userId"`
	Nickname      string      `json:"nickname"`
	CreatedAt     time.Time   `json:"createdAt"`
	AssignedRoles []uuid.UUID `json:"assignedRoles"`
}
