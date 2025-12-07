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
	Server() repositories.ServerRepo
	Message() repositories.MessageRepo
}

type MessageService struct {
	uow repositories.UnitOfWork[MessageRepos]
}

func NewMessageService(uow repositories.UnitOfWork[MessageRepos]) interfaces.MessageService {
	return &MessageService{uow}
}

func (s *MessageService) Create(ctx context.Context, params command.CreateMessageCommand) (res command.CreateMessageCommandResult, err error) {
	if params.IsTargetChannel {
		msg, err := entities.NewMessage((*entities.ChannelId)(&params.TargetId), nil, (*entities.UserId)(params.UserId), entities.AuthorType(params.AuthorType), params.Content, nil)
		if err != nil {
			return res, err
		}

		err = s.uow.Do(ctx, func(ctx context.Context, repos MessageRepos) error {
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
	// TODO: here
	return res, entities.NewError(entities.ErrCodeForbidden, "method not implemented", nil)
}

func (s *MessageService) Delete(ctx context.Context, params command.DeleteMessageCommand) error {
	return s.uow.Do(ctx, func(ctx context.Context, repos MessageRepos) error {
		msg, err := repos.Message().Find(ctx, entities.MessageId(params.MessageId))
		if err != nil {
			return entities.GetErrOrDefault(err, entities.ErrCodeDepFail, "cannot get message")
		}

		if msg.ChannelId != nil {
			if !msg.IsAuthor(entities.UserId(params.UserId)) && !params.HasMessageManagement {
				return entities.NewError(entities.ErrCodeForbidden, "user don't have permission to delete message", nil)
			}

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
