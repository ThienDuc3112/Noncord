package postgres

import (
	"backend/internal/domain/entities"
	"backend/internal/infra/db/postgres/gen"

	"github.com/google/uuid"
)

func fromDbMembership(membership gen.Membership, roles []uuid.UUID) *entities.Membership {
	rolesMap := make(map[entities.RoleId]bool)
	for _, id := range roles {
		rolesMap[entities.RoleId(id)] = true
	}
	return &entities.Membership{
		Id:        entities.MembershipId(membership.ID),
		ServerId:  entities.ServerId(membership.ServerID),
		UserId:    entities.UserId(membership.UserID),
		Nickname:  membership.Nickname,
		CreatedAt: membership.CreatedAt,
		Roles:     rolesMap,
	}
}
