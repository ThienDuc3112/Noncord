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

type PGChannelRepo struct {
	q *gen.Queries
}

func (r *PGChannelRepo) Find(ctx context.Context, id e.ChannelId) (*e.Channel, error) {
	channel, err := r.q.FindChannelById(ctx, uuid.UUID(id))
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, e.NewError(e.ErrCodeNoObject, "no channel by this id", err)
	} else if err != nil {
		return nil, err
	}

	return fromDbChannel(channel), nil
}

func (r *PGChannelRepo) FindIds(ctx context.Context, ids []e.ChannelId) ([]*e.Channel, error) {
	var mapper arrutil.MapFn[e.ChannelId, uuid.UUID] = func(input e.ChannelId) (target uuid.UUID, find bool) {
		return uuid.UUID(input), true
	}
	channels, err := r.q.FindChannelsByIds(ctx, arrutil.Map(ids, mapper))
	if err != nil {
		return nil, err
	}

	return arrutil.Map(channels, func(c gen.Channel) (target *e.Channel, find bool) {
		return fromDbChannel(c), true
	}), nil
}

func (r *PGChannelRepo) FindByServerId(ctx context.Context, serverId e.ServerId) ([]*e.Channel, error) {
	channels, err := r.q.FindChannelsByServerId(ctx, uuid.UUID(serverId))
	if err != nil {
		return nil, err
	}

	return arrutil.Map(channels, func(c gen.Channel) (target *e.Channel, find bool) {
		return fromDbChannel(c), true
	}), nil
}

func (r *PGChannelRepo) GetServerMaxChannelOrder(ctx context.Context, serverId e.ServerId) (int32, error) {
	return r.q.GetServerMaxOrdering(ctx, uuid.UUID(serverId))

}

// func (r *PGChannelRepo) FindRoleOverrides(ctx context.Context, id e.ChannelId) ([]*e.ChannelRolePermissionOverride, error) {
// 	return nil, fmt.Errorf("Not implemented")
// }

// func (r *PGChannelRepo) FindRoleOverrideByRoleId(ctx context.Context, id e.ChannelId, roleId e.RoleId) (*e.ChannelRolePermissionOverride, error) {
// 	return nil, fmt.Errorf("Not implemented")
// }

// func (r *PGChannelRepo) FindUserOverrides(ctx context.Context, id e.ChannelId) (*e.ChannelPermOverwrite, error) {
// 	return nil, fmt.Errorf("Not implemented")
// }

// func (r *PGChannelRepo) FindUserOverrideByUserId(ctx context.Context, id e.ChannelId, userId e.UserId) (*e.ChannelPermOverwrite, error) {
// 	return nil, fmt.Errorf("Not implemented")
// }

func (r *PGChannelRepo) Save(ctx context.Context, channel *e.Channel) (*e.Channel, error) {
	c, err := r.q.SaveChannel(ctx, gen.SaveChannelParams{
		ID:             uuid.UUID(channel.Id),
		CreatedAt:      channel.CreatedAt,
		UpdatedAt:      channel.UpdatedAt,
		DeletedAt:      channel.DeletedAt,
		Name:           channel.Name,
		Description:    channel.Description,
		ServerID:       uuid.UUID(channel.ServerId),
		Ordering:       int16(channel.Order),
		ParentCategory: (*uuid.UUID)(channel.ParentCategory),
	})
	if err != nil {
		return nil, err
	}

	if err = pullAndPushEvents(ctx, r.q, channel.PullsEvents()); err != nil {
		return nil, err
	}

	return fromDbChannel(c), nil
}

// func (r *PGChannelRepo) SaveRoleOverride(ctx context.Context, perm *e.ChannelRolePermissionOverride) (*e.ChannelRolePermissionOverride, error) {
// 	return nil, fmt.Errorf("Not implemented")
// }

// func (r *PGChannelRepo) SaveUserOverride(ctx context.Context, perm *e.ChannelPermOverwrite) (*e.ChannelPermOverwrite, error) {
// 	return nil, fmt.Errorf("Not implemented")
// }

func (r *PGChannelRepo) Delete(ctx context.Context, id e.ChannelId) error {
	return r.q.DeleteChannel(ctx, uuid.UUID(id))
}

// func (r *PGChannelRepo) DeleteRoleOverride(ctx context.Context, id e.ChannelId, roleId e.RoleId) error {
// 	return fmt.Errorf("Not implemented")
// }
//
// func (r *PGChannelRepo) DeleteUserOverride(ctx context.Context, id e.ChannelId, userId e.UserId) error {
// 	return fmt.Errorf("Not implemented")
// }

func (r *PGChannelRepo) FindByUserServers(ctx context.Context, userId e.UserId) ([]e.ChannelId, error) {
	ids, err := r.q.FindAllChannelInUserServers(ctx, uuid.UUID(userId))
	if err != nil {
		return nil, err
	}
	return arrutil.Map(ids, func(row gen.FindAllChannelInUserServersRow) (target e.ChannelId, find bool) {
		return e.ChannelId(row.ID), true
	}), nil
}

var _ repositories.ChannelRepo = &PGChannelRepo{}
