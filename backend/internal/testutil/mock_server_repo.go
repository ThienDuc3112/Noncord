package testutil

import (
	e "backend/internal/domain/entities"
	"backend/internal/domain/repositories"
	"context"

	"github.com/gookit/goutil/arrutil"
)

type MockServerRepo struct {
	db map[e.ServerId]*e.Server
}

func NewPGServerRepo(init []*e.Server) repositories.ServerRepo {
	mock := &MockServerRepo{
		db: make(map[e.ServerId]*e.Server),
	}
	for _, server := range init {
		mock.db[server.Id] = server
	}
	return mock
}

func (r *MockServerRepo) Save(ctx context.Context, server *e.Server) (*e.Server, error) {
	r.db[server.Id] = server
	return server, nil
}

// Find returns a server by ID
func (r *MockServerRepo) Find(ctx context.Context, id e.ServerId) (*e.Server, error) {
	if s, ok := r.db[id]; ok {
		return s, nil
	}
	return nil, e.NewError(e.ErrCodeNoObject, "server not found", nil)
}

// FindByIds returns multiple servers by IDs
func (r *MockServerRepo) FindByIds(ctx context.Context, ids []e.ServerId) ([]*e.Server, error) {
	return arrutil.Map(ids, func(id e.ServerId) (*e.Server, bool) {
		server, ok := r.db[id]
		return server, ok
	}), nil
}

func (r *MockServerRepo) FindByInvitationId(context.Context, e.InvitationId) (*e.Server, error)
func (r *MockServerRepo) FindByUser(context.Context, e.UserId) ([]*e.Server, error)

var _ repositories.ServerRepo = &MockServerRepo{}
