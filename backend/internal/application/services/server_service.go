package services

import (
	"backend/internal/application/command"
	"backend/internal/application/interfaces"
	"backend/internal/application/mapper"
	"backend/internal/application/query"
	"backend/internal/domain/entities"
	"backend/internal/domain/repositories"
	"context"
	"fmt"
)

type ServerService struct {
	r repositories.ServerRepo
}

func NewServerService(sr repositories.ServerRepo) interfaces.ServerService {
	return &ServerService{
		r: sr,
	}
}

func (s *ServerService) Create(ctx context.Context, params command.CreateServerCommand) (command.CreateServerCommandResult, error) {
	server := entities.NewServer(entities.UserId(params.UserId), params.Name, "", "", "", false)
	_, err := s.r.Save(ctx, server)
	if err != nil {
		return command.CreateServerCommandResult{}, entities.GetErrOrDefault(err, entities.ErrCodeDepFail, "cannot save server failed")
	}

	return command.CreateServerCommandResult{
		Result: mapper.ServerToResult(server),
	}, nil
}

func (s *ServerService) Get(ctx context.Context, params query.GetServer) (query.GetServerResult, error) {
	server, err := s.r.Find(ctx, entities.ServerId(params.ServerId))
	if err != nil {
		return query.GetServerResult{}, entities.GetErrOrDefault(err, entities.ErrCodeDepFail, "cannot get server")
	}
	return query.GetServerResult{
		Result: mapper.ServerToResult(server),
	}, nil
}

func (s *ServerService) Update(context.Context, *entities.Server) (*entities.Server, error) {
	return nil, fmt.Errorf("unimplemented")
}

func (s *ServerService) Delete(ctx context.Context, name string) (*entities.Server, error) {
	return nil, fmt.Errorf("unimplemented")
}
