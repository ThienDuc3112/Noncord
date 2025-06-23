package services

import (
	"backend/internal/application/command"
	"backend/internal/application/interfaces"
	"backend/internal/application/mapper"
	"backend/internal/domain/entities"
	"backend/internal/domain/ports"
	"backend/internal/domain/repositories"
	"context"
	"database/sql"
	"errors"
	"fmt"

	"golang.org/x/crypto/bcrypt"
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

func (s *AuthService) Register(ctx context.Context, cmd command.RegisterCommand) (command.RegisterCommandResult, error) {
	if len(cmd.Password) < 8 || len(cmd.Password) > 72 {
		return command.RegisterCommandResult{}, entities.NewError(entities.ErrCodeValidationError, "password must be between 8 and 72 characters long", nil)
	}

	password, err := bcrypt.GenerateFromPassword([]byte(cmd.Password), bcrypt.DefaultCost)
	if err != nil {
		return command.RegisterCommandResult{}, entities.NewError(entities.ErrCodeDepFail, "cannot create new password", err)
	}

	user := entities.NewUser(entities.NewUserParam{
		Username:    cmd.Username,
		Email:       cmd.Email,
		DisplayName: "",
		AboutMe:     "",
		Password:    string(password),
		AvatarUrl:   "",
		BannerUrl:   "",
		Flags:       entities.UserFlagUser,
	})
	err = user.Validate()
	if err != nil {
		return command.RegisterCommandResult{}, err
	}

	err = s.userRepo.Save(ctx, user)
	var domainErr *entities.ChatError
	if errors.As(err, &domainErr) {
		return command.RegisterCommandResult{}, domainErr
	} else if err != nil {
		return command.RegisterCommandResult{}, entities.NewError(entities.ErrCodeDepFail, "cannot save user", err)
	}

	return command.RegisterCommandResult{
		Result: mapper.NewUserResultFromUserEntity(user),
	}, nil
}

func (s *AuthService) Login(ctx context.Context, param command.LoginCommand) (command.LoginCommandResult, error) {
	var user *entities.User
	var err error
	if entities.IsValidEmail(param.Username) {
		user, err = s.userRepo.FindByEmail(ctx, param.Username)
	} else {
		user, err = s.userRepo.FindByUsername(ctx, param.Username)
	}
	var domainErr *entities.ChatError
	if errors.As(err, &domainErr) {
		return command.LoginCommandResult{}, domainErr
	} else if err != nil {
		return command.LoginCommandResult{}, err
	}

	if user == nil {
		return command.LoginCommandResult{}, entities.NewError(entities.ErrCodeInvalidAssertion, "nil user despite nil error", nil)
	}

	if user.Password == "" {
		return command.LoginCommandResult{}, entities.NewError(entities.ErrCodeInvalidAction, "sso user cannot sign in with password", nil)
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(param.Password))
	if err != nil {
		return command.LoginCommandResult{}, entities.NewError(entities.ErrCodeLogicFailure, "invalid password", err)
	}

	// TODO: generate tokens here

	return command.LoginCommandResult{}, fmt.Errorf("not implemented")
}

func (s *AuthService) Logout(ctx context.Context, param command.LogoutCommand) error {
	return fmt.Errorf("not implemented")
}

func (s *AuthService) Refresh(ctx context.Context, param command.RefreshCommand) (command.RefreshCommandResult, error) {
	return command.RefreshCommandResult{}, fmt.Errorf("not implemented")

}
