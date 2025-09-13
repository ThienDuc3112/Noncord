package response

import (
	"github.com/google/uuid"
)

type GetInvitationResponse struct {
	Id     uuid.UUID     `json:"id"`
	Server ServerPreview `json:"server"`
}
