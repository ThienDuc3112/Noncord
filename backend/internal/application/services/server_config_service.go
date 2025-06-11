package services

import "backend/internal/domain/entities"

type ServerConfigService interface {
	CreateServer(serverName string) (*entities.Server, error)
	UpdateServer(server *entities.Server) error
	DeleteServer(serverId entities.ServerId) error

	CreateCategory(serverId entities.ServerId, name string) (*entities.Category, error)
	UpdateCategory(category *entities.Category) error
	DeleteCategory(categoryId entities.CategoryId) error

	CreateRole(serverId entities.ServerId, name string) (*entities.Role, error)
	UpdateRole(role *entities.Role) error
	DeleteRole(roleId entities.RoleId) error
}
