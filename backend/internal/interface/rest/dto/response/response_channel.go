package response

import (
	"time"

	"github.com/google/uuid"
)

type Channel struct {
	Id             uuid.UUID  `json:"id"`
	CreatedAt      time.Time  `json:"createdAt"`
	UpdatedAt      time.Time  `json:"updatedAt"`
	Name           string     `json:"name"`
	Description    string     `json:"description"`
	ServerId       uuid.UUID  `json:"serverId"`
	Order          uint16     `json:"order"`
	ParentCategory *uuid.UUID `json:"parentCategory"`
}
