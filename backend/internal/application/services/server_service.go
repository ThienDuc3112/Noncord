package services

import (
	"backend/internal/application/command"
	"backend/internal/application/interfaces"
	"backend/internal/application/query"
	"backend/internal/domain/entities"
	"context"
	"fmt"
)

type ServerService struct {
}

func NewServerService() interfaces.ServerService {
	return &ServerService{}
}

func (s *ServerService) Create(context.Context, command.CreateServerCommand) (command.CreateServerCommandResult, error) {
	return command.CreateServerCommandResult{}, fmt.Errorf("unimplemented")
}

func (s *ServerService) Get(context.Context, query.GetServer) (query.GetServerResult, error) {
	return query.GetServerResult{}, fmt.Errorf("unimplemented")
}

func (s *ServerService) Update(context.Context, *entities.Server) (*entities.Server, error) {
	return nil, fmt.Errorf("unimplemented")
}

func (s *ServerService) Delete(ctx context.Context, name string) (*entities.Server, error) {
	return nil, fmt.Errorf("unimplemented")
}
