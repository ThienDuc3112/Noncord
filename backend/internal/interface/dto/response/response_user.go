package response

import (
	"time"

	"github.com/google/uuid"
)

type GetUser struct {
	Id          uuid.UUID `json:"id"`
	CreatedAt   time.Time `json:"createdAt"`
	Username    string    `json:"username"`
	DisplayName string    `json:"displayName"`
	AboutMe     string    `json:"aboutMe"`
	Email       string    `json:"email"`
	Disabled    bool      `json:"disabled"`
	AvatarUrl   string    `json:"avatarUrl"`
	BannerUrl   string    `json:"bannerUrl"`
	Flags       uint16    `json:"flags"`
}
