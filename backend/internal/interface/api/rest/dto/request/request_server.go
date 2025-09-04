package request

import (
	"net/http"

	"github.com/google/uuid"
)

type NewServer struct {
	Name string `json:"name" example:"My very good server" validate:"required,max=256"`
}

func (r *NewServer) Bind(_ *http.Request) error {
	return validate.Struct(r)
}

type UpdateServer struct {
	Name                *string       `json:"name"`
	Description         *string       `json:"description"`
	IconUrl             *string       `json:"iconUrl"`
	BannerUrl           *string       `json:"bannerUrl"`
	NeedApproval        *bool         `json:"needApproval"`
	AnnouncementChannel uuid.NullUUID `json:"announcementChannel"`
}

func (r *UpdateServer) Bind(_ *http.Request) error {
	return validate.Struct(r)
}
