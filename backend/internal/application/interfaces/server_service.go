package interfaces

import (
	"backend/internal/application/command"
	"backend/internal/application/query"
	"context"
)

type ServerService interface {
	Create(context.Context, command.CreateServerCommand) (command.CreateServerCommandResult, error)
	Get(context.Context, query.GetServer) (query.GetServerResult, error)
	GetServers(context.Context, query.GetServers) (query.GetServersResult, error)
	GetServersUserIn(context.Context, query.GetServersUserIn) (query.GetServersUserInResult, error)
	Update(context.Context, command.UpdateServerCommand) (command.UpdateServerCommandResult, error)
	Delete(context.Context, command.DeleteServerCommand) error
}
