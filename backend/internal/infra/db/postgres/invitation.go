package postgres

import (
	e "backend/internal/domain/entities"
	"backend/internal/domain/repositories"
	"context"
	"database/sql"
)

type PGInviteRepo struct {
	db *sql.DB
}

func (r *PGInviteRepo) Find(ctx context.Context, id e.InvititationId) (*e.Invititation, error)
func (r *PGInviteRepo) FindByServerId(ctx context.Context, serverId e.ServerId) ([]*e.Invititation, error)
func (r *PGInviteRepo) Save(ctx context.Context, invite *e.Invititation) (*e.Invititation, error)

var _ repositories.InvitationRepo = &PGInviteRepo{}
