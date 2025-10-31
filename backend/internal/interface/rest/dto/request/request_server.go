package request

import (
	"net/http"
	"time"

	"github.com/google/uuid"
)

type NewServer struct {
	Name string `json:"name" example:"My very good server" validate:"required,max=256"`
}

func (r *NewServer) Bind(_ *http.Request) error {
	return validate.Struct(r)
}

type UpdateServer struct {
	Name                *string    `json:"name" validate:"required_without_all=Description IconUrl BannerUrl NeedApproval AnnouncementChannel DefaultPermission"`
	Description         *string    `json:"description" validate:"required_without_all=Name IconUrl BannerUrl NeedApproval AnnouncementChannel DefaultPermission"`
	IconUrl             *string    `json:"iconUrl" validate:"required_without_all=Name Description BannerUrl NeedApproval AnnouncementChannel DefaultPermission"`
	BannerUrl           *string    `json:"bannerUrl" validate:"required_without_all=Name Description IconUrl NeedApproval AnnouncementChannel DefaultPermission"`
	NeedApproval        *bool      `json:"needApproval" validate:"required_without_all=Name Description IconUrl BannerUrl AnnouncementChannel DefaultPermission"`
	AnnouncementChannel *uuid.UUID `json:"announcementChannel" validate:"required_without_all=Name Description IconUrl BannerUrl NeedApproval DefaultPermission"`
	DefaultPermission   *int64     `json:"defaultPermission" validate:"required_without_all=Name Description IconUrl BannerUrl NeedApproval AnnouncementChannel"`
}

func (r *UpdateServer) Bind(_ *http.Request) error {
	return validate.Struct(r)
}

type NewInvitation struct {
	ExpiresAt      *time.Time `json:"expiresAt"`
	BypassApproval bool       `json:"bypassApproval"`
	JoinLimit      int32      `json:"joinLimit"`
}

func (r *NewInvitation) Bind(_ *http.Request) error {
	return validate.Struct(r)
}
