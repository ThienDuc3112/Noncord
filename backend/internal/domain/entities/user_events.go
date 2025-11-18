package entities

import (
	"backend/internal/domain/events"
	"time"

	"github.com/google/uuid"
)

const (
	EventUserCreated            = "user.created"
	EventUserUsernameUpdated    = "user.username_updated"
	EventUserDisplayNameUpdated = "user.display_name_updated"
	EventUserAboutMeUpdated     = "user.about_me_updated"
	EventUserEmailUpdated       = "user.email_updated"
	EventUserPasswordUpdated    = "user.password_updated"
	EventUserAvatarURLUpdated   = "user.avatar_url_updated"
	EventUserBannerURLUpdated   = "user.banner_url_updated"
	EventUserFlagsChanged       = "user.flags_changed"
	EventUserDisabledChanged    = "user.disabled_changed"
	EventUserVerifiedChanged    = "user.verified_changed"
	EventUserDeleted            = "user.deleted"

	UserCreatedSchemaVersion            = 1
	UserUsernameUpdatedSchemaVersion    = 1
	UserDisplayNameUpdatedSchemaVersion = 1
	UserAboutMeUpdatedSchemaVersion     = 1
	UserEmailUpdatedSchemaVersion       = 1
	UserPasswordUpdatedSchemaVersion    = 1
	UserAvatarURLUpdatedSchemaVersion   = 1
	UserBannerURLUpdatedSchemaVersion   = 1
	UserFlagsChangedSchemaVersion       = 1
	UserDisabledChangedSchemaVersion    = 1
	UserVerifiedChangedSchemaVersion    = 1
	UserDeletedSchemaVersion            = 1
)

// Optional: emit on NewUser() if you want creation in the stream.
type UserCreated struct {
	events.Base
	Username    string    `json:"username"`
	DisplayName string    `json:"display_name"`
	Email       string    `json:"email"`
	AvatarURL   string    `json:"avatar_url,omitempty"`
	BannerURL   string    `json:"banner_url,omitempty"`
	Flags       UserFlags `json:"flags"`
}

func NewUserCreated(u *User) UserCreated {
	return UserCreated{
		Base:        events.NewBase("user", uuid.UUID(u.Id), EventUserCreated, UserCreatedSchemaVersion),
		Username:    u.Username,
		DisplayName: u.DisplayName,
		Email:       u.Email,
		AvatarURL:   u.AvatarUrl,
		BannerURL:   u.BannerUrl,
		Flags:       u.Flags,
	}
}

type UserUsernameUpdated struct {
	events.Base
	Old string `json:"old"`
	New string `json:"new"`
}

func NewUserUsernameUpdated(u *User, old string) UserUsernameUpdated {
	return UserUsernameUpdated{
		Base: events.NewBase("user", uuid.UUID(u.Id), EventUserUsernameUpdated, UserUsernameUpdatedSchemaVersion),
		Old:  old,
		New:  u.Username,
	}
}

type UserDisplayNameUpdated struct {
	events.Base
	Old string `json:"old"`
	New string `json:"new"`
}

func NewUserDisplayNameUpdated(u *User, old string) UserDisplayNameUpdated {
	return UserDisplayNameUpdated{
		Base: events.NewBase("user", uuid.UUID(u.Id), EventUserDisplayNameUpdated, UserDisplayNameUpdatedSchemaVersion),
		Old:  old,
		New:  u.DisplayName,
	}
}

type UserAboutMeUpdated struct {
	events.Base
	Old string `json:"old"`
	New string `json:"new"`
}

func NewUserAboutMeUpdated(u *User, old string) UserAboutMeUpdated {
	return UserAboutMeUpdated{
		Base: events.NewBase("user", uuid.UUID(u.Id), EventUserAboutMeUpdated, UserAboutMeUpdatedSchemaVersion),
		Old:  old,
		New:  u.AboutMe,
	}
}

type UserEmailUpdated struct {
	events.Base
	Old string `json:"old"`
	New string `json:"new"`
}

func NewUserEmailUpdated(u *User, old string) UserEmailUpdated {
	return UserEmailUpdated{
		Base: events.NewBase("user", uuid.UUID(u.Id), EventUserEmailUpdated, UserEmailUpdatedSchemaVersion),
		Old:  old,
		New:  u.Email,
	}
}

type UserPasswordUpdated struct {
	events.Base
	// NEVER include previous password hashes in real systems.
	// We include "touched" only for audit trails without leaking secrets.
	Touched bool `json:"touched"`
}

func NewUserPasswordUpdated(u *User, _old string) UserPasswordUpdated {
	return UserPasswordUpdated{
		Base:    events.NewBase("user", uuid.UUID(u.Id), EventUserPasswordUpdated, UserPasswordUpdatedSchemaVersion),
		Touched: true,
	}
}

