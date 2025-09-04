package postgres

import (
	"backend/internal/domain/entities"
	"backend/internal/infra/db/postgres/gen"
)

func fromDbMembership(membership gen.Membership) *entities.Membership {
	return &entities.Membership{
		ServerId:  entities.ServerId(membership.ServerID),
		UserId:    entities.UserId(membership.UserID),
		Nickname:  membership.Nickname,
		CreatedAt: membership.CreatedAt,
	}
}
