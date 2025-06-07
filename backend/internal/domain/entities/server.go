package entities

import (
	"time"

	"github.com/google/uuid"
)

type ServerId uuid.UUID

type Server struct {
	Id           ServerId
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    *time.Time
	Name         string
	Description  string
	IconUrl      string
	BannerUrl    string
	NeedApproval bool

	Categories []Category

	DefaultRole         *RoleId
	AnnouncementChannel *ChannelId
}

func (s *Server) Validate() error {
	if s.Name == "" {
		return NewError(ErrCodeValidationError, "cannot have empty server name", nil)
	}
	if len(s.Categories) > 255 {
		return NewError(ErrCodeValidationError, "cannot have more than 255 categories", nil)
	}
	if s.DefaultRole == nil {
		return NewError(ErrCodeValidationError, "server have no @everyone role", nil)
	}
	if s.IconUrl != "" && !emailReg.MatchString(s.IconUrl) {
		return NewError(ErrCodeValidationError, "icon invalid url", nil)
	}
	if s.BannerUrl != "" && !emailReg.MatchString(s.BannerUrl) {
		return NewError(ErrCodeValidationError, "banner invalid url", nil)
	}

	return nil
}

type CategoryId uuid.UUID

type Category struct {
	Id        CategoryId
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
	Name      string
	Order     uint8
}

type Membership struct {
	ServerId  ServerId
	UserId    UserId
	CreatedAt time.Time
}

type BanEntry struct {
	ServerId  ServerId
	UserId    UserId
	CreatedAt time.Time
}

type InvititationId uuid.UUID

type Invititation struct {
	Id             InvititationId
	CreatedAt      time.Time
	ExpiredAt      *time.Time
	BypassApproval bool
	ServerId       ServerId
	JoinLimit      int32
}

type JoinRequestId uuid.UUID

type JoinRequest struct {
	Id            JoinRequestId
	CreatedAt     time.Time
	ApprovedAt    *time.Time
	RevokedAt     *time.Time
	ServerId      ServerId
	UserId        UserId
	ApproverId    *UserId
	ApprovedState bool
}
