package services

import (
	"backend/internal/application/command"
	"backend/internal/application/interfaces"
	"backend/internal/application/mapper"
	"backend/internal/domain/entities"
	"backend/internal/domain/repositories"
	"context"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type AuthRepos interface {
	User() repositories.UserRepo
	Session() repositories.SessionRepo
}

type AuthService struct {
	uow    repositories.UnitOfWork[AuthRepos]
	secret string
}

func NewAuthService(uow repositories.UnitOfWork[AuthRepos], secret string) interfaces.AuthService {
	return &AuthService{
		uow:    uow,
		secret: secret,
	}
}

func (s *AuthService) Register(ctx context.Context, cmd command.RegisterCommand) (res command.RegisterCommandResult, err error) {
	if len(cmd.Password) < 8 || len(cmd.Password) > 72 {
		return res, entities.NewError(entities.ErrCodeValidationError, "password must be between 8 and 72 characters long", nil)
	}

	password, err := bcrypt.GenerateFromPassword([]byte(cmd.Password), bcrypt.DefaultCost)
	if err != nil {
		return res, entities.NewError(entities.ErrCodeDepFail, "cannot create new password", err)
	}

	user := entities.NewUser(entities.NewUserParam{
		Username:    strings.ToLower(cmd.Username),
		Email:       strings.ToLower(cmd.Email),
		DisplayName: cmd.Username,
		AboutMe:     "",
		Password:    string(password),
		AvatarUrl:   "",
		BannerUrl:   "",
		Flags:       entities.UserFlagUser,
	})
	err = user.Validate()
	if err != nil {
		return res, entities.GetErrOrDefault(err, entities.ErrCodeValidationError, "validation failed")
	}

	err = s.uow.Do(ctx, func(ctx context.Context, repos AuthRepos) error {
		return repos.User().Save(ctx, user)
	})
	if err != nil {
		return command.RegisterCommandResult{}, entities.GetErrOrDefault(err, entities.ErrCodeDepFail, "cannot save user")
	}

	return command.RegisterCommandResult{
		Result: mapper.NewUserResultFromUserEntity(user),
	}, nil
}

type AccessTokenClaim struct {
	UserId    string             `json:"userId"`
	Username  string             `json:"username"`
	UserFlags entities.UserFlags `json:"userFlags"`

	jwt.RegisteredClaims
}

func (s *AuthService) Login(ctx context.Context, param command.LoginCommand) (res command.LoginCommandResult, err error) {
	err = s.uow.Do(ctx, func(ctx context.Context, repos AuthRepos) error {
		var user *entities.User
		if entities.IsValidEmail(param.Username) {
			user, err = repos.User().FindByEmail(ctx, param.Username)
		} else {
			user, err = repos.User().FindByUsername(ctx, param.Username)
		}
		if err != nil {
			return entities.GetErrOrDefault(err, entities.ErrCodeDepFail, "cannot get user")
		}

		if user == nil {
			return entities.NewError(entities.ErrCodeInvalidAssertion, "nil user despite nil error", nil)
		}

		if user.Password == "" {
			return entities.NewError(entities.ErrCodeForbidden, "sso user cannot sign in with password", nil)
		}

		err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(param.Password))
		if err != nil {
			return entities.NewError(entities.ErrCodeUnauth, "invalid password", err)
		}

		now := time.Now()
		session := entities.NewSession(user.Id, now.Add(time.Hour*24*30), param.UserAgent)
		accessTokenClaim := jwt.NewWithClaims(jwt.SigningMethodHS256, AccessTokenClaim{
			UserId:    uuid.UUID(user.Id).String(),
			Username:  user.Username,
			UserFlags: user.Flags,
			RegisteredClaims: jwt.RegisteredClaims{
				Issuer:    "Noncord",
				ExpiresAt: jwt.NewNumericDate(now.Add(time.Minute * 30)),
				IssuedAt:  jwt.NewNumericDate(now),
			},
		})
		accessToken, err := accessTokenClaim.SignedString([]byte(s.secret))
		if err != nil {
			return entities.NewError(entities.ErrCodeDepFail, "fail to generate access token", err)
		}
		err = repos.Session().Save(ctx, session)
		if err != nil {
			return entities.GetErrOrDefault(err, entities.ErrCodeDepFail, "cannot save session")
		}

		res = command.LoginCommandResult{
			AccessToken:  accessToken,
			RefreshToken: session.Token,
		}
		return nil
	})
	return res, err
}

