package mapper

import (
	"backend/internal/application/common"
	"backend/internal/domain/entities"

	"github.com/google/uuid"
)

func ServerToResult(s *entities.Server) *common.Server {
	return &common.Server{
		Id:                  uuid.UUID(s.Id),
		CreatedAt:           s.CreatedAt,
		UpdatedAt:           s.UpdatedAt,
		DeletedAt:           s.DeletedAt,
		Name:                s.Name,
		Description:         s.Description,
		IconUrl:             s.IconUrl,
		BannerUrl:           s.BannerUrl,
		NeedApproval:        s.NeedApproval,
		AnnouncementChannel: (*uuid.UUID)(s.AnnouncementChannel),
		DefaultRole:         (*uuid.UUID)(s.DefaultRole),
	}
}

func ServerToPreview(s *entities.Server) common.ServerPreview {
	return common.ServerPreview{
		Id:          uuid.UUID(s.Id),
		Name:        s.Name,
		Description: s.Description,
		IconUrl:     s.IconUrl,
		BannerUrl:   s.BannerUrl,
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
