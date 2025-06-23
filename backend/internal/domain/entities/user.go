package entities

import (
	"time"

	"github.com/google/uuid"
)

type UserId uuid.UUID

type UserFlags uint16

const (
	UserFlagUser      UserFlags = 0
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
	Verified    bool
	AvatarUrl   string
	BannerUrl   string
	Flags       UserFlags
}

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
	if len(u.AboutMe) > 1024 {
		return NewError(ErrCodeValidationError, "about me cannot be longer than 512 characters", nil)
	}
	if !IsValidUsername(u.Username) {
		return NewError(ErrCodeValidationError, "username must only contain alphanumeric character '_' or '-' and cannot start or end with '_' or '-' ", nil)
	}
	if !IsValidEmail(u.Email) {
		return NewError(ErrCodeValidationError, "invalid email", nil)
	}
	if len(u.Email) > 256 {
		return NewError(ErrCodeValidationError, "email too long, likely not valid, please use another one", nil)
	}
	if u.AvatarUrl != "" && !IsValidUrl(u.AvatarUrl) {
		return NewError(ErrCodeValidationError, "avatar invalid url", nil)
	}
	if len(u.AvatarUrl) > 2048 {
		return NewError(ErrCodeValidationError, "avatar url too long", nil)
	}
	if u.BannerUrl != "" && !IsValidUrl(u.BannerUrl) {
		return NewError(ErrCodeValidationError, "banner invalid url", nil)
	}
	if len(u.BannerUrl) > 2048 {
		return NewError(ErrCodeValidationError, "banner url too long", nil)
	}

	return nil
}

type NewUserParam struct {
	Username    string
	DisplayName string
	AboutMe     string
	Email       string
	Password    string
	AvatarUrl   string
	BannerUrl   string
	Flags       UserFlags
}

// func NewUser(username, displayName, aboutMe, email, password, avatarUrl, bannerUrl string, flags UserFlags) *User {
func NewUser(param NewUserParam) *User {
	return &User{
		Id:          UserId(uuid.New()),
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
		DeletedAt:   nil,
		Username:    param.Username,
		DisplayName: param.DisplayName,
		AboutMe:     param.AboutMe,
		Email:       param.Email,
		Password:    param.Password,
		Disabled:    false,
		Verified:    false,
		AvatarUrl:   param.AvatarUrl,
		BannerUrl:   param.BannerUrl,
		Flags:       param.Flags,
	}
}

type Theme string

const (
	LightTheme Theme = "LIGHT"
	DarkTheme  Theme = "DARK"
)

type DMAllowOption uint16

const (
	DMAllowFriend DMAllowOption = 0
	DMAllowMember DMAllowOption = 1
	DMAllowAll    DMAllowOption = 2
)

type DMFilterOption uint16

const (
	DMFilterNone      DMFilterOption = 0
	DMFilterNonFriend DMFilterOption = 1
	DMFilterAll       DMFilterOption = 2
)

type FriendRequestPermissionBits uint16

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

type ReactionNotificationOption uint16

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

func NewUserSettings(uid UserId, lang string, dmOption DMAllowOption, dmFilter DMFilterOption, friendReqPerm FriendRequestPermissionBits, colAnaPerm bool, theme Theme, showEmote bool, notiSetting NotificationBits, afkDur time.Duration) *UserSettings {
	return &UserSettings{
		UserId:                     uid,
		Language:                   lang,
		DMAllowOption:              dmOption,
		DMFilterOption:             dmFilter,
		FriendRequestPermission:    friendReqPerm,
		CollectAnalyticsPermission: colAnaPerm,
		Theme:                      theme,
		ShowEmote:                  showEmote,
		NotificationSettings:       notiSetting,
		AFKTimeout:                 afkDur,
	}
}

type FriendRequest struct {
	RequesterId  UserId
	TargetUserId UserId
	Message      string
	CreatedAt    time.Time
}

func NewFriendRequest(requester, target UserId, msg string) *FriendRequest {
	return &FriendRequest{
		RequesterId:  requester,
		TargetUserId: target,
		Message:      msg,
		CreatedAt:    time.Now(),
	}
}