func (s *AuthService) Logout(ctx context.Context, param command.LogoutCommand) error {
	return s.uow.Do(ctx, func(ctx context.Context, repos AuthRepos) error {
		session, err := repos.Session().FindByToken(ctx, param.RefreshToken)
		if err != nil {
			return entities.GetErrOrDefault(err, entities.ErrCodeDepFail, "cannot get session")
		}

		session.ExpiresAt = time.Now()

		return entities.GetErrOrDefault(repos.Session().Save(ctx, session), entities.ErrCodeDepFail, "cannot set session")

	})
}

func (s *AuthService) Refresh(ctx context.Context, param command.RefreshCommand) (res command.RefreshCommandResult, err error) {
	err = s.uow.Do(ctx, func(ctx context.Context, repos AuthRepos) error {
		session, err := repos.Session().FindByToken(ctx, param.RefreshToken)
		if err != nil {
			return entities.GetErrOrDefault(err, entities.ErrCodeDepFail, "cannot get session")
		}

		user, err := repos.User().Find(ctx, session.UserId)
		if err != nil {
			return entities.GetErrOrDefault(err, entities.ErrCodeDepFail, "cannot get user")
		}

		now := time.Now()
		session.Token = entities.RandomToken()
		session.RotationCount += 1
		session.ExpiresAt = now.Add(30 * 24 * time.Hour)

		accessTokenClaim := jwt.NewWithClaims(jwt.SigningMethodHS256, AccessTokenClaim{
			UserId:    uuid.UUID(user.Id).String(),
			Username:  user.Username,
			UserFlags: user.Flags,
			RegisteredClaims: jwt.RegisteredClaims{
				Issuer:    "Noncord",
				ExpiresAt: jwt.NewNumericDate(now.Add(time.Minute * 30)),
				IssuedAt:  jwt.NewNumericDate(now),
			},
		})
		accessToken, err := accessTokenClaim.SignedString([]byte(s.secret))
		if err != nil {
			return entities.NewError(entities.ErrCodeDepFail, "fail to generate access token", err)
		}
		err = repos.Session().Save(ctx, session)
		if err != nil {
			return entities.GetErrOrDefault(err, entities.ErrCodeDepFail, "cannot save session")
		}

		res = command.RefreshCommandResult{
			AccessToken:  accessToken,
			RefreshToken: session.Token,
		}
		return nil
	})

	return res, err
}

func (s *AuthService) Authenticate(ctx context.Context, param command.AuthenticateCommand) (res command.AuthenticateCommandResult, err error) {
	token, err := jwt.ParseWithClaims(param.AccessToken, &AccessTokenClaim{}, func(token *jwt.Token) (any, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, entities.NewError(entities.ErrCodeUnauth, "invalid token", nil)
		}
		return []byte(s.secret), nil
	}, jwt.WithValidMethods([]string{"HS256"}))

	if err != nil {
		return res, entities.GetErrOrDefault(err, entities.ErrCodeUnauth, "cannot verify token")
	}

	claims, ok := token.Claims.(*AccessTokenClaim)
	if !ok {
		return res, entities.GetErrOrDefault(err, entities.ErrCodeUnauth, "deformed token structure")
	} else if !token.Valid {
		return res, entities.GetErrOrDefault(err, entities.ErrCodeUnauth, "invalid token, potentially expired")
	}

	userId, err := uuid.Parse(claims.UserId)
	if err != nil {
		return res, entities.GetErrOrDefault(err, entities.ErrCodeUnauth, "invalid token, invalid user id")
	}

	err = s.uow.Do(ctx, func(ctx context.Context, repos AuthRepos) error {
		user, err := repos.User().Find(ctx, entities.UserId(userId))
		if err != nil {
			return entities.GetErrOrDefault(err, entities.ErrCodeDepFail, "cannot find user")
		}

		res = command.AuthenticateCommandResult{
			User: mapper.NewUserResultFromUserEntity(user),
		}
		return nil
	})
	return res, err
}
