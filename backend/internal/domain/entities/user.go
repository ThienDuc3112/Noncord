package entities

import (
	"regexp"
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

var emailReg = regexp.MustCompile(`^[\w\-\.]+@([\w-]+\.)+[\w-]{2,}$`)
var usernameReg = regexp.MustCompile(`^[a-zA-Z0-9](?:[a-zA-Z0-9_-]{1,30}[a-zA-Z0-9])$`)
var urlReg = regexp.MustCompile(`(?:http[s]?:\/\/.)?(?:www\.)?[-a-zA-Z0-9@%._\+~#=]{2,256}\.[a-z]{2,6}\b(?:[-a-zA-Z0-9@:%_\+.~#?&\/\/=]*)`)

func (u *User) Validate() error {
	if len(u.Username) < 3 {
		return NewError(ErrCodeValidationError, "username must be longer than 3 characters", nil)
	}
	if len(u.Username) > 32 {
		return NewError(ErrCodeValidationError, "username cannot be longer than 32 characters", nil)
	}
	if len(u.DisplayName) > 128 {
		return NewError(ErrCodeValidationError, "display name cannot be longer than 128 characters", nil)
	}
	if !usernameReg.MatchString(u.Username) {
		return NewError(ErrCodeValidationError, "username must only contain alphanumeric character '_' or '-' and cannot start or end with '_' or '-' ", nil)
	}
	if !emailReg.MatchString(u.Email) {
		return NewError(ErrCodeValidationError, "invalid email", nil)
	}
	if u.AvatarUrl != "" && !urlReg.MatchString(u.AvatarUrl) {
		return NewError(ErrCodeValidationError, "avatar invalid url", nil)
	}
	if u.BannerUrl != "" && !urlReg.MatchString(u.BannerUrl) {
		return NewError(ErrCodeValidationError, "banner invalid url", nil)
	}

	return nil
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

type NotificationBits uint16

const (
	NotifyOnMentionEveryone   NotificationBits = 1 << 1
	NotifyOnMentionRole       NotificationBits = 1 << 2
	NotifyOnMentionDirect     NotificationBits = 1 << 3
	NotifyOnReply             NotificationBits = 1 << 4
	NotifyOnReactionServerMsg NotificationBits = 1 << 5
	NotifyOnReactionDM        NotificationBits = 1 << 6
	NotifyOnServerMessage     NotificationBits = 1 << 7
	NotifyOnGroupMessage      NotificationBits = 1 << 8
	NotifyOnDM                NotificationBits = 1 << 9
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
	NotificationSettings NotificationBits
	AFKTimeout           time.Duration
}

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

type Friendship struct {
	User1Id   UserId
	User2Id   UserId
	CreatedAt time.Time
}

type FriendRequest struct {
	RequesterId  UserId
	TargetUserId UserId
	CreatedAt    time.Time
}
