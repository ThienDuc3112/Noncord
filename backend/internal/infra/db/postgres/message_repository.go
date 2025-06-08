package postgres

import (
	e "backend/internal/domain/entities"
	"backend/internal/domain/repositories"
	"context"
	"database/sql"
	"time"
)

type PGMessageRepo struct {
	db *sql.DB
}

func (r *PGMessageRepo) Find(ctx context.Context, id e.MessageId) (*e.Message, error)
func (r *PGMessageRepo) FindByChannelId(ctx context.Context, channelId e.ChannelId, before time.Time, limit int32) ([]*e.Message, error)
func (r *PGMessageRepo) FindByGroupId(ctx context.Context, groupId e.DMGroupId, before time.Time, limit int32) ([]*e.Message, error)
func (r *PGMessageRepo) Save(ctx context.Context, msg *e.Message) (*e.Message, error)
func (r *PGMessageRepo) Delete(ctx context.Context, id e.MessageId) error

var _ repositories.MessageRepo = &PGMessageRepo{}
