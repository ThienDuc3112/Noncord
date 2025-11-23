package interfaces

import (
	"backend/internal/application/command"
	"backend/internal/application/query"
	"context"
)

type InviteService interface {
	CreateInvitation(context.Context, command.CreateInvitationCommand) (command.CreateInvitationCommandResult, error)
	UpdateInvitation(context.Context, command.UpdateInvitationCommand) (command.UpdateInvitationCommandResult, error)
	InvalidateInvitation(context.Context, command.InvalidateInvitationCommand) error
}

type InviteQueries interface {
	GetInvitationById(context.Context, query.GetInvitation) (query.GetInvitationResult, error)
	GetInvitationsByServerId(context.Context, query.GetInvitationsByServerId) (query.GetInvitationsByServerIdResult, error)
}
