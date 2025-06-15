package services

import "backend/internal/domain/entities"

type PermissionService interface {
	EffectivePermissions(channelId entities.ChannelId, userId entities.UserId) entities.ServerPermissionBits
}
