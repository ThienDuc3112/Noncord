package mapper

import (
	"backend/internal/application/common"
	"backend/internal/domain/entities"

	"github.com/google/uuid"
)

func InvitationToResult(s *entities.Invitation) *common.Invitation {
	return &common.Invitation{
		Id:             uuid.UUID(s.Id),
		CreatedAt:      s.CreatedAt,
		ServerId:       uuid.UUID(s.ServerId),
		ExpiresAt:      s.ExpiresAt,
		BypassApproval: s.BypassApproval,
		JoinLimit:      s.JoinLimit,
		JoinCount:      s.JoinCount,
	}
}

// func ResultToServer(s *common.Server) *entities.Server {
// 	return &entities.Server{
// 		Id:                  entities.ServerId(s.Id),
// 		CreatedAt:           s.CreatedAt,
// 		UpdatedAt:           s.UpdatedAt,
// 		DeletedAt:           s.DeletedAt,
// 		Name:                s.Name,
// 		Description:         s.Description,
// 		IconUrl:             s.IconUrl,
// 		BannerUrl:           s.BannerUrl,
// 		NeedApproval:        s.NeedApproval,
// 		DefaultRole:         (*entities.RoleId)(s.DefaultRole),
// 		AnnouncementChannel: (*entities.ChannelId)(s.AnnouncementChannel),
// 	}
// }
