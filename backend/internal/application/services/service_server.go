package services

import (
	"backend/internal/application/command"
	"backend/internal/application/common"
	"backend/internal/application/interfaces"
	"backend/internal/application/mapper"
	"backend/internal/application/query"
	"backend/internal/domain/entities"
	"backend/internal/domain/repositories"
	"context"

	"github.com/google/uuid"
	"github.com/gookit/goutil/arrutil"
)

type ServerService struct {
	sr repositories.ServerRepo
	mr repositories.MemberRepo
	cr repositories.ChannelRepo
}

func NewServerService(sr repositories.ServerRepo, mr repositories.MemberRepo, cr repositories.ChannelRepo) interfaces.ServerService {
	return &ServerService{
		sr: sr,
		mr: mr,
		cr: cr,
	}
}

func (s *ServerService) Create(ctx context.Context, params command.CreateServerCommand) (command.CreateServerCommandResult, error) {
	server, err := entities.NewServer(entities.UserId(params.UserId), params.Name, "", "", "", false)
	if err != nil {
		return command.CreateServerCommandResult{}, err
	}

	server, err = s.sr.Save(ctx, server)
	if err != nil {
		return command.CreateServerCommandResult{}, entities.GetErrOrDefault(err, entities.ErrCodeDepFail, "cannot save server")
	}

	_, err = s.mr.Save(ctx, entities.NewMembership(server.Id, entities.UserId(params.UserId), params.UserDisplayName))
	if err != nil {
		return command.CreateServerCommandResult{}, entities.GetErrOrDefault(err, entities.ErrCodeDepFail, "cannot save membership")
	}

	channel := entities.NewChannel("text channel", "Your first channel", server.Id, 1, nil)
	channel, err = s.cr.Save(ctx, channel)
	if err != nil {
		return command.CreateServerCommandResult{}, entities.GetErrOrDefault(err, entities.ErrCodeDepFail, "cannot save channel")
	}
	// TODO: Create default role

	if err = server.UpdateAnnouncementChannel(&channel.Id); err != nil {
		return command.CreateServerCommandResult{}, err
	}

	server, err = s.sr.Save(ctx, server)
	if err != nil {
		return command.CreateServerCommandResult{}, entities.GetErrOrDefault(err, entities.ErrCodeDepFail, "cannot save server")
	}

	return command.CreateServerCommandResult{
		Result: mapper.ServerToResult(server),
	}, nil
}

func (s *ServerService) Get(ctx context.Context, params query.GetServer) (query.GetServerResult, error) {
	server, err := s.sr.Find(ctx, entities.ServerId(params.ServerId))
	if err != nil {
		return query.GetServerResult{}, entities.GetErrOrDefault(err, entities.ErrCodeDepFail, "cannot get server")
	}

	if params.UserId == nil {
		return query.GetServerResult{
			Preview: mapper.ServerToPreview(server),
		}, nil
	}

	if _, err = s.mr.Find(ctx, entities.UserId(*params.UserId), server.Id); err != nil {
		return query.GetServerResult{
			Preview: mapper.ServerToPreview(server),
		}, nil
	}

	channels, err := s.cr.FindByServerId(ctx, server.Id)
	if err != nil {
		return query.GetServerResult{}, entities.GetErrOrDefault(err, entities.ErrCodeDepFail, "cannot get channels")
	}

	return query.GetServerResult{
		Preview: mapper.ServerToPreview(server),
		Full:    mapper.ServerToResult(server),
		Channel: arrutil.Map(channels, func(c *entities.Channel) (target *common.Channel, find bool) {
			return mapper.ChannelToResult(c), true
		}),
	}, nil
}

func (s *ServerService) GetServers(ctx context.Context, params query.GetServers) (query.GetServersResult, error) {
	var mapFn arrutil.MapFn[uuid.UUID, entities.ServerId] = func(input uuid.UUID) (target entities.ServerId, find bool) {
		return entities.ServerId(input), true
	}
	servers, err := s.sr.FindByIds(ctx, arrutil.Map(params.ServerIds, mapFn))
	if err != nil {
		return query.GetServersResult{}, entities.GetErrOrDefault(err, entities.ErrCodeDepFail, "cannot get server")
	}

	return query.GetServersResult{
		Result: arrutil.Map(servers, func(server *entities.Server) (target *common.Server, find bool) {
			return mapper.ServerToResult(server), true
		}),
	}, nil
}

