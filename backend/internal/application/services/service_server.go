package services

import (
	"backend/internal/application/command"
	"backend/internal/application/common"
	"backend/internal/application/interfaces"
	"backend/internal/application/mapper"
	"backend/internal/application/query"
	"backend/internal/domain/entities"
	"backend/internal/domain/repositories"
	"context"

	"github.com/google/uuid"
	"github.com/gookit/goutil/arrutil"
)

type ServerService struct {
	sr repositories.ServerRepo
}

func NewServerService(sr repositories.ServerRepo) interfaces.ServerService {
	return &ServerService{
		sr: sr,
	}
}

func (s *ServerService) Create(ctx context.Context, params command.CreateServerCommand) (command.CreateServerCommandResult, error) {
	server := entities.NewServer(entities.UserId(params.UserId), params.Name, "", "", "", false)
	_, err := s.sr.Save(ctx, server)
	if err != nil {
		return command.CreateServerCommandResult{}, entities.GetErrOrDefault(err, entities.ErrCodeDepFail, "cannot save server failed")
	}

	// TODO: Create default channel and role

	return command.CreateServerCommandResult{
		Result: mapper.ServerToResult(server),
	}, nil
}

func (s *ServerService) Get(ctx context.Context, params query.GetServer) (query.GetServerResult, error) {
	server, err := s.sr.Find(ctx, entities.ServerId(params.ServerId))
	if err != nil {
		return query.GetServerResult{}, entities.GetErrOrDefault(err, entities.ErrCodeDepFail, "cannot get server")
	}

	// TODO: Check if user is a member of the server or not
	// params.UserId

	return query.GetServerResult{
		Result: mapper.ServerToResult(server),
	}, nil
}

func (s *ServerService) GetServers(ctx context.Context, params query.GetServers) (query.GetServersResult, error) {
	var mapFn arrutil.MapFn[uuid.UUID, entities.ServerId] = func(input uuid.UUID) (target entities.ServerId, find bool) {
		return entities.ServerId(input), true
	}
	servers, err := s.sr.FindByIds(ctx, arrutil.Map(params.ServerIds, mapFn))
	if err != nil {
		return query.GetServersResult{}, entities.GetErrOrDefault(err, entities.ErrCodeDepFail, "cannot get server")
	}

	return query.GetServersResult{
		Result: arrutil.Map(servers, func(server *entities.Server) (target *common.Server, find bool) {
			return mapper.ServerToResult(server), true
		}),
	}, nil
}

func (s *ServerService) Update(ctx context.Context, params command.UpdateServerCommand) (command.UpdateServerCommandResult, error) {
	server, err := s.sr.Find(ctx, entities.ServerId(params.ServerId))
	if err != nil {
		return command.UpdateServerCommandResult{}, entities.GetErrOrDefault(err, entities.ErrCodeDepFail, "cannot get server")
	}

	// TODO: Update with mod and role permission
	if !server.IsOwner(entities.UserId(params.UserId)) {
		return command.UpdateServerCommandResult{}, entities.NewError(entities.ErrCodeForbidden, "not authorized", err)
	}

	if params.Updates.Name != nil {
		if err = server.UpdateName(*params.Updates.Name); err != nil {
			return command.UpdateServerCommandResult{}, err
		}
	}
	if params.Updates.Description != nil {
		if err = server.UpdateDescription(*params.Updates.Description); err != nil {
			return command.UpdateServerCommandResult{}, err
		}
	}
	if params.Updates.IconUrl != nil {
		if err = server.UpdateIconUrl(*params.Updates.IconUrl); err != nil {
			return command.UpdateServerCommandResult{}, err
		}
	}
	if params.Updates.BannerUrl != nil {
		if err = server.UpdateBannerUrl(*params.Updates.BannerUrl); err != nil {
			return command.UpdateServerCommandResult{}, err
		}
	}
	if params.Updates.NeedApproval != nil {
		if err = server.UpdateNeedApproval(*params.Updates.NeedApproval); err != nil {
			return command.UpdateServerCommandResult{}, err
		}
	}
	if params.Updates.AnnouncementChannel.Valid {
		if err = server.UpdateAnnouncementChannel((*entities.ChannelId)(&params.Updates.AnnouncementChannel.UUID)); err != nil {
			return command.UpdateServerCommandResult{}, err
		}
	}

	server, err = s.sr.Save(ctx, server)
	if err != nil {
		return command.UpdateServerCommandResult{}, entities.GetErrOrDefault(err, entities.ErrCodeDepFail, "cannot update server")
	}

	return command.UpdateServerCommandResult{
		Result: mapper.ServerToResult(server),
	}, nil
}

