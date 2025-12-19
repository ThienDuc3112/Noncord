package ws

import (
	"backend/internal/application/command"
	"backend/internal/application/interfaces"
	"context"
	"log/slog"

	"github.com/google/uuid"
)

func authMiddleware(authService interfaces.AuthService, token string) *uuid.UUID {
	slog.Info("Verifying user")

	res, err := authService.Authenticate(context.Background(), command.AuthenticateCommand{AccessToken: token})
	if err != nil {
		return nil
	}

	return res.UserId
}