type UserAvatarURLUpdated struct {
	events.Base
	Old string `json:"old"`
	New string `json:"new"`
}

func NewUserAvatarURLUpdated(u *User, old string) UserAvatarURLUpdated {
	return UserAvatarURLUpdated{
		Base: events.NewBase("user", uuid.UUID(u.Id), EventUserAvatarURLUpdated, UserAvatarURLUpdatedSchemaVersion),
		Old:  old,
		New:  u.AvatarUrl,
	}
}

type UserBannerURLUpdated struct {
	events.Base
	Old string `json:"old"`
	New string `json:"new"`
}

func NewUserBannerURLUpdated(u *User, old string) UserBannerURLUpdated {
	return UserBannerURLUpdated{
		Base: events.NewBase("user", uuid.UUID(u.Id), EventUserBannerURLUpdated, UserBannerURLUpdatedSchemaVersion),
		Old:  old,
		New:  u.BannerUrl,
	}
}

type UserFlagsChanged struct {
	events.Base
	Old UserFlags `json:"old"`
	New UserFlags `json:"new"`
}

func NewUserFlagsChanged(u *User, old UserFlags) UserFlagsChanged {
	return UserFlagsChanged{
		Base: events.NewBase("user", uuid.UUID(u.Id), EventUserFlagsChanged, UserFlagsChangedSchemaVersion),
		Old:  old,
		New:  u.Flags,
	}
}

type UserDisabledChanged struct {
	events.Base
	Old bool `json:"old"`
	New bool `json:"new"`
}

func NewUserDisabledChanged(u *User, old bool) UserDisabledChanged {
	return UserDisabledChanged{
		Base: events.NewBase("user", uuid.UUID(u.Id), EventUserDisabledChanged, UserDisabledChangedSchemaVersion),
		Old:  old,
		New:  u.Disabled,
	}
}

type UserVerifiedChanged struct {
	events.Base
	Old bool `json:"old"`
	New bool `json:"new"`
}

func NewUserVerifiedChanged(u *User, old bool) UserVerifiedChanged {
	return UserVerifiedChanged{
		Base: events.NewBase("user", uuid.UUID(u.Id), EventUserVerifiedChanged, UserVerifiedChangedSchemaVersion),
		Old:  old,
		New:  u.Verified,
	}
}

type UserDeleted struct {
	events.Base
	OldDeletedAt *time.Time `json:"old_deleted_at,omitempty"`
	NewDeletedAt *time.Time `json:"new_deleted_at,omitempty"`
}

func NewUserDeleted(u *User, old *time.Time) UserDeleted {
	return UserDeleted{
		Base:         events.NewBase("user", uuid.UUID(u.Id), EventUserDeleted, UserDeletedSchemaVersion),
		OldDeletedAt: old,
		NewDeletedAt: u.DeletedAt,
	}
}

func init() {
	events.Register(EventUserCreated, UserCreatedSchemaVersion, func() events.DomainEvent { return UserCreated{} })
	events.Register(EventUserUsernameUpdated, UserUsernameUpdatedSchemaVersion, func() events.DomainEvent { return UserUsernameUpdated{} })
	events.Register(EventUserDisplayNameUpdated, UserDisplayNameUpdatedSchemaVersion, func() events.DomainEvent { return UserDisplayNameUpdated{} })
	events.Register(EventUserAboutMeUpdated, UserAboutMeUpdatedSchemaVersion, func() events.DomainEvent { return UserAboutMeUpdated{} })
	events.Register(EventUserEmailUpdated, UserEmailUpdatedSchemaVersion, func() events.DomainEvent { return UserEmailUpdated{} })
	events.Register(EventUserPasswordUpdated, UserPasswordUpdatedSchemaVersion, func() events.DomainEvent { return UserPasswordUpdated{} })
	events.Register(EventUserAvatarURLUpdated, UserAvatarURLUpdatedSchemaVersion, func() events.DomainEvent { return UserAvatarURLUpdated{} })
	events.Register(EventUserBannerURLUpdated, UserBannerURLUpdatedSchemaVersion, func() events.DomainEvent { return UserBannerURLUpdated{} })
	events.Register(EventUserFlagsChanged, UserFlagsChangedSchemaVersion, func() events.DomainEvent { return UserFlagsChanged{} })
	events.Register(EventUserDisabledChanged, UserDisabledChangedSchemaVersion, func() events.DomainEvent { return UserDisabledChanged{} })
	events.Register(EventUserVerifiedChanged, UserVerifiedChangedSchemaVersion, func() events.DomainEvent { return UserVerifiedChanged{} })
	events.Register(EventUserDeleted, UserDeletedSchemaVersion, func() events.DomainEvent { return UserDeleted{} })
}
