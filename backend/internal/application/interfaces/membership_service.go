package interfaces

import (
	"backend/internal/application/command"
	"context"
)

type MembershipService interface {
	JoinServer(context.Context, command.JoinServerCommand) (command.JoinServerCommandResult, error)
	LeaveServer(context.Context, command.LeaveServerCommand) error
	Kick(context.Context, command.KickCommand) error
	Ban(context.Context, command.BanCommand) error
	SetNickname(context.Context, command.SetNickname) error
}

type MembershipQueries interface {
}

// - [ ] Assign role
// - [ ] Update channel user permission
// - [ ] Get server by user in
