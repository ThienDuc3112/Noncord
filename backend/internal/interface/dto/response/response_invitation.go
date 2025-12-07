package response

import (
	"time"

	"github.com/google/uuid"
)

type GetInvitationResponse struct {
	Id     uuid.UUID     `json:"id"`
	Server ServerPreview `json:"server"`
}

type Invitation struct {
	Id             uuid.UUID  `json:"id"`
	ServerId       uuid.UUID  `json:"serverId"`
	CreatedAt      time.Time  `json:"createdAt"`
	ExpiresAt      *time.Time `json:"expiresAt"`
	BypassApproval bool       `json:"bypassApproval"`
	JoinLimit      int32      `json:"joinLimit"`
	JoinCount      int32      `json:"joinCount"`
}

type JoinServerResponse struct {
	Server     ServerPreview `json:"server"`
	Membership Membership    `json:"membership"`
}
