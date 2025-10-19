package postgres

import (
	e "backend/internal/domain/entities"
	"backend/internal/domain/repositories"
	"backend/internal/infra/db/postgres/gen"
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/gookit/goutil/arrutil"
	"github.com/jackc/pgx/v5"
)

type PGInvitationRepo struct {
	repo *gen.Queries
}

func NewPGInvitationRepo(db gen.DBTX) repositories.InvitationRepo {
	return &PGInvitationRepo{
		repo: gen.New(db),
	}
}

func (r *PGInvitationRepo) Find(ctx context.Context, id e.InvitationId) (*e.Invitation, error) {
	inv, err := r.repo.FindInvitationById(ctx, uuid.UUID(id))
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, e.NewError(e.ErrCodeNoObject, "invitation expired or invalid", err)
	} else if err != nil {
		return nil, err
	}

	return fromDbInvitation(inv), nil
}

func (r *PGInvitationRepo) FindByServerId(ctx context.Context, serverId e.ServerId) ([]*e.Invitation, error) {
	invs, err := r.repo.FindInvitationsByServerId(ctx, uuid.UUID(serverId))
	if err != nil {
		return nil, err
	}

	return arrutil.Map(invs, func(inv gen.Invitation) (target *e.Invitation, find bool) {
		return fromDbInvitation(inv), true
	}), nil
}

func (r *PGInvitationRepo) Save(ctx context.Context, invite *e.Invitation) (*e.Invitation, error) {
	i, err := r.repo.SaveInvitation(ctx, gen.SaveInvitationParams{
		ID:             uuid.UUID(invite.Id),
		ServerID:       uuid.UUID(invite.ServerId),
		CreatedAt:      invite.CreatedAt,
		ExpiredAt:      invite.ExpiresAt,
		BypassApproval: invite.BypassApproval,
		JoinLimit:      invite.JoinLimit,
		JoinCount:      invite.JoinCount,
	})
	if err != nil {
		return nil, err
	}

	if err = pullAndPushEvents(ctx, r.repo, invite.PullsEvents()); err != nil {
		return nil, err
	}

	return fromDbInvitation(i), nil
}

var _ repositories.InvitationRepo = &PGInvitationRepo{}
