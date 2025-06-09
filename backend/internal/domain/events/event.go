package events

import (
	"backend/internal/domain/entities"
	"time"

	"github.com/google/uuid"
)

type DomainEvent[T any] struct {
	TargetId  uuid.UUID
	Timestamp time.Time
	Action    string
	Payload   T
}

type BanEvent DomainEvent[*entities.BanEntry]
type ChannelEvent DomainEvent[*entities.Channel]
type DMGroupEvent DomainEvent[*entities.DMGroup]
type EmoteEvent DomainEvent[*entities.Emote]
type MemberEvent DomainEvent[*entities.Membership]
type MessageEvent DomainEvent[*entities.Message]
type RoleEvent DomainEvent[*entities.Role]
type ServerEvent DomainEvent[*entities.Server]
type UserEvent DomainEvent[*entities.User]
type UserServerNotificationEvent DomainEvent[*entities.UserNotification]
type OtherEvent DomainEvent[any]
