package services

import "backend/internal/domain/entities"

type ChannelService interface {
	CreateChannel(serverId entities.ServerId, name string) (entities.Channel, error)
	AddRoleOverride(channelId entities.ChannelId, roleId entities.RoleId, perm entities.ServerPermissionBits) (*entities.ChannelRolePermissionOverride, error)
}
