package mapper

import (
	"backend/internal/application/common"
	"backend/internal/domain/entities"

	"github.com/google/uuid"
)

func MembershipToResult(m *entities.Membership) *common.Membership {
	return &common.Membership{
		ServerId:  uuid.UUID(m.ServerId),
		UserId:    uuid.UUID(m.UserId),
		Nickname:  m.Nickname,
		CreatedAt: m.CreatedAt,
	}
}
