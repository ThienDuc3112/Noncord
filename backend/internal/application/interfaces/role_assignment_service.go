package services

import "backend/internal/domain/entities"

type RoleAssignmentService interface {
	Assign(serverId entities.ServerId, userId entities.UserId, roleId entities.RoleId) error
	Remove(serverId entities.ServerId, userId entities.UserId, roleId entities.RoleId) error
}
