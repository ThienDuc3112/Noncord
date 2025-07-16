package common

import (
	"time"

	"github.com/google/uuid"
)

type Server struct {
	Id           uuid.UUID
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    *time.Time
	Name         string
	Description  string
	IconUrl      string
	BannerUrl    string
	NeedApproval bool

	DefaultRole         *uuid.UUID
	AnnouncementChannel *uuid.UUID
}
