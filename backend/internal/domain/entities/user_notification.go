package entities

import "time"

type Scope string

const (
	ScopeServer  Scope = "SERVER"
	ScopeChannel Scope = "CHANNEL"
	ScopeDM      Scope = "DM"
)

type UserNotification struct {
	UserId               UserId
	ServerId             ServerId
	UpdatedAt            time.Time
	Scope                Scope
	NotificationSettings NotificationBits
}
