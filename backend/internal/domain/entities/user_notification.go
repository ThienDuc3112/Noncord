package entities

import (
	"time"

	"github.com/google/uuid"
)

type Scope string

const (
	ScopeServer  Scope = "SERVER"
	ScopeChannel Scope = "CHANNEL"
	ScopeDM      Scope = "DM"
)

type UserNotification struct {
	UserId               UserId
	ReferenceId          uuid.UUID
	UpdatedAt            time.Time
	Scope                Scope
	NotificationSettings NotificationBits
}

func NewUserNotification(uid UserId, refId uuid.UUID, scope Scope, setting NotificationBits) *UserNotification {
	return &UserNotification{
		UserId:               uid,
		ReferenceId:          refId,
		UpdatedAt:            time.Now(),
		Scope:                scope,
		NotificationSettings: setting,
	}
}
