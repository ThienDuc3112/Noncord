package request

import (
	"net/http"

	"github.com/google/uuid"
)

type CreateChannel struct {
	ServerId       uuid.UUID  `json:"serverId" validate:"required"`
	Name           string     `json:"name" validate:"required"`
	Description    string     `json:"description"`
	ParentCategory *uuid.UUID `json:"parentCategory"`
}

func (r *CreateChannel) Bind(_ *http.Request) error {
	return validate.Struct(r)
}

type UpdateChannel struct {
	Name        *string `json:"name" validate:"required_without_all=Description"`
	Description *string `json:"description" validate:"required_without_all=Name"`
}

func (r *UpdateChannel) Bind(_ *http.Request) error {
	return validate.Struct(r)
}
