package entities

import (
	"time"

	"github.com/google/uuid"
)

type DMGroupMember struct {
	Member   UserId
	JoinedAt time.Time
}

type DMGroupId uuid.UUID

type DMGroup struct {
	Id        DMGroupId
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
	Name      string
	IconUrl   string
	IsGroup   bool
	Members   []DMGroupMember
}

func (g *DMGroup) Validate() error {
	if len(g.Name) > 64 {
		return NewError(ErrCodeValidationError, "name cannot exceed 64 characters", nil)
	}
	if !g.IsGroup && g.IconUrl != "" {
		return NewError(ErrCodeValidationError, "direct message cannot set icon url", nil)
	}
	if g.IconUrl != "" && !emailReg.MatchString(g.IconUrl) {
		return NewError(ErrCodeValidationError, "invalid icon url", nil)
	}
	if len(g.IconUrl) > 2048 {
		return NewError(ErrCodeValidationError, "icon url too long", nil)
	}
	if !g.IsGroup && len(g.Members) > 2 {
		return NewError(ErrCodeValidationError, "direct message cannot have more than 2 members", nil)
	}
	if g.IsGroup && len(g.Members) > 100 {
		return NewError(ErrCodeValidationError, "chat group cannot have more than 100 members, consider making a server", nil)
	}
	return nil
}
