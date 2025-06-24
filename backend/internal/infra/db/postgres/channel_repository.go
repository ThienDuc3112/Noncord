package postgres

import (
	e "backend/internal/domain/entities"
	"backend/internal/domain/repositories"
	"backend/internal/infra/db/postgres/gen"
	"context"
	"fmt"
)

type PGChannelRepo struct {
	db gen.DBTX
}

func (r *PGChannelRepo) Find(ctx context.Context, id e.ChannelId) (*e.Channel, error) {
	return nil, fmt.Errorf("Not implemented")
}

func (r *PGChannelRepo) FindIds(ctx context.Context, ids []e.ChannelId) ([]*e.Channel, error) {
	return nil, fmt.Errorf("Not implemented")
}

func (r *PGChannelRepo) FindByServerId(ctx context.Context, serverId e.ServerId) ([]*e.Channel, error) {
	return nil, fmt.Errorf("Not implemented")
}

func (r *PGChannelRepo) FindRoleOverrides(ctx context.Context, id e.ChannelId) ([]*e.ChannelRolePermissionOverride, error) {
	return nil, fmt.Errorf("Not implemented")
}

func (r *PGChannelRepo) FindRoleOverrideByRoleId(ctx context.Context, id e.ChannelId, roleId e.RoleId) (*e.ChannelRolePermissionOverride, error) {
	return nil, fmt.Errorf("Not implemented")
}

func (r *PGChannelRepo) FindUserOverrides(ctx context.Context, id e.ChannelId) (*e.ChannelUserPermissionOverride, error) {
	return nil, fmt.Errorf("Not implemented")
}

func (r *PGChannelRepo) FindUserOverrideByUserId(ctx context.Context, id e.ChannelId, userId e.UserId) (*e.ChannelUserPermissionOverride, error) {
	return nil, fmt.Errorf("Not implemented")
}

func (r *PGChannelRepo) Save(ctx context.Context, channel *e.Channel) (*e.Channel, error) {
	return nil, fmt.Errorf("Not implemented")
}

func (r *PGChannelRepo) SaveRoleOverride(ctx context.Context, perm *e.ChannelRolePermissionOverride) (*e.ChannelRolePermissionOverride, error) {
	return nil, fmt.Errorf("Not implemented")
}

func (r *PGChannelRepo) SaveUserOverride(ctx context.Context, perm *e.ChannelUserPermissionOverride) (*e.ChannelUserPermissionOverride, error) {
	return nil, fmt.Errorf("Not implemented")
}

func (r *PGChannelRepo) Delete(ctx context.Context, id e.ChannelId) error {
	return fmt.Errorf("Not implemented")
}

func (r *PGChannelRepo) DeleteRoleOverride(ctx context.Context, id e.ChannelId, roleId e.RoleId) error {
	return fmt.Errorf("Not implemented")
}

func (r *PGChannelRepo) DeleteUserOverride(ctx context.Context, id e.ChannelId, userId e.UserId) error {
	return fmt.Errorf("Not implemented")
}

var _ repositories.ChannelRepo = &PGChannelRepo{}
