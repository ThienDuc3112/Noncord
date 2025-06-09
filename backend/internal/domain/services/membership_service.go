package services

import "backend/internal/domain/entities"

type MembershipService interface {
	JoinServer(serverId entities.ServerId, userId entities.UserId, inviteId entities.InvititationId) error
	LeaveServer(serverId entities.ServerId, userId entities.UserId) error
	TransferOwnership(serverId entities.ServerId, userId entities.UserId)
}
