package entities

import (
	"time"

	"github.com/google/uuid"
)

type UserId uuid.UUID

type UserFlags uint8

const (
	UserFlagBot       UserFlags = 1
	UserFlagVerified  UserFlags = 2
	UserFlagModerator UserFlags = 3
	UserFlagStaff     UserFlags = 4
	UserFlagQA        UserFlags = 5
	UserFlagDev       UserFlags = 6
)

type User struct {
	Id          UserId
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   *time.Time
	Username    string
	DisplayName string
	AboutMe     string
	Email       string
	Password    string
	Disabled    bool
	AvatarUrl   string
	BannerUrl   string
	Flags       UserFlags
}

type Theme string

const (
	LightTheme Theme = "LIGHT"
	DarkTheme  Theme = "DARK"
)

type DMAllowOption uint8

const (
	DMAllowFriend DMAllowOption = 0
	DMAllowMember DMAllowOption = 1
	DMAllowAll    DMAllowOption = 2
)

type DMFilterOption uint8

const (
	DMFilterNone      DMFilterOption = 0
	DMFilterNonFriend DMFilterOption = 1
	DMFilterAll       DMFilterOption = 2
)

type FriendRequestPermissionBits uint8

const (
	FriendRequest2ndFriend FriendRequestPermissionBits = 1 << 1
	FriendRequestMember    FriendRequestPermissionBits = 1 << 2
	FriendRequestEveryone  FriendRequestPermissionBits = 1 << 3
)

type ReactionNotificationOption uint8

const (
	ReactionNotificationAll  = 0
	ReactionNotificationDM   = 1
	ReactionNotificationNone = 2
)

type UserSettings struct {
	UserId UserId

	// General
	Language string

	// Privacy settings
	DMAllowOption              DMAllowOption
	DMFilterOption             DMFilterOption
	FriendRequestPermission    FriendRequestPermissionBits
	CollectAnalyticsPermission bool

	// UI settings
	Theme Theme

	// Chat settings
	ShowEmote bool

	// Default Notifications
	ReactionNotification ReactionNotificationOption
	AFKTimeout           time.Duration
}

type Friendship struct {
	User1Id   UserId
	User2Id   UserId
	CreatedAt time.Time
}
