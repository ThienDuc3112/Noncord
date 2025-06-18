package interfaces

import "backend/internal/application/command"

type AuthService interface {
	Register(command.RegisterCommand) (command.RegisterCommandResult, error)
	Login(command.LoginCommand) (command.LoginCommandResult, error)
	Logout(command.LogoutCommand) error
	Refresh(command.RefreshCommand) (command.RefreshCommandResult, error)
}
