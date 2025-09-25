package postgres

import (
	"backend/internal/domain/entities"
	"backend/internal/infra/db/postgres/gen"
)

func fromDbInvitation(inv gen.Invitation) *entities.Invitation {
	return &entities.Invitation{
		Id:             entities.InvitationId(inv.ID),
		ServerId:       entities.ServerId(inv.ServerID),
		CreatedAt:      inv.CreatedAt,
		ExpiresAt:      inv.ExpiredAt,
		BypassApproval: inv.BypassApproval,
		JoinLimit:      inv.JoinLimit,
		JoinCount:      inv.JoinCount,
	}
}
