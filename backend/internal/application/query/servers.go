package query

import (
	"backend/internal/application/common"
	"time"

	"github.com/google/uuid"
)

type GetServer struct {
	ServerId uuid.UUID
	UserId   *uuid.UUID
}

type Membership struct {
	ServerId  uuid.UUID
	UserId    uuid.UUID
	Nickname  string
	CreatedAt time.Time
}

type GetServerResult struct {
	Preview    common.ServerPreview
	Full       *common.Server
	Channel    []common.Channel
	Roles      []common.Role
	Membership Membership
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
