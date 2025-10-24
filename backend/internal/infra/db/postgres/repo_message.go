package postgres

import (
	e "backend/internal/domain/entities"
	"backend/internal/domain/repositories"
	"backend/internal/infra/db/postgres/gen"
	"context"
	"fmt"
	"time"
)

type PGMessageRepo struct {
	q *gen.Queries
}

func (r *PGMessageRepo) Find(ctx context.Context, id e.MessageId) (*e.Message, error) {
	return nil, fmt.Errorf("Not implemented")
}

func (r *PGMessageRepo) FindByChannelId(ctx context.Context, channelId e.ChannelId, before time.Time, limit int32) ([]*e.Message, error) {
	return nil, fmt.Errorf("Not implemented")
}

func (r *PGMessageRepo) FindByGroupId(ctx context.Context, groupId e.DMGroupId, before time.Time, limit int32) ([]*e.Message, error) {
	return nil, fmt.Errorf("Not implemented")
}

func (r *PGMessageRepo) Save(ctx context.Context, msg *e.Message) (*e.Message, error) {
	return nil, fmt.Errorf("Not implemented")
}

func (r *PGMessageRepo) Delete(ctx context.Context, id e.MessageId) error {
	return fmt.Errorf("Not implemented")
}

var _ repositories.MessageRepo = &PGMessageRepo{}
