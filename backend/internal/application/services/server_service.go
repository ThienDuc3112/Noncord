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

	// TODO: Create default channel and role

	return query.GetServerResult{
		Result: mapper.ServerToResult(server),
	}, nil
}

func (s *ServerService) Update(ctx context.Context, params command.UpdateServerCommand) error {
	server, err := s.r.Find(ctx, entities.ServerId(params.ServerId))
	if err != nil {
		return entities.GetErrOrDefault(err, entities.ErrCodeDepFail, "cannot get server")
	}

	if server.Owner != entities.UserId(params.UserId) {
		return entities.NewError(entities.ErrCodeForbidden, "not authorized", err)
	}
	// TODO: Update with mod and role permission

	if params.Updates.Name != "" {
		server.Name = params.Updates.Name
	}
	if params.Updates.Description != "" {
		server.Description = params.Updates.Description
	}
	if params.Updates.IconUrl != "" {
		server.IconUrl = params.Updates.IconUrl
	}
	if params.Updates.BannerUrl != "" {
		server.BannerUrl = params.Updates.BannerUrl
	}
	if params.Updates.NeedApproval != server.NeedApproval {
		server.NeedApproval = params.Updates.NeedApproval
	}
	if params.Updates.DefaultRole.Valid {
		server.DefaultRole = (*entities.RoleId)(&params.Updates.DefaultRole.UUID)
	}
	if params.Updates.AnnouncementChannel.Valid {
		server.AnnouncementChannel = (*entities.ChannelId)(&params.Updates.AnnouncementChannel.UUID)
	}

	if err = server.Validate(); err != nil {
		return err
	}

	_, err = s.r.Save(ctx, server)
	return entities.GetErrOrDefault(err, entities.ErrCodeDepFail, "cannot update server")
}

func (s *ServerService) Delete(ctx context.Context, name string) (*entities.Server, error) {

	return nil, fmt.Errorf("unimplemented")
}
