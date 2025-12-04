package query

import "github.com/google/uuid"

type GetVisibleChannelsInServer struct {
	ServerId uuid.UUID
	UserId   uuid.UUID
}
