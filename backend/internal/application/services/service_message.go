package services

import (
	"backend/internal/application/command"
	"backend/internal/application/interfaces"
	"backend/internal/application/mapper"
	"backend/internal/domain/entities"
	"backend/internal/domain/repositories"
	"context"
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

func (s *MessageService) getChannelContext(ctx context.Context, repos MessageRepos, channelId entities.ChannelId, userId entities.UserId) (*entities.Channel, *entities.Server, *entities.Membership, error) {
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

func (s *MessageService) Create(ctx context.Context, params command.CreateMessageCommand) (res command.CreateMessageCommandResult, err error) {
	if params.IsTargetChannel {
		msg, err := entities.NewMessage((*entities.ChannelId)(&params.TargetId), nil, (*entities.UserId)(params.UserId), entities.AuthorType(params.AuthorType), params.Content, nil)
		if err != nil {
			return res, err
		}

		err = s.uow.Do(ctx, func(ctx context.Context, repos MessageRepos) error {
			_, _, _, err = s.getChannelContext(ctx, repos, entities.ChannelId(params.TargetId), *msg.Author)
			if err != nil {
				return err
			}
			// TODO: Check permission with roles, channel overwrite and stuff

			msg, err = repos.Message().Save(ctx, msg)
			if err != nil {
				return entities.GetErrOrDefault(err, entities.ErrCodeDepFail, "failed to save message")
			}

			res = command.CreateMessageCommandResult{
				Result: mapper.MessageToResult(msg),
			}
			return nil
		})

		return res, err
	} else {
		return res, entities.NewError(entities.ErrCodeForbidden, "dm group not implemented", nil)
	}

}

func (s *MessageService) CreateSystemMessage(ctx context.Context, params command.CreateSystemMessageCommand) error {
	return s.uow.Do(ctx, func(ctx context.Context, repos MessageRepos) error {
		s, err := repos.Server().Find(ctx, entities.ServerId(params.ServerId))
		if err != nil {
			return entities.GetErrOrDefault(err, entities.ErrCodeDepFail, "cannot get server")
		}
		if s.AnnouncementChannel == nil {
			return nil
		}
		// TODO: Check permission with roles, channel overwrite and stuff
		msg, err := entities.NewMessage(s.AnnouncementChannel, nil, nil, entities.AuthorTypeSystem, params.Content, nil)
		if err != nil {
			return err
		}

		msg, err = repos.Message().Save(ctx, msg)
		if err != nil {
			return entities.GetErrOrDefault(err, entities.ErrCodeDepFail, "failed to save message")
		}

		return nil
	})
}

func (s *MessageService) Update(context.Context, command.UpdateMessageCommand) (res command.UpdateMessageCommandResult, err error) {
	return res, entities.NewError(entities.ErrCodeForbidden, "method not implemented", nil)
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
