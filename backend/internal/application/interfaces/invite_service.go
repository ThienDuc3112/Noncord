package interfaces

import (
	"backend/internal/application/command"
	"backend/internal/application/query"
	"context"
)

type InviteService interface {
	GetInvitationById(context.Context, query.GetInvitation) (query.GetInvitationResult, error)
	GetInvitationsByServerId(context.Context, query.GetInvitationsByServerId) (query.GetInvitationsByServerIdResult, error)
	CreateInvitation(context.Context, command.CreateInvitationCommand) (command.CreateInvitationCommandResult, error)
	UpdateInvitation(context.Context, command.UpdateInvitationCommand) (command.UpdateInvitationCommandResult, error)
}
