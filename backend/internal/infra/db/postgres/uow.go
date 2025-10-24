package postgres

import (
	"backend/internal/domain/entities"
	"backend/internal/domain/repositories"
	"backend/internal/infra/db/postgres/gen"
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type pgRepoBundle struct {
	q *gen.Queries
}

func newRepoBundle(q *gen.Queries) repositories.RepoBundle {
	return &pgRepoBundle{q}
}

func (b *pgRepoBundle) Ban() repositories.BanRepo {
	return &PGBanRepo{b.q}
}
func (b *pgRepoBundle) Channel() repositories.ChannelRepo {
	return &PGChannelRepo{b.q}
}
func (b *pgRepoBundle) DMGroup() repositories.DMGroupRepo {
	return &PGDMGroupRepo{b.q}
}
func (b *pgRepoBundle) Emote() repositories.EmoteRepo {
	return &PGEmoteRepo{b.q}
}
func (b *pgRepoBundle) Invitation() repositories.InvitationRepo {
	return &PGInvitationRepo{b.q}
}
func (b *pgRepoBundle) Member() repositories.MemberRepo {
	return &PGMemberRepo{b.q}
}
func (b *pgRepoBundle) Message() repositories.MessageRepo {
	return &PGMessageRepo{b.q}
}
func (b *pgRepoBundle) Role() repositories.RoleRepo {
	return &PGRoleRepo{b.q}
}
func (b *pgRepoBundle) Server() repositories.ServerRepo {
	return &PGServerRepo{b.q}
}
func (b *pgRepoBundle) Session() repositories.SessionRepo {
	return &PGSessionRepo{b.q}
}
func (b *pgRepoBundle) UserNotification() repositories.UserNotificationRepo {
	return &PGUserNotiRepo{b.q}
}
func (b *pgRepoBundle) User() repositories.UserRepo {
	return &PGUserRepo{b.q}
}

type baseUoW struct {
	pool *pgxpool.Pool
}

func NewBaseUoW(pool *pgxpool.Pool) repositories.BaseUnitOfWork { return &baseUoW{pool} }

func (u *baseUoW) Do(ctx context.Context, fn func(ctx context.Context, repos repositories.RepoBundle) error) error {
	tx, err := u.pool.Begin(ctx)
	if err != nil {
		return entities.NewError(entities.ErrCodeDepFail, "Cannot start a transaction", err)
	}
	defer func() { _ = tx.Rollback(ctx) }()

	qtx := gen.New(tx)
	full := newRepoBundle(qtx)

	if err = fn(ctx, full); err != nil {
		return err
	}
	return entities.GetErrOrDefault(tx.Commit(ctx), entities.ErrCodeDepFail, "Cannot commit changes")
}

type scopedUoW[T any] struct {
	base      repositories.BaseUnitOfWork
	projector func(repositories.RepoBundle) T
}

func NewScopedUoW[T any](base repositories.BaseUnitOfWork, projector func(repositories.RepoBundle) T) repositories.UnitOfWork[T] {
	return &scopedUoW[T]{base: base, projector: projector}
}

func (u *scopedUoW[T]) Do(ctx context.Context, fn func(ctx context.Context, repos T) error) error {
	return u.base.Do(ctx, func(ctx context.Context, full repositories.RepoBundle) error {
		view := u.projector(full) // narrow the bundle
		return fn(ctx, view)
	})
}
