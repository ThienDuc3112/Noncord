package repositories

import (
	e "backend/internal/domain/entities"
	"context"
)

type ChannelRepo interface {
	Find(context.Context, e.ChannelId) (*e.Channel, error)
	FindIds(context.Context, []e.ChannelId) ([]*e.Channel, error)
	FindByServerId(context.Context, e.ServerId) ([]*e.Channel, error)

	GetServerMaxChannelOrder(context.Context, e.ServerId) (int32, error)

	FindRoleOverrides(context.Context, e.ChannelId) ([]*e.ChannelRolePermissionOverride, error)
	FindRoleOverrideByRoleId(context.Context, e.ChannelId, e.RoleId) (*e.ChannelRolePermissionOverride, error)

	FindUserOverrides(context.Context, e.ChannelId) (*e.ChannelUserPermissionOverride, error)
	FindUserOverrideByUserId(context.Context, e.ChannelId, e.UserId) (*e.ChannelUserPermissionOverride, error)

	Save(context.Context, *e.Channel) (*e.Channel, error)
	SaveRoleOverride(context.Context, *e.ChannelRolePermissionOverride) (*e.ChannelRolePermissionOverride, error)
	SaveUserOverride(context.Context, *e.ChannelUserPermissionOverride) (*e.ChannelUserPermissionOverride, error)

	Delete(context.Context, e.ChannelId) error
	DeleteRoleOverride(context.Context, e.ChannelId, e.RoleId) error
	DeleteUserOverride(context.Context, e.ChannelId, e.UserId) error
}
