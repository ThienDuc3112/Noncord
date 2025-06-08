package entities

import (
	"time"

	"github.com/google/uuid"
)

type EmoteId uuid.UUID

type Emote struct {
	Id        EmoteId
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
	ServerId  ServerId
	Name      string
	IconUrl   string
}

func (e *Emote) Validate() error {
	if e.IconUrl != "" && !emailReg.MatchString(e.IconUrl) {
		return NewError(ErrCodeValidationError, "invalid icon url", nil)
	}
	if len(e.IconUrl) > 2048 {
		return NewError(ErrCodeValidationError, "icon url too long", nil)
	}
	if len(e.Name) > 64 {
		return NewError(ErrCodeValidationError, "name cannot exceed 64 characters", nil)
	}

	return nil
}
