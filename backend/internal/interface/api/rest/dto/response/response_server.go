package response

import (
	"time"

	"github.com/google/uuid"
)

type NewServerResponse struct {
	Id uuid.UUID `json:"id"`
}

type ServerPreview struct {
	Id        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	IconUrl   string    `json:"iconUrl"`
	BannerUrl string    `json:"bannerUrl"`
}

type GetServersResponse struct {
	Result []ServerPreview `json:"result"`
}

type GetServerResponse struct {
	Id          uuid.UUID `json:"id"`
	CreatedAt   time.Time `json:"createdAt,omitempty,omitzero"`
	UpdatedAt   time.Time `json:"updatedAt,omitempty,omitzero"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	IconUrl     string    `json:"iconUrl"`
	BannerUrl   string    `json:"bannerUrl"`

	// Channels
	// Members
}

type UpdateServerResponse struct {
	Id          uuid.UUID `json:"id"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	IconUrl     string    `json:"iconUrl"`
	BannerUrl   string    `json:"bannerUrl"`
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

type GetServerInvitationsResponse struct {
	Result []Invitation `json:"result"`
}
