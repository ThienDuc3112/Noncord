package postgres

import (
	"backend/internal/application/common"
	"backend/internal/domain/entities"
	"backend/internal/infra/db/postgres/gen"
)

func fromDbRole(r gen.Role) *entities.Role {
	return &entities.Role{
		Id:           entities.RoleId(r.ID),
		CreatedAt:    r.CreatedAt,
		UpdatedAt:    r.UpdatedAt,
		DeletedAt:    r.DeletedAt,
		Name:         r.Name,
		ServerId:     entities.ServerId(r.ServerID),
		Permissions:  entities.ServerPermissionBits(r.Permissions),
		AllowMention: r.AllowMention,
		Priority:     uint16(r.Priority),
		Color:        uint32(r.Color),
	}
}

func toCommonRole(r gen.Role) common.Role {
	return common.Role{
		Id:           r.ID,
		Name:         r.Name,
		Color:        uint32(r.Color),
		Priority:     uint16(r.Priority),
		AllowMention: r.AllowMention,
		Permissions:  entities.ServerPermissionBits(r.Permissions).ToFlagArray(),
		ServerId:     r.ServerID,
	}
}
