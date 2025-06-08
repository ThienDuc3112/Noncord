package entities

import "time"

type UserServerSettingsOverride struct {
	UserId               UserId
	ServerId             ServerId
	UpdatedAt            time.Time
	NotificationSettings NotificationBits
}

type UserChannelSettingsOverride struct {
	UserId               UserId
	ChannelId            ChannelId
	UpdatedAt            time.Time
	NotificationSettings NotificationBits
}

type UserDMSettingsOverride struct {
	UserId               UserId
	DMGroupId            DMGroupId
	UpdatedAt            time.Time
	NotificationSettings NotificationBits
}
