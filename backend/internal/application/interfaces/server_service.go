package interfaces

import (
	"backend/internal/application/command"
	"backend/internal/application/query"
	"backend/internal/domain/entities"
	"context"
)

type ServerService interface {
	Create(context.Context, command.CreateServerCommand) (command.CreateServerCommandResult, error)
	Get(context.Context, query.GetServer) (query.GetServerResult, error)
	Update(context.Context, *entities.Server) (*entities.Server, error)
	Delete(ctx context.Context, name string) (*entities.Server, error)
}
