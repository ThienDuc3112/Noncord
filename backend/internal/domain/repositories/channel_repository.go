package repositories

import (
	e "backend/internal/domain/entities"
	"context"
)

type ChannelRepo interface {
	Find(ctx context.Context, id e.ChannelId) (*e.Channel, error)
	FindIds(ctx context.Context, ids []e.ChannelId) ([]*e.Channel, error)
	FindByServerId(ctx context.Context, serverId e.ServerId) ([]*e.Channel, error)

	FindRoleOverrides(ctx context.Context, id e.ChannelId) ([]*e.ChannelRolePermissionOverride, error)
	FindRoleOverrideByRoleId(ctx context.Context, id e.ChannelId, roleId e.RoleId) (*e.ChannelRolePermissionOverride, error)

	FindUserOverrides(ctx context.Context, id e.ChannelId) (*e.ChannelUserPermissionOverride, error)
	FindUserOverrideByUserId(ctx context.Context, id e.ChannelId, userId e.UserId) (*e.ChannelUserPermissionOverride, error)

	Save(ctx context.Context, channel *e.Channel) (*e.Channel, error)
	SaveRoleOverride(ctx context.Context, perm *e.ChannelRolePermissionOverride) (*e.ChannelRolePermissionOverride, error)
	SaveUserOverride(ctx context.Context, perm *e.ChannelUserPermissionOverride) (*e.ChannelUserPermissionOverride, error)

	Delete(ctx context.Context, id e.ChannelId) error
	DeleteRoleOverride(ctx context.Context, id e.ChannelId, roleId e.RoleId) error
	DeleteUserOverride(ctx context.Context, id e.ChannelId, userId e.UserId) error
}
