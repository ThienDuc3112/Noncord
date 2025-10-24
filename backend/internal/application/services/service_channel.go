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
	"fmt"

	"github.com/gookit/goutil/arrutil"
)

type ChannelRepos interface {
	Channel() repositories.ChannelRepo
	Server() repositories.ServerRepo
	Member() repositories.MemberRepo
}

type ChannelService struct {
	uow repositories.UnitOfWork[ChannelRepos]
}

func NewChannelService(uow repositories.UnitOfWork[ChannelRepos]) interfaces.ChannelService {
	return &ChannelService{uow: uow}
}

func (s *ChannelService) Create(ctx context.Context, params command.CreateChannelCommand) (res command.CreateChannelCommandResult, err error) {
	err = s.uow.Do(ctx, func(ctx context.Context, repos ChannelRepos) error {
		server, err := repos.Server().Find(ctx, entities.ServerId(params.ServerId))
		if err != nil {
			return entities.GetErrOrDefault(err, entities.ErrCodeDepFail, "cannot get server details")
		}

		// TODO: Add role check
		if !server.IsOwner(entities.UserId(params.UserId)) {
			return entities.NewError(entities.ErrCodeForbidden, "user don't have permission to delete channel", err)
		}

		maxOrder, err := repos.Channel().GetServerMaxChannelOrder(ctx, entities.ServerId(params.ServerId))
		if err != nil {
			return entities.GetErrOrDefault(err, entities.ErrCodeDepFail, "cannot get server details (max ordering)")
		}

		channel, err := repos.Channel().Save(ctx,
			entities.NewChannel(params.Name, params.Description, entities.ServerId(params.ServerId), uint16(maxOrder)+1, (*entities.CategoryId)(params.ParentCategory)))
		if err != nil {
			return entities.GetErrOrDefault(err, entities.ErrCodeDepFail, "cannot save channel")
		}

		res = command.CreateChannelCommandResult{
			Result: mapper.ChannelToResult(channel),
		}
		return nil
	})

	return res, err
}

func (s *ChannelService) Get(ctx context.Context, params query.GetChannel) (res query.GetChannelResult, err error) {
	err = s.uow.Do(ctx, func(ctx context.Context, repos ChannelRepos) error {
		channel, err := repos.Channel().Find(ctx, entities.ChannelId(params.ChannelId))
		if err != nil {
			return entities.GetErrOrDefault(err, entities.ErrCodeDepFail, "cannot get channel")
		}

		// TODO: Check view permission
		_, err = repos.Member().Find(ctx, entities.UserId(params.UserId), channel.ServerId)
		if err != nil {
			return entities.GetErrOrDefault(err, entities.ErrCodeNoObject, "channel not found")
		}

		res = query.GetChannelResult{
			Result: mapper.ChannelToResult(channel),
		}
		return nil
	})

	return res, err
}

func (s *ChannelService) GetChannelsByServer(ctx context.Context, params query.GetChannelsByServer) (res query.GetChannelsByServerResult, err error) {
	err = s.uow.Do(ctx, func(ctx context.Context, repos ChannelRepos) error {
		channels, err := repos.Channel().FindByServerId(ctx, entities.ServerId(params.ServerId))
		if err != nil {
			return entities.GetErrOrDefault(err, entities.ErrCodeDepFail, "cannot get channels")
		}

		res = query.GetChannelsByServerResult{
			Result: arrutil.Map(channels, func(channel *entities.Channel) (target *common.Channel, find bool) {
				return mapper.ChannelToResult(channel), true
			}),
		}
		return nil
	})

	return res, err
}

func (s *ChannelService) Update(ctx context.Context, params command.UpdateChannelCommand) (command.UpdateChannelCommandResult, error) {
	return command.UpdateChannelCommandResult{}, fmt.Errorf("not implemented")
}

func (s *ChannelService) Delete(ctx context.Context, params command.DeleteChannelCommand) error {
	return s.uow.Do(ctx, func(ctx context.Context, repos ChannelRepos) error {
		channel, err := repos.Channel().Find(ctx, entities.ChannelId(params.ChannelId))
		if err != nil {
			return entities.GetErrOrDefault(err, entities.ErrCodeDepFail, "cannot get channel")
		}

		server, err := repos.Server().Find(ctx, channel.ServerId)
		if err != nil {
			return entities.GetErrOrDefault(err, entities.ErrCodeDepFail, "cannot get server details")
		}

		// TODO: Add role check
		if !server.IsOwner(entities.UserId(params.UserId)) {
			return entities.NewError(entities.ErrCodeForbidden, "user don't have permission to delete channel", err)
		}

		channel.Delete()
		channel, err = repos.Channel().Save(ctx, channel)
		return entities.GetErrOrDefault(err, entities.ErrCodeDepFail, "cannot delete server")
	})
}
