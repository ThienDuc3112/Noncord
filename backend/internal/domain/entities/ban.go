package entities

import "time"

type BanEntry struct {
	ServerId  ServerId
	UserId    UserId
	CreatedAt time.Time
}
