package response

import (
	"time"

	"github.com/google/uuid"
)

type NewServerResponse struct {
	Id uuid.UUID `json:"id"`
}

type GetServerResponse struct {
	Id          uuid.UUID
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Name        string
	Description string
	IconUrl     string
	BannerUrl   string

	// Channels
	// Members
}

type UpdateServerResponse struct {
}