func (s *ServerService) Delete(ctx context.Context, param command.DeleteServerCommand) error {
	server, err := s.sr.Find(ctx, entities.ServerId(param.ServerId))
	if err != nil {
		return entities.GetErrOrDefault(err, entities.ErrCodeDepFail, "cannot get server")
	}

	if server.Owner != entities.UserId(param.UserId) {
		return entities.NewError(entities.ErrCodeForbidden, "user is not the owner of the server", nil)
	}

	return entities.GetErrOrDefault(s.sr.Delete(ctx, server.Id), entities.ErrCodeDepFail, "cannot delete server")
}

// func (s *ServerService) CreateInvitation(ctx context.Context, param command.CreateInvitationCommand) (command.CreateInvitationCommandResult, error) {
// 	server, err := s.sr.Find(ctx, entities.ServerId(param.ServerId))
// 	if err != nil {
// 		return command.CreateInvitationCommandResult{}, entities.GetErrOrDefault(err, entities.ErrCodeDepFail, "cannot get server")
// 	}
//
// 	if !server.IsOwner(entities.UserId(param.UserId)) {
// 		return command.CreateInvitationCommandResult{}, entities.NewError(entities.ErrCodeForbidden, "not enough permission to create invitation", nil)
// 	}
// 	// TODO: Check permission in role as well
//
// 	invitation := entities.NewInvitation(entities.ServerId(param.ServerId), param.ExpiresAt, param.BypassApproval, param.JoinLimit)
// 	newInv, err := s.ir.Save(ctx, invitation)
// 	if err != nil {
// 		return command.CreateInvitationCommandResult{}, err
// 	}
//
// 	return command.CreateInvitationCommandResult{
// 		Result: mapper.InvitationToResult(newInv),
// 	}, nil
// }
//
// func (s *ServerService) UpdateInvitation(ctx context.Context, params command.UpdateInvitationCommand) (command.UpdateInvitationCommandResult, error) {
// 	server, err := s.sr.FindByInvitationId(ctx, entities.InvitationId(params.InvitationId))
// 	if err != nil {
// 		return command.UpdateInvitationCommandResult{}, entities.GetErrOrDefault(err, entities.ErrCodeDepFail, "cannot get server")
// 	}
//
// 	if !server.IsOwner(entities.UserId(params.UserId)) {
// 		return command.UpdateInvitationCommandResult{}, entities.NewError(entities.ErrCodeForbidden, "not enough permission to create invitation", nil)
// 	}
// 	// TODO: Check permission in role as well
//
// 	inv, err := s.ir.Find(ctx, entities.InvitationId(params.InvitationId))
// 	if err != nil {
// 		return command.UpdateInvitationCommandResult{}, entities.GetErrOrDefault(err, entities.ErrCodeDepFail, "cannot get invitation")
// 	}
//
// 	if params.Updates.BypassApproval != nil {
// 		if err = inv.UpdateBypassApproval(*params.Updates.BypassApproval); err != nil {
// 			return command.UpdateInvitationCommandResult{}, err
// 		}
// 	}
// 	if params.Updates.JoinLimit != nil {
// 		if err = inv.UpdateJoinLimit(*params.Updates.JoinLimit); err != nil {
// 			return command.UpdateInvitationCommandResult{}, err
// 		}
// 	}
//
// 	newInv, err := s.ir.Save(ctx, inv)
// 	if err != nil {
// 		return command.UpdateInvitationCommandResult{}, entities.GetErrOrDefault(err, entities.ErrCodeDepFail, "cannot update invitation")
// 	}
// 	return command.UpdateInvitationCommandResult{
// 		Result: mapper.InvitationToResult(newInv),
// 	}, nil
// }
