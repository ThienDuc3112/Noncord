package services

import "backend/internal/domain/entities"

type ModerationServices interface {
	BanUser(serverId entities.ServerId, userId entities.UserId) error
	UnbanUser(serverId entities.ServerId, userId entities.UserId) error
}
