package interfaces

import (
	"backend/internal/application/command"
	"context"
)

type MembershipService interface {
	JoinServer(context.Context, command.JoinServerCommand) error
	LeaveServer(context.Context, command.LeaveServerCommand) error
	KickServer(context.Context, command.KickServerCommand) error
	BanServer(context.Context, command.BanServerCommand) error
	// - [ ] Assign role
	// - [ ] Update channel user permission
	// - [ ] Get server by user in
	// - [ ] Set nickname
}
