package repositories

import "context"

type UnitOfWork[T any] interface {
	Do(ctx context.Context, fn func(ctx context.Context, repos T) error) error
}

type RepoBundle interface {
	Ban() BanRepo
	Channel() ChannelRepo
	DMGroup() DMGroupRepo
	Emote() EmoteRepo
	Invitation() InvitationRepo
	Member() MemberRepo
	Message() MessageRepo
	Role() RoleRepo
	Server() ServerRepo
	Session() SessionRepo
	UserNotification() UserNotificationRepo
	User() UserRepo
}

type BaseUnitOfWork interface {
	Do(ctx context.Context, fn func(ctx context.Context, repos RepoBundle) error) error
}
