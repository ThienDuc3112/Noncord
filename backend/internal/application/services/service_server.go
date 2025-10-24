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

type ServerRepos interface {
	Channel() repositories.ChannelRepo
	Server() repositories.ServerRepo
	Member() repositories.MemberRepo
}

type ServerService struct {
	uow repositories.UnitOfWork[ServerRepos]
}

func NewServerService(uow repositories.UnitOfWork[ServerRepos]) interfaces.ServerService {
	return &ServerService{uow}
}

func (s *ServerService) Create(ctx context.Context, params command.CreateServerCommand) (res command.CreateServerCommandResult, err error) {
	server, err := entities.NewServer(entities.UserId(params.UserId), params.Name, "", "", "", false)
	if err != nil {
		return command.CreateServerCommandResult{}, err
	}

	err = s.uow.Do(ctx, func(ctx context.Context, repos ServerRepos) error {
		server, err = repos.Server().Save(ctx, server)
		if err != nil {
			return entities.GetErrOrDefault(err, entities.ErrCodeDepFail, "cannot save server")
		}

		_, err = repos.Member().Save(ctx, entities.NewMembership(server.Id, entities.UserId(params.UserId), params.UserDisplayName))
		if err != nil {
			return entities.GetErrOrDefault(err, entities.ErrCodeDepFail, "cannot save membership")
		}

		channel := entities.NewChannel("text channel", "Your first channel", server.Id, 1, nil)
		channel, err = repos.Channel().Save(ctx, channel)
		if err != nil {
			return entities.GetErrOrDefault(err, entities.ErrCodeDepFail, "cannot save channel")
		}
		// TODO: Create default role

		if err = server.UpdateAnnouncementChannel(&channel.Id); err != nil {
			return err
		}

		server, err = repos.Server().Save(ctx, server)
		if err != nil {
			return entities.GetErrOrDefault(err, entities.ErrCodeDepFail, "cannot save server")
		}

		res = command.CreateServerCommandResult{
			Result: mapper.ServerToResult(server),
		}
		return nil
	})

	return res, err
}

func (s *ServerService) Get(ctx context.Context, params query.GetServer) (res query.GetServerResult, err error) {
	err = s.uow.Do(ctx, func(ctx context.Context, repos ServerRepos) error {
		server, err := repos.Server().Find(ctx, entities.ServerId(params.ServerId))
		if err != nil {
			return entities.GetErrOrDefault(err, entities.ErrCodeDepFail, "cannot get server")
		}

		if params.UserId == nil {
			res = query.GetServerResult{
				Preview: mapper.ServerToPreview(server),
			}
			return nil
		}

		if _, err = repos.Member().Find(ctx, entities.UserId(*params.UserId), server.Id); err != nil {
			res = query.GetServerResult{
				Preview: mapper.ServerToPreview(server),
			}
			return nil
		}

		channels, err := repos.Channel().FindByServerId(ctx, server.Id)
		if err != nil {
			return entities.GetErrOrDefault(err, entities.ErrCodeDepFail, "cannot get channels")
		}

		res = query.GetServerResult{
			Preview: mapper.ServerToPreview(server),
			Full:    mapper.ServerToResult(server),
			Channel: arrutil.Map(channels, func(c *entities.Channel) (target *common.Channel, find bool) {
				return mapper.ChannelToResult(c), true
			}),
		}
		return nil
	})

	return res, err
}

func (s *ServerService) GetServers(ctx context.Context, params query.GetServers) (res query.GetServersResult, err error) {
	var mapFn arrutil.MapFn[uuid.UUID, entities.ServerId] = func(input uuid.UUID) (target entities.ServerId, find bool) {
		return entities.ServerId(input), true
	}

	err = s.uow.Do(ctx, func(ctx context.Context, repos ServerRepos) error {
		servers, err := repos.Server().FindByIds(ctx, arrutil.Map(params.ServerIds, mapFn))
		if err != nil {
			return entities.GetErrOrDefault(err, entities.ErrCodeDepFail, "cannot get server")
		}

		res = query.GetServersResult{
			Result: arrutil.Map(servers, func(server *entities.Server) (target *common.Server, find bool) {
				return mapper.ServerToResult(server), true
			}),
		}
		return nil
	})

	return res, err
}

func (s *ServerService) GetServersUserIn(ctx context.Context, params query.GetServersUserIn) (res query.GetServersUserInResult, err error) {
	err = s.uow.Do(ctx, func(ctx context.Context, repos ServerRepos) error {
		servers, err := repos.Server().FindByUser(ctx, entities.UserId(params.UserId))
		if err != nil {
			return entities.GetErrOrDefault(err, entities.ErrCodeDepFail, "cannot get server")
		}

		res = query.GetServersUserInResult{
			Result: arrutil.Map(servers, func(server *entities.Server) (target *common.Server, find bool) {
				return mapper.ServerToResult(server), true
			}),
		}
		return nil
	})

	return res, err
}

func (s *ServerService) Update(ctx context.Context, params command.UpdateServerCommand) (res command.UpdateServerCommandResult, err error) {
	err = s.uow.Do(ctx, func(ctx context.Context, repos ServerRepos) error {
		server, err := repos.Server().Find(ctx, entities.ServerId(params.ServerId))
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
		if params.Updates.DefaultPermission != nil {
			if err = server.UpdateDefaultPermission(entities.ServerPermissionBits(*params.Updates.DefaultPermission)); err != nil {
				return err
			}
		}

		server, err = repos.Server().Save(ctx, server)
		if err != nil {
			return entities.GetErrOrDefault(err, entities.ErrCodeDepFail, "cannot update server")
		}

		res = command.UpdateServerCommandResult{
			Result: mapper.ServerToResult(server),
		}
		return nil
	})

	return res, err
}

func (s *ServerService) Delete(ctx context.Context, param command.DeleteServerCommand) error {
	return s.uow.Do(ctx, func(ctx context.Context, repos ServerRepos) error {
		server, err := repos.Server().Find(ctx, entities.ServerId(param.ServerId))
		if err != nil {
			return entities.GetErrOrDefault(err, entities.ErrCodeDepFail, "cannot get server")
		}

		if server.Owner != entities.UserId(param.UserId) {
			return entities.NewError(entities.ErrCodeForbidden, "user is not the owner of the server", nil)
		}

		server.Delete()
		server, err = repos.Server().Save(ctx, server)
		return entities.GetErrOrDefault(err, entities.ErrCodeDepFail, "cannot delete server")
	})
}
