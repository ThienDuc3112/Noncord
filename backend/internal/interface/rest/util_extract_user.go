package rest

import (
	"context"
	"log/slog"

	"github.com/google/uuid"
)

func extractUserId(ctx context.Context) *uuid.UUID {
	userId, ok := ctx.Value(userIdKey).(*uuid.UUID)
	if ok == false || userId == nil {
		slog.Info("context don't have user id, this logcially shouldn't happen")
		return nil
	}
	return userId
}
