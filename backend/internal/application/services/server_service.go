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

	// TODO: Update with mod and role permission
	if !server.IsOwner(entities.UserId(params.UserId)) {
		return entities.NewError(entities.ErrCodeForbidden, "not authorized", err)
	}

	if params.Updates.Name != nil {
		if err = server.UpdateName(*params.Updates.Name); err != nil {
			return err
		}
	}
	if params.Updates.Description != nil {
		if err = server.UpdateDescription(*params.Updates.Description); err != nil {
			return err
		}
	}
	if params.Updates.IconUrl != nil {
		if err = server.UpdateIconUrl(*params.Updates.IconUrl); err != nil {
			return err
		}
	}
	if params.Updates.BannerUrl != nil {
		if err = server.UpdateBannerUrl(*params.Updates.BannerUrl); err != nil {
			return err
		}
	}
	if params.Updates.NeedApproval != nil {
		if err = server.UpdateNeedApproval(*params.Updates.NeedApproval); err != nil {
			return err
		}
	}
	if params.Updates.AnnouncementChannel.Valid {
		if err = server.UpdateAnnouncementChannel((*entities.ChannelId)(&params.Updates.AnnouncementChannel.UUID)); err != nil {
			return err
		}
	}

	_, err = s.r.Save(ctx, server)
	return entities.GetErrOrDefault(err, entities.ErrCodeDepFail, "cannot update server")
}

func (s *ServerService) Delete(ctx context.Context, param command.DeleteServerCommand) error {
	server, err := s.r.Find(ctx, entities.ServerId(param.ServerId))
	if err != nil {
		return entities.GetErrOrDefault(err, entities.ErrCodeDepFail, "cannot get server")
	}
	if server.Owner != entities.UserId(param.UserId) {
		return entities.NewError(entities.ErrCodeForbidden, "user is not the owner of the server", nil)
	}

	return fmt.Errorf("unimplemented")
}
