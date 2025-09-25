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

	DefaultPermission   int64
	AnnouncementChannel *uuid.UUID
}

type ServerPreview struct {
	Id          uuid.UUID
	Name        string
	IconUrl     string
	BannerUrl   string
	Description string
}
