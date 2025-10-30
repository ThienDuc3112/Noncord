package postgres

import (
	e "backend/internal/domain/entities"
	"backend/internal/infra/db/postgres/gen"
)

func fromDbServer(s gen.Server) *e.Server {
	return &e.Server{
		Id:                  e.ServerId(s.ID),
		CreatedAt:           s.CreatedAt,
		UpdatedAt:           s.UpdatedAt,
		DeletedAt:           s.DeletedAt,
		Name:                s.Name,
		Description:         s.Description,
		IconUrl:             s.IconUrl,
		BannerUrl:           s.BannerUrl,
		NeedApproval:        s.NeedApproval,
		DefaultPermission:   e.ServerPermissionBits(s.DefaultPermission),
		AnnouncementChannel: (*e.ChannelId)(s.AnnouncementChannel),
		Owner:               e.UserId(s.Owner),
	}
}
