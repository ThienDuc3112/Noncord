package response

import "github.com/google/uuid"

type Role struct {
	Id           uuid.UUID `json:"id"`
	Name         string    `json:"name"`
	Color        uint32    `json:"color"`
	Priority     uint16    `json:"priority"`
	AllowMention bool      `json:"allowMention"`
	Permissions  []string  `json:"permissions"`
	ServerId     uuid.UUID `json:"serverId"`
}
