package interfaces

import (
	"backend/internal/application/command"
	"context"
)

type MembershipService interface {
	JoinServer(context.Context, command.AuthenticateCommand)
	// - [ ] Join server
	// - [ ] Leave server
	// - [ ] Kick member
	// - [ ] Ban member
	// - [ ] Assign role
	// - [ ] Update channel user permission
	// - [ ] Get server by user in
	// - [ ] Set nickname
}
