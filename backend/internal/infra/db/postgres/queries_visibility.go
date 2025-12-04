package postgres

import (
	"backend/internal/application/interfaces"
	"backend/internal/application/query"
	"backend/internal/domain/entities"
	"backend/internal/infra/db/postgres/gen"
	"context"

	"github.com/google/uuid"
	"github.com/gookit/goutil/arrutil"
	"github.com/jackc/pgx/v5/pgxpool"
)

type PGVisibilityQueries struct {
	q *gen.Queries
}

func NewPGVisibilityQueries(conn *pgxpool.Conn) interfaces.VisibilityQueries {
	return &PGVisibilityQueries{
		q: gen.New(conn),
	}
}

func (q *PGVisibilityQueries) GetVisibleChannels(ctx context.Context, userId uuid.UUID) (uuid.UUIDs, error) {
	channels, err := q.q.FindAllChannelInUserServers(ctx, userId)
	if err != nil {
		return nil, entities.NewError(entities.ErrCodeDepFail, "cannot get channels", err)
	}

	rolesPerm, err := q.q.FindAllUserRolePermission(ctx, userId)
	if err != nil {
		return nil, entities.NewError(entities.ErrCodeDepFail, "cannot get user permission (roles)", err)
	}

	roleOverwrite, err := q.q.FindAllChannelUserRoleOverwrite(ctx, userId)
	if err != nil {
		return nil, entities.NewError(entities.ErrCodeDepFail, "cannot get channel permission (role)", err)
	}

	userOverwrite, err := q.q.FindAllChannelUserOverwrite(ctx, userId)
	if err != nil {
		return nil, entities.NewError(entities.ErrCodeDepFail, "cannot get channel permission (user)", err)
	}

	chanEffPerms := effectivePermPerChannel(getChannelsHelperParams{
		Channels: arrutil.Map(channels, func(row gen.FindAllChannelInUserServersRow) (target channel, find bool) {
			return channel{id: row.ID, serverId: row.ServerID}, true
		}),
		RolePerms: arrutil.Map(rolesPerm, func(row gen.FindAllUserRolePermissionRow) (target rolePerm, find bool) {
			return rolePerm{serverId: row.ServerID, roleId: row.RoleID, permission: row.Permissions}, true
		}),
		UserOverwrite: arrutil.Map(userOverwrite, func(row gen.FindAllChannelUserOverwriteRow) (target overwrite, find bool) {
			return overwrite{allow: row.Allow, deny: row.Deny, channelId: row.ChannelID, targetId: userId}, true
		}),
		RoleOverwrite: arrutil.Map(roleOverwrite, func(row gen.FindAllChannelUserRoleOverwriteRow) (target overwrite, find bool) {
			return overwrite{allow: row.Allow, deny: row.Deny, channelId: row.ChannelID, targetId: row.RoleID}, true
		}),
	})

	res := arrutil.Map(channels, func(c gen.FindAllChannelInUserServersRow) (uuid.UUID, bool) {
		return c.ID, entities.ServerPermissionBits(chanEffPerms[c.ID]).HasAny(entities.PermViewChannel, entities.PermAdministrator)
	})

	return res, nil

}

func (q *PGVisibilityQueries) GetVisibleChannelsInServer(ctx context.Context, params query.GetVisibleChannelsInServer) (uuid.UUIDs, error) {
	channels, err := q.q.FindChannelsByServerId(ctx, params.ServerId)
	if err != nil {
		return nil, entities.NewError(entities.ErrCodeDepFail, "cannot get channels", err)
	}

	rolesPerm, err := q.q.FindAllUserServerRolePermission(ctx, gen.FindAllUserServerRolePermissionParams{
		UserID:   params.UserId,
		ServerID: params.ServerId,
	})
	if err != nil {
		return nil, entities.NewError(entities.ErrCodeDepFail, "cannot get user permission (roles)", err)
	}

	roleOverwrite, err := q.q.FindAllChannelUserServerRoleOverwrite(ctx, gen.FindAllChannelUserServerRoleOverwriteParams{
		UserID:   params.UserId,
		ServerID: params.ServerId,
	})
	if err != nil {
		return nil, entities.NewError(entities.ErrCodeDepFail, "cannot get channel permission (role)", err)
	}

	userOverwrite, err := q.q.FindAllChannelServerUserOverwrite(ctx, gen.FindAllChannelServerUserOverwriteParams{
		UserID:   params.UserId,
		ServerID: params.ServerId,
	})
	if err != nil {
		return nil, entities.NewError(entities.ErrCodeDepFail, "cannot get channel permission (user)", err)
	}

	chanEffPerms := effectivePermPerChannel(getChannelsHelperParams{
		Channels: arrutil.Map(channels, func(c gen.Channel) (target channel, find bool) {
			return channel{id: c.ID, serverId: c.ServerID}, true
		}),
		RolePerms: arrutil.Map(rolesPerm, func(row gen.FindAllUserServerRolePermissionRow) (target rolePerm, find bool) {
			return rolePerm{serverId: row.ServerID, roleId: row.RoleID, permission: row.Permissions}, true
		}),
		UserOverwrite: arrutil.Map(userOverwrite, func(row gen.FindAllChannelServerUserOverwriteRow) (target overwrite, find bool) {
			return overwrite{allow: row.Allow, deny: row.Deny, channelId: row.ChannelID, targetId: params.UserId}, true
		}),
		RoleOverwrite: arrutil.Map(roleOverwrite, func(row gen.FindAllChannelUserServerRoleOverwriteRow) (target overwrite, find bool) {
			return overwrite{allow: row.Allow, deny: row.Deny, channelId: row.ChannelID, targetId: row.RoleID}, true
		}),
	})

	res := arrutil.Map(channels, func(c gen.Channel) (uuid.UUID, bool) {
		return c.ID, entities.ServerPermissionBits(chanEffPerms[c.ID]).HasAny(entities.PermViewChannel, entities.PermAdministrator)
	})

	return res, nil
}

func (q *PGVisibilityQueries) GetVisibleServers(ctx context.Context, userId uuid.UUID) (uuid.UUIDs, error) {
	// TODO: here
	return nil, nil
}
