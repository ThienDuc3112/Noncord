package entities

import "time"

type BanEntry struct {
	ServerId  ServerId
	UserId    UserId
	CreatedAt time.Time
}

func NewBanEntry(serverId ServerId, userId UserId) *BanEntry {
	return &BanEntry{
		ServerId:  serverId,
		UserId:    userId,
		CreatedAt: time.Now(),
	}
}
