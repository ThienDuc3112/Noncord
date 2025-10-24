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

	"github.com/gookit/goutil/arrutil"
)

type ChannelService struct {
	cr repositories.ChannelRepo
	sr repositories.ServerRepo
	mr repositories.MemberRepo
}

func NewChannelService(cr repositories.ChannelRepo, sr repositories.ServerRepo, mr repositories.MemberRepo) interfaces.ChannelService {
	return &ChannelService{
		cr: cr,
		sr: sr,
		mr: mr,
	}
}

func (s *ChannelService) Create(ctx context.Context, params command.CreateChannelCommand) (command.CreateChannelCommandResult, error) {
	server, err := s.sr.Find(ctx, entities.ServerId(params.ServerId))
	if err != nil {
		return command.CreateChannelCommandResult{}, entities.GetErrOrDefault(err, entities.ErrCodeDepFail, "cannot get server details")
	}

	// TODO: Add role check
	if !server.IsOwner(entities.UserId(params.UserId)) {
		return command.CreateChannelCommandResult{}, entities.NewError(entities.ErrCodeForbidden, "user don't have permission to delete channel", err)
	}

	maxOrder, err := s.cr.GetServerMaxChannelOrder(ctx, entities.ServerId(params.ServerId))
	if err != nil {
		return command.CreateChannelCommandResult{}, entities.GetErrOrDefault(err, entities.ErrCodeDepFail, "cannot get server details (max ordering)")
	}

	channel, err := s.cr.Save(ctx,
		entities.NewChannel(params.Name, params.Description, entities.ServerId(params.ServerId), uint16(maxOrder)+1, (*entities.CategoryId)(params.ParentCategory)))
	if err != nil {
		return command.CreateChannelCommandResult{}, entities.GetErrOrDefault(err, entities.ErrCodeDepFail, "cannot save channel")
	}

	return command.CreateChannelCommandResult{
		Result: mapper.ChannelToResult(channel),
	}, nil
}

func (s *ChannelService) Get(ctx context.Context, params query.GetChannel) (query.GetChannelResult, error) {
	channel, err := s.cr.Find(ctx, entities.ChannelId(params.ChannelId))
	if err != nil {
		return query.GetChannelResult{}, entities.GetErrOrDefault(err, entities.ErrCodeDepFail, "cannot get channel")
	}

	// TODO: Check view permission
	_, err = s.mr.Find(ctx, entities.UserId(params.UserId), channel.ServerId)
	if err != nil {
		return query.GetChannelResult{}, entities.GetErrOrDefault(err, entities.ErrCodeNoObject, "channel not found")
	}

	return query.GetChannelResult{
		Result: mapper.ChannelToResult(channel),
	}, nil
}

func (s *ChannelService) GetChannelsByServer(ctx context.Context, params query.GetChannelsByServer) (query.GetChannelsByServerResult, error) {
	channels, err := s.cr.FindByServerId(ctx, entities.ServerId(params.ServerId))
	if err != nil {
		return query.GetChannelsByServerResult{}, entities.GetErrOrDefault(err, entities.ErrCodeDepFail, "cannot get channels")
	}

	return query.GetChannelsByServerResult{
		Result: arrutil.Map(channels, func(channel *entities.Channel) (target *common.Channel, find bool) {
			return mapper.ChannelToResult(channel), true
		}),
	}, nil
}

func (s *ChannelService) Update(ctx context.Context, params command.UpdateChannelCommand) (command.UpdateChannelCommandResult, error) {

	return command.UpdateChannelCommandResult{}, nil
}

func (s *ChannelService) Delete(ctx context.Context, params command.DeleteChannelCommand) error {
	channel, err := s.cr.Find(ctx, entities.ChannelId(params.ChannelId))
	if err != nil {
		return entities.GetErrOrDefault(err, entities.ErrCodeDepFail, "cannot get channel")
	}

	server, err := s.sr.Find(ctx, channel.ServerId)
	if err != nil {
		return entities.GetErrOrDefault(err, entities.ErrCodeDepFail, "cannot get server details")
	}

	// TODO: Add role check
	if !server.IsOwner(entities.UserId(params.UserId)) {
		return entities.NewError(entities.ErrCodeForbidden, "user don't have permission to delete channel", err)
	}

	channel.Delete()
	channel, err = s.cr.Save(ctx, channel)
	return entities.GetErrOrDefault(err, entities.ErrCodeDepFail, "cannot delete server")
}
