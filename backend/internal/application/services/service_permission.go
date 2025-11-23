package services

import (
	"backend/internal/application/interfaces"
	"backend/internal/application/query"
	"backend/internal/domain/entities"
	"backend/internal/domain/repositories"
	"context"

	"github.com/google/uuid"
	"github.com/gookit/goutil/arrutil"
)

type PermissionRepos interface {
	Member() repositories.MemberRepo
	Server() repositories.ServerRepo
	Channel() repositories.ChannelRepo
}

type PermissionService struct {
	uow repositories.UnitOfWork[PermissionRepos]
}

func NewPermissionQueries(uow repositories.UnitOfWork[PermissionRepos]) interfaces.PermissionQueries {
	return &PermissionService{uow}
}

func (s *PermissionService) getChannelContext(ctx context.Context, repos PermissionRepos, channelId entities.ChannelId, userId entities.UserId) (*entities.Channel, *entities.Server, *entities.Membership, error) {
	channel, err := repos.Channel().Find(ctx, channelId)
	if err != nil {
		return nil, nil, nil, entities.GetErrOrDefault(err, entities.ErrCodeDepFail, "cannot get message's channel")
	}

	server, membership, err := s.getServerContext(ctx, repos, channel.ServerId, userId)
	if err != nil {
		return nil, nil, nil, err
	}

	return channel, server, membership, nil
}

func (s *PermissionService) getServerContext(ctx context.Context, repos PermissionRepos, serverId entities.ServerId, userId entities.UserId) (*entities.Server, *entities.Membership, error) {
	server, err := repos.Server().Find(ctx, serverId)
	if err != nil {
		return nil, nil, entities.GetErrOrDefault(err, entities.ErrCodeDepFail, "cannot get message's server")
	}

	membership, err := repos.Member().Find(ctx, userId, server.Id)
	if err != nil {
		if derr, ok := err.(*entities.ChatError); ok && derr.Code == entities.ErrCodeNoObject {
			return nil, nil, entities.NewError(entities.ErrCodeForbidden, "user not in server to view message", nil)
		}
		return nil, nil, entities.GetErrOrDefault(err, entities.ErrCodeDepFail, "cannot get user's server membership detail")
	}

	return server, membership, nil
}

func (s *PermissionService) getChannelEffectivePerm(ctx context.Context, channelId entities.ChannelId, userId entities.UserId) (res entities.ServerPermissionBits, err error) {
	err = s.uow.Do(ctx, func(ctx context.Context, repos PermissionRepos) error {
		_, _, _, err := s.getChannelContext(ctx, repos, channelId, userId)
		if err != nil {
			return err
		}

		// TODO: Actually calculate effective perm
		negOne := int64(-1)
		res = entities.ServerPermissionBits(negOne)

		return nil
	})

	return res, err
}

func (s *PermissionService) getServerEffectivePerm(ctx context.Context, serverId entities.ServerId, userId entities.UserId) (res entities.ServerPermissionBits, err error) {
	err = s.uow.Do(ctx, func(ctx context.Context, repos PermissionRepos) error {
		_, _, err := s.getServerContext(ctx, repos, serverId, userId)
		if err != nil {
			return err
		}

		// TODO: Actually calculate effective perm
		negOne := int64(-1)
		res = entities.ServerPermissionBits(negOne)

		return nil
	})

	return res, err
}

func (s *PermissionService) ChannelHasAll(ctx context.Context, params query.CheckChannelPerm) (res bool, err error) {
	effBit, err := s.getChannelEffectivePerm(ctx, entities.ChannelId(params.ChannelId), entities.UserId(params.UserId))
	if err != nil {
		return false, err
	}

	return effBit.HasAll(params.Permission), err
}

func (s *PermissionService) ChannelHasAny(ctx context.Context, params query.CheckChannelPerm) (bool, error) {
	effBit, err := s.getChannelEffectivePerm(ctx, entities.ChannelId(params.ChannelId), entities.UserId(params.UserId))
	if err != nil {
		return false, err
	}

	return effBit.HasAny(params.Permission), err
}

func (s *PermissionService) ServerHasAll(ctx context.Context, params query.CheckServerPerm) (bool, error) {
	effBit, err := s.getServerEffectivePerm(ctx, entities.ServerId(params.ServerId), entities.UserId(params.UserId))
	if err != nil {
		return false, err
	}

	return effBit.HasAll(params.Permission), err
}

func (s *PermissionService) ServerHasAny(ctx context.Context, params query.CheckServerPerm) (bool, error) {
	effBit, err := s.getServerEffectivePerm(ctx, entities.ServerId(params.ServerId), entities.UserId(params.UserId))
	if err != nil {
		return false, err
	}

	return effBit.HasAny(params.Permission), err
}

func (s *PermissionService) GetVisibleChannels(ctx context.Context, userId uuid.UUID) (res uuid.UUIDs, err error) {
	err = s.uow.Do(ctx, func(ctx context.Context, repos PermissionRepos) error {
		channelIds, err := repos.Channel().FindByUserServers(ctx, entities.UserId(userId))
		if err != nil {
			return entities.GetErrOrDefault(err, entities.ErrCodeDepFail, "cannot get channels")
		}

		res = arrutil.Map(channelIds, func(id entities.ChannelId) (uuid.UUID, bool) { return uuid.UUID(id), true })
		return nil
	})

	return res, err
}

func (s *PermissionService) GetVisibleServers(ctx context.Context, userId uuid.UUID) (uuid.UUIDs, error) {
	return nil, nil
}
