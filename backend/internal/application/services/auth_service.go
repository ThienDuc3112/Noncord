package services

import (
	"backend/internal/application/command"
	"backend/internal/application/interfaces"
	"backend/internal/domain/repositories"
	"database/sql"
	"fmt"
)

type AuthService struct {
	userRepo repositories.UserRepo
	connPool *sql.DB
}

func NewAuthService(repo repositories.UserRepo, connPool *sql.DB) interfaces.AuthService {
	return &AuthService{
		userRepo: repo,
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
