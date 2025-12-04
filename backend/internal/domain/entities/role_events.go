package entities

import (
	"backend/internal/domain/events"

	"github.com/google/uuid"
)

const (
	EventRoleCreated = "role.created"

	RoleCreatedSchemaVersion = 1
)

type RoleCreated struct {
	events.Base

	Name         string               `json:"name"`
	Color        uint32               `json:"color"`
	Priority     uint16               `json:"priority"`
	AllowMention bool                 `json:"allowMention"`
	Permissions  ServerPermissionBits `json:"permissions"`
	ServerId     ServerId             `json:"serverId"`
}

func NewRoleCreated(r *Role) RoleCreated {
	return RoleCreated{
		Base: events.NewBase("role", uuid.UUID(r.Id), EventRoleCreated, RoleCreatedSchemaVersion),

		Name:         r.Name,
		Color:        r.Color,
		Priority:     r.Priority,
		AllowMention: r.AllowMention,
		Permissions:  r.Permissions,
		ServerId:     r.ServerId,
	}
}
