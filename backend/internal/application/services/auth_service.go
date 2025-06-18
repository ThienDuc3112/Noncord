package services

import (
	"backend/internal/application/command"
	"backend/internal/application/interfaces"
	"backend/internal/domain/ports"
	"backend/internal/domain/repositories"
	"database/sql"
	"fmt"
)

type AuthService struct {
	userRepo repositories.UserRepo
	sessRepo ports.SessionRepository
	connPool *sql.DB
}

func NewAuthService(ur repositories.UserRepo, sr ports.SessionRepository, connPool *sql.DB) interfaces.AuthService {
	return &AuthService{
		userRepo: ur,
		sessRepo: sr,
		connPool: connPool,
	}
}

func (s *AuthService) Register(command.RegisterCommand) (command.RegisterCommandResult, error) {
	return command.RegisterCommandResult{}, fmt.Errorf("not implemented")
}

func (s *AuthService) Login(command.LoginCommand) (command.LoginCommandResult, error) {
	return command.LoginCommandResult{}, fmt.Errorf("not implemented")
}

func (s *AuthService) Logout(command.LogoutCommand) error {
	return fmt.Errorf("not implemented")
}

func (s *AuthService) Refresh(command.RefreshCommand) (command.RefreshCommandResult, error) {
	return command.RefreshCommandResult{}, fmt.Errorf("not implemented")

}
