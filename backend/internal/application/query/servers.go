package query

import (
	"backend/internal/application/common"

	"github.com/google/uuid"
)

type GetServer struct {
	ServerId uuid.UUID
	UserId   *uuid.UUID
}

type GetServerResult struct {
	Preview common.ServerPreview
	Full    *common.Server
	Channel []common.Channel
	Roles   []common.Role
}

type GetServers struct {
	ServerIds []uuid.UUID
}

type GetServersResult struct {
	Result []common.Server
}

type GetServersUserIn struct {
	UserId uuid.UUID
}

type GetServersUserInResult struct {
	Result []common.Server
}
