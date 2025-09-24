package interfaces

import (
	"backend/internal/application/command"
	"backend/internal/application/query"
	"context"
)

type ChannelService interface {
	Create(context.Context, command.CreateChannelCommand) (command.CreateChannelCommandResult, error)
	Get(context.Context, query.GetChannel) (query.GetChannelResult, error)
	GetChannelsByServer(context.Context, query.GetChannelsByServer) (query.GetChannelsByServerResult, error)
	Update(context.Context, command.UpdateChannelCommand) (command.UpdateChannelCommandResult, error)
	Delete(context.Context, command.DeleteChannelCommand) error
}
