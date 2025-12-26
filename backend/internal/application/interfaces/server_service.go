package interfaces

import (
	"backend/internal/application/command"
	"backend/internal/application/query"
	"context"

	"github.com/google/uuid"
)

type ServerService interface {
	Create(context.Context, command.CreateServerCommand) (command.CreateServerCommandResult, error)
	UpdateMetadata(context.Context, command.UpdateServerCommand) (command.UpdateServerCommandResult, error)
	UpsertRole(context.Context, command.UpsertRoleCommand) (command.UpsertRoleCommandResult, error)
	ReorderRoles(context.Context) error
	Delete(context.Context, command.DeleteServerCommand) error
	DeleteRole(context.Context, command.DeleteRoleCommand) error
}

type ServerQueries interface {
	Get(context.Context, query.GetServer) (query.GetServerResult, error)
	GetServers(context.Context, query.GetServers) (query.GetServersResult, error)
	GetServersUserIn(context.Context, query.GetServersUserIn) (query.GetServersUserInResult, error)
	GetServerIdsUserIn(context.Context, query.GetServersUserIn) (uuid.UUIDs, error)
}
