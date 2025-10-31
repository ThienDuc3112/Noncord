package rest

import (
	"backend/internal/application/common"
	"context"
	"log"
)

func extractUser(ctx context.Context) *common.UserResult {
	user, ok := ctx.Value("user").(*common.UserResult)
	if ok == false || user == nil {
		log.Println("[extractUser] context don't have user, this logcially shouldn't happen")
		return nil
	}
	return user
}
