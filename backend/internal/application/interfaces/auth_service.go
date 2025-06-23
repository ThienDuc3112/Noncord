package interfaces

import (
	"backend/internal/application/command"
	"context"
)

type AuthService interface {
	Register(context.Context, command.RegisterCommand) (command.RegisterCommandResult, error)
	Login(context.Context, command.LoginCommand) (command.LoginCommandResult, error)
	Logout(context.Context, command.LogoutCommand) error
	Refresh(context.Context, command.RefreshCommand) (command.RefreshCommandResult, error)
}
