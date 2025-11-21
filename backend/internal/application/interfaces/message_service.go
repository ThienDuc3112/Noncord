package interfaces

import (
	"backend/internal/application/command"
	"backend/internal/application/query"
	"context"
)

type MessageService interface {
	Get(context.Context, query.GetMessage) (query.GetMessageResult, error)
	GetByGroupId(context.Context, query.GetMessagesByGroupId) (query.GetMessagesByGroupIdResult, error)
	GetByChannelId(context.Context, query.GetMessagesByChannelId) (query.GetMessagesByChannelIdResult, error)

	Create(context.Context, command.CreateMessageCommand) (command.CreateMessageCommandResult, error)
	CreateSystemMessage(context.Context, command.CreateSystemMessageCommand) error
	Update(context.Context, command.UpdateMessageCommand) (command.UpdateMessageCommandResult, error)
	Delete(context.Context, command.DeleteMessageCommand) error
}
