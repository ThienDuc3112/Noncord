package entities

import (
	"backend/internal/domain/events"
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
	events.Recorder

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
	if len(u.DisplayName) == 0 {
		return NewError(ErrCodeValidationError, "display name cannot be empty", nil)
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

type UpdateUserParam struct {
	Username    *string
	DisplayName *string
	AboutMe     *string
	Email       *string
	Password    *string
	AvatarUrl   *string
	BannerUrl   *string
	Flags       *UserFlags
	Disabled    *bool
	Verified    *bool
}

func (u *User) Update(p UpdateUserParam) error {
	changed := false

	// Username
	if p.Username != nil && *p.Username != u.Username {
		old := u.Username
		u.Username = *p.Username
		if err := u.Validate(); err != nil { // validate new state
			u.Username = old
			return err
		}
		u.Record(NewUserUsernameUpdated(u, old))
		changed = true
	}

	// DisplayName
	if p.DisplayName != nil && *p.DisplayName != u.DisplayName {
		old := u.DisplayName
		u.DisplayName = *p.DisplayName
		if err := u.Validate(); err != nil {
			u.DisplayName = old
			return err
		}
		u.Record(NewUserDisplayNameUpdated(u, old))
		changed = true
	}

	// AboutMe
	if p.AboutMe != nil && *p.AboutMe != u.AboutMe {
		old := u.AboutMe
		u.AboutMe = *p.AboutMe
		if err := u.Validate(); err != nil {
			u.AboutMe = old
			return err
		}
		u.Record(NewUserAboutMeUpdated(u, old))
		changed = true
	}

	// Email
	if p.Email != nil && *p.Email != u.Email {
		old := u.Email
		u.Email = *p.Email
		if err := u.Validate(); err != nil {
			u.Email = old
			return err
		}
		u.Record(NewUserEmailUpdated(u, old))
		changed = true
	}

	// Password (usually hashed upstream)
	if p.Password != nil && *p.Password != u.Password {
		old := u.Password
		u.Password = *p.Password
		// no Validate() needed here (policy-specific), but keep if you have rules
		u.Record(NewUserPasswordUpdated(u, old))
		changed = true
	}

	// AvatarUrl
	if p.AvatarUrl != nil && *p.AvatarUrl != u.AvatarUrl {
		old := u.AvatarUrl
		u.AvatarUrl = *p.AvatarUrl
		if err := u.Validate(); err != nil {
			u.AvatarUrl = old
			return err
		}
		u.Record(NewUserAvatarURLUpdated(u, old))
		changed = true
	}

	// BannerUrl
	if p.BannerUrl != nil && *p.BannerUrl != u.BannerUrl {
		old := u.BannerUrl
		u.BannerUrl = *p.BannerUrl
		if err := u.Validate(); err != nil {
			u.BannerUrl = old
			return err
		}
		u.Record(NewUserBannerURLUpdated(u, old))
		changed = true
	}

	// Flags
	if p.Flags != nil && *p.Flags != u.Flags {
		old := u.Flags
		u.Flags = *p.Flags
		u.Record(NewUserFlagsChanged(u, old))
		changed = true
	}

	// Disabled
	if p.Disabled != nil && *p.Disabled != u.Disabled {
		old := u.Disabled
		u.Disabled = *p.Disabled
		u.Record(NewUserDisabledChanged(u, old))
		changed = true
	}

	// Verified
	if p.Verified != nil && *p.Verified != u.Verified {
		old := u.Verified
		u.Verified = *p.Verified
		u.Record(NewUserVerifiedChanged(u, old))
		changed = true
	}

	if changed {
		u.UpdatedAt = time.Now()
	}
	return nil
}

// Soft delete. Idempotent.
func (u *User) Delete() {
	now := time.Now()
	var old *time.Time
	if u.DeletedAt != nil {
		// already deleted; keep old for event payload and refresh timestamp if you want
		old = u.DeletedAt
	}
	u.DeletedAt = &now
	u.Record(NewUserDeleted(u, old))
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
