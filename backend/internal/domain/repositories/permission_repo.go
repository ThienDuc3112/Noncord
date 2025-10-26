package repositories

import (
	"backend/internal/domain/entities"
	"context"
)

type UserChannelPermissionResult struct {
	ServerOwnerId     entities.UserId
	ServerDefaultPerm entities.ServerPermissionBits
	AssignedRoles     []entities.Role
	RoleOverwrite     []entities.ChannelPermOverwrite
	UserOverwrite     *entities.ChannelPermOverwrite
}

type PermissionRepo interface {
	GetUserChannelPermission(context.Context, entities.ChannelId, entities.UserId) (UserChannelPermissionResult, error)
}