func (s *ServerService) GetServersUserIn(ctx context.Context, params query.GetServersUserIn) (query.GetServersUserInResult, error) {
	servers, err := s.sr.FindByUser(ctx, entities.UserId(params.UserId))
	if err != nil {
		return query.GetServersUserInResult{}, entities.GetErrOrDefault(err, entities.ErrCodeDepFail, "cannot get server")
	}

	return query.GetServersUserInResult{
		Result: arrutil.Map(servers, func(server *entities.Server) (target *common.Server, find bool) {
			return mapper.ServerToResult(server), true
		}),
	}, nil
}

func (s *ServerService) Update(ctx context.Context, params command.UpdateServerCommand) (command.UpdateServerCommandResult, error) {
	server, err := s.sr.Find(ctx, entities.ServerId(params.ServerId))
	if err != nil {
		return command.UpdateServerCommandResult{}, entities.GetErrOrDefault(err, entities.ErrCodeDepFail, "cannot get server")
	}

	// TODO: Update with mod and role permission
	if !server.IsOwner(entities.UserId(params.UserId)) {
		return command.UpdateServerCommandResult{}, entities.NewError(entities.ErrCodeForbidden, "not authorized", err)
	}

	if params.Updates.Name != nil {
		if err = server.UpdateName(*params.Updates.Name); err != nil {
			return command.UpdateServerCommandResult{}, err
		}
	}
	if params.Updates.Description != nil {
		if err = server.UpdateDescription(*params.Updates.Description); err != nil {
			return command.UpdateServerCommandResult{}, err
		}
	}
	if params.Updates.IconUrl != nil {
		if err = server.UpdateIconUrl(*params.Updates.IconUrl); err != nil {
			return command.UpdateServerCommandResult{}, err
		}
	}
	if params.Updates.BannerUrl != nil {
		if err = server.UpdateBannerUrl(*params.Updates.BannerUrl); err != nil {
			return command.UpdateServerCommandResult{}, err
		}
	}
	if params.Updates.NeedApproval != nil {
		if err = server.UpdateNeedApproval(*params.Updates.NeedApproval); err != nil {
			return command.UpdateServerCommandResult{}, err
		}
	}
	if params.Updates.AnnouncementChannel.Valid {
		if err = server.UpdateAnnouncementChannel((*entities.ChannelId)(&params.Updates.AnnouncementChannel.UUID)); err != nil {
			return command.UpdateServerCommandResult{}, err
		}
	}
	if params.Updates.DefaultPermission != nil {
		if err = server.UpdateDefaultPermission(entities.ServerPermissionBits(*params.Updates.DefaultPermission)); err != nil {
			return command.UpdateServerCommandResult{}, err
		}
	}

	server, err = s.sr.Save(ctx, server)
	if err != nil {
		return command.UpdateServerCommandResult{}, entities.GetErrOrDefault(err, entities.ErrCodeDepFail, "cannot update server")
	}

	return command.UpdateServerCommandResult{
		Result: mapper.ServerToResult(server),
	}, nil
}

func (s *ServerService) Delete(ctx context.Context, param command.DeleteServerCommand) error {
	server, err := s.sr.Find(ctx, entities.ServerId(param.ServerId))
	if err != nil {
		return entities.GetErrOrDefault(err, entities.ErrCodeDepFail, "cannot get server")
	}

	if server.Owner != entities.UserId(param.UserId) {
		return entities.NewError(entities.ErrCodeForbidden, "user is not the owner of the server", nil)
	}

	server.Delete()
	server, err = s.sr.Save(ctx, server)
	return entities.GetErrOrDefault(err, entities.ErrCodeDepFail, "cannot delete server")
}
