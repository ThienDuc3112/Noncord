package postgres

import (
	"backend/internal/application/common"
	e "backend/internal/domain/entities"
	"backend/internal/infra/db/postgres/gen"
)

func fromDbServer(s gen.Server, rolesMap map[e.RoleId]*e.Role) *e.Server {
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
		DefaultRole:         (e.RoleId)(s.DefaultRole),
		AnnouncementChannel: (*e.ChannelId)(s.AnnouncementChannel),
		Owner:               e.UserId(s.Owner),
		Roles:               rolesMap,
	}
}

func toCommonServer(s gen.Server) common.Server {
	return common.Server{
		Id:                  s.ID,
		CreatedAt:           s.CreatedAt,
		UpdatedAt:           s.UpdatedAt,
		DeletedAt:           s.DeletedAt,
		Name:                s.Name,
		Description:         s.Description,
		IconUrl:             s.IconUrl,
		BannerUrl:           s.BannerUrl,
		NeedApproval:        s.NeedApproval,
		Owner:               s.Owner,
		DefaultRole:         s.DefaultRole,
		AnnouncementChannel: s.AnnouncementChannel,
	}
}
