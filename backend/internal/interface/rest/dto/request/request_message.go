package request

import (
	"net/http"

	"github.com/google/uuid"
)

type CreateMessage struct {
	TargetId        uuid.UUID `json:"targetId" validate:"required"`
	IsTargetChannel bool      `json:"isTargetChannel"`
	Content         string    `json:"content" validate:"required"`
}

func (r *CreateMessage) Bind(_ *http.Request) error {
	return validate.Struct(r)
}
