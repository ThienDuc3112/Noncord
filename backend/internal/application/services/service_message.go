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
	"time"

	"github.com/gookit/goutil/arrutil"
)

type MessageRepos interface {
	Member() repositories.MemberRepo
	Server() repositories.ServerRepo
	Channel() repositories.ChannelRepo
	Message() repositories.MessageRepo
}

type MessageService struct {
	uow repositories.UnitOfWork[MessageRepos]
}

func NewMessageService(uow repositories.UnitOfWork[MessageRepos]) interfaces.MessageService {
	return &MessageService{uow}
}

func (s *MessageService) getChannelContext(ctx context.Context, repos MessageRepos, channelId entities.ChannelId, userId entities.UserId) (*entities.Channel, *entities.Server, *entities.Membership, *entities.ChatError) {
	channel, err := repos.Channel().Find(ctx, channelId)
	if err != nil {
		return nil, nil, nil, entities.GetErrOrDefault(err, entities.ErrCodeDepFail, "cannot get message's channel")
	}

	server, err := repos.Server().Find(ctx, channel.ServerId)
	if err != nil {
		return nil, nil, nil, entities.GetErrOrDefault(err, entities.ErrCodeDepFail, "cannot get message's server")
	}

	membership, err := repos.Member().Find(ctx, userId, server.Id)
	if err != nil {
		if derr, ok := err.(*entities.ChatError); ok && derr.Code == entities.ErrCodeNoObject {
			return nil, nil, nil, entities.NewError(entities.ErrCodeForbidden, "user not in server to view message", nil)
		}
		return nil, nil, nil, entities.GetErrOrDefault(err, entities.ErrCodeDepFail, "cannot get user's server membership detail")
	}

	return channel, server, membership, nil
}

func (s *MessageService) Get(ctx context.Context, params query.GetMessage) (res query.GetMessageResult, err error) {
	err = s.uow.Do(ctx, func(ctx context.Context, repos MessageRepos) error {
		msg, err := repos.Message().Find(ctx, entities.MessageId(params.MessageId))
		if err != nil {
			return entities.GetErrOrDefault(err, entities.ErrCodeDepFail, "cannot get message")
		}

		if msg.ChannelId != nil {
			_, _, _, err := s.getChannelContext(ctx, repos, *msg.ChannelId, entities.UserId(params.UserId))
			if err != nil {
				return err
			}
			// TODO: Check permission with roles, channel overwrite and stuff

			res = query.GetMessageResult{
				Result: mapper.MessageToResult(msg),
			}
			return nil
		} else {
			return entities.NewError(entities.ErrCodeForbidden, "dm group not implemented", nil)
		}
	})

	return res, err
}

func (s *MessageService) GetByGroupId(ctx context.Context, params query.GetMessagesByGroupId) (res query.GetMessagesByGroupIdResult, err error) {
	err = s.uow.Do(ctx, func(ctx context.Context, repos MessageRepos) error {
		before := time.Now()
		if params.Before != nil {
			before = *params.Before
		}
		limit := int32(100)
		if params.Limit != nil && *params.Limit <= 500 && *params.Limit > 0 {
			limit = *params.Limit
		}

		_, err := repos.Message().FindByGroupId(ctx, entities.DMGroupId(params.GroupId), before, limit)
		if err != nil {
			return entities.GetErrOrDefault(err, entities.ErrCodeDepFail, "cannot get messages")
		}

		return entities.NewError(entities.ErrCodeForbidden, "dm group not implemented", nil)
	})

	return res, err
}

func (s *MessageService) GetByChannelId(ctx context.Context, params query.GetMessagesByChannelId) (res query.GetMessagesByChannelIdResult, err error) {
	err = s.uow.Do(ctx, func(ctx context.Context, repos MessageRepos) error {
		_, _, _, derr := s.getChannelContext(ctx, repos, entities.ChannelId(params.ChannelId), entities.UserId(params.UserId))
		if derr != nil {
			return derr
		}
		// TODO: Check permission with roles, channel overwrite and stuff

		before := time.Now()
		if params.Before != nil {
			before = *params.Before
		}
		limit := int32(100)
		if params.Limit != nil && *params.Limit <= 500 && *params.Limit > 0 {
			limit = *params.Limit
		}

		msgs, err := repos.Message().FindByChannelId(ctx, entities.ChannelId(params.ChannelId), before, limit)
		if err != nil {
			return entities.GetErrOrDefault(err, entities.ErrCodeDepFail, "cannot get messages")
		}

		res = query.GetMessagesByChannelIdResult{
			Result: arrutil.Map(msgs, func(m *entities.Message) (target *common.Message, find bool) {
				return mapper.MessageToResult(m), true
			}),
		}
		return nil
	})

	return res, err
}

func (s *MessageService) Create(ctx context.Context, params command.CreateMessageCommand) (res command.CreateMessageCommandResult, err error) {
	msg, err := entities.NewMessage((*entities.ChannelId)(params.ChannelId), (*entities.DMGroupId)(params.GroupId), entities.UserId(params.UserId), params.Content, nil)
	if err != nil {
		return res, err
	}

	err = s.uow.Do(ctx, func(ctx context.Context, repos MessageRepos) error {
		if msg.ChannelId != nil {
			_, _, _, err = s.getChannelContext(ctx, repos, *msg.ChannelId, msg.Author)
			if err != nil {
				return err
			}
			// TODO: Check permission with roles, channel overwrite and stuff

			res = command.CreateMessageCommandResult{
				Result: mapper.MessageToResult(msg),
			}
			return nil
		} else {
			return entities.NewError(entities.ErrCodeForbidden, "dm group not implemented", nil)
		}
	})

	return res, err
}

func (s *MessageService) Update(context.Context, command.UpdateMessageCommand) (res command.UpdateMessageCommandResult, err error) {

	return res, err
}

func (s *MessageService) Delete(ctx context.Context, params command.DeleteMessageCommand) error {
	return s.uow.Do(ctx, func(ctx context.Context, repos MessageRepos) error {
		msg, err := repos.Message().Find(ctx, entities.MessageId(params.MessageId))
		if err != nil {
			return entities.GetErrOrDefault(err, entities.ErrCodeDepFail, "cannot get message")
		}

		if msg.ChannelId != nil {
			_, server, membership, derr := s.getChannelContext(ctx, repos, *msg.ChannelId, entities.UserId(params.UserId))
			if derr != nil {
				return derr
			}

			if !server.IsOwner(membership.UserId) && !msg.IsAuthor(membership.UserId) {
				return entities.NewError(entities.ErrCodeForbidden, "user don't have permission to delete message", nil)
			}
			// TODO: Add other permission check later

			err = msg.Delete()
			if err != nil {
				return err
			}

			msg, err = repos.Message().Save(ctx, msg)
			if err != nil {
				return entities.GetErrOrDefault(err, entities.ErrCodeDepFail, "cannot delete message")
			}
			return nil
		} else {
			return entities.NewError(entities.ErrCodeForbidden, "dm group not implemented", nil)
		}
	})
}
