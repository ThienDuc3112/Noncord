package common

import (
	"time"

	"github.com/google/uuid"
)

type UserResult struct {
	Id          uuid.UUID
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Username    string
	DisplayName string
	AboutMe     string
	Email       string
	Disabled    bool
	Verified    bool
	AvatarUrl   string
	BannerUrl   string
	Flags       uint16
}
