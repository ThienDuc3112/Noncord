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

	"github.com/gookit/goutil/arrutil"
)

type InvitationService struct {
	sr repositories.ServerRepo
	ir repositories.InvitationRepo
}

func NewInvitationService(sr repositories.ServerRepo, ir repositories.InvitationRepo) interfaces.InviteService {
	return &InvitationService{
		sr: sr,
		ir: ir,
	}
}

func (s *InvitationService) CreateInvitation(ctx context.Context, param command.CreateInvitationCommand) (command.CreateInvitationCommandResult, error) {
	server, err := s.sr.Find(ctx, entities.ServerId(param.ServerId))
	if err != nil {
		return command.CreateInvitationCommandResult{}, entities.GetErrOrDefault(err, entities.ErrCodeDepFail, "cannot get server")
	}

	if !server.IsOwner(entities.UserId(param.UserId)) {
		return command.CreateInvitationCommandResult{}, entities.NewError(entities.ErrCodeForbidden, "not enough permission to create invitation", nil)
	}
	// TODO: Check permission in role as well

	invitation := entities.NewInvitation(entities.ServerId(param.ServerId), param.ExpiresAt, param.BypassApproval, param.JoinLimit)
	newInv, err := s.ir.Save(ctx, invitation)
	if err != nil {
		return command.CreateInvitationCommandResult{}, err
	}

	return command.CreateInvitationCommandResult{
		Result: mapper.InvitationToResult(newInv),
	}, nil
}

func (s *InvitationService) UpdateInvitation(ctx context.Context, params command.UpdateInvitationCommand) (command.UpdateInvitationCommandResult, error) {
	server, err := s.sr.FindByInvitationId(ctx, entities.InvitationId(params.InvitationId))
	if err != nil {
		return command.UpdateInvitationCommandResult{}, entities.GetErrOrDefault(err, entities.ErrCodeDepFail, "cannot get server")
	}

	if !server.IsOwner(entities.UserId(params.UserId)) {
		return command.UpdateInvitationCommandResult{}, entities.NewError(entities.ErrCodeForbidden, "not enough permission to create invitation", nil)
	}
	// TODO: Check permission in role as well

	inv, err := s.ir.Find(ctx, entities.InvitationId(params.InvitationId))
	if err != nil {
		return command.UpdateInvitationCommandResult{}, entities.GetErrOrDefault(err, entities.ErrCodeDepFail, "cannot get invitation")
	}

	if params.Updates.BypassApproval != nil {
		if err = inv.UpdateBypassApproval(*params.Updates.BypassApproval); err != nil {
			return command.UpdateInvitationCommandResult{}, err
		}
	}
	if params.Updates.JoinLimit != nil {
		if err = inv.UpdateJoinLimit(*params.Updates.JoinLimit); err != nil {
			return command.UpdateInvitationCommandResult{}, err
		}
	}

	newInv, err := s.ir.Save(ctx, inv)
	if err != nil {
		return command.UpdateInvitationCommandResult{}, entities.GetErrOrDefault(err, entities.ErrCodeDepFail, "cannot update invitation")
	}
	return command.UpdateInvitationCommandResult{
		Result: mapper.InvitationToResult(newInv),
	}, nil
}

func (s *InvitationService) GetInvitationById(ctx context.Context, params query.GetInvitation) (query.GetInvitationResult, error) {
	inv, err := s.ir.Find(ctx, entities.InvitationId(params.InvitationId))
	if err != nil {
		return query.GetInvitationResult{}, entities.GetErrOrDefault(err, entities.ErrCodeDepFail, "cannot get invitation")
	}

	return query.GetInvitationResult{
		Result: mapper.InvitationToResult(inv),
	}, nil
}

func (s *InvitationService) GetInvitationsByServerId(ctx context.Context, params query.GetInvitationsByServerId) (query.GetInvitationsByServerIdResult, error) {
	server, err := s.sr.Find(ctx, entities.ServerId(params.ServerId))
	if err != nil {
		return query.GetInvitationsByServerIdResult{}, entities.GetErrOrDefault(err, entities.ErrCodeDepFail, "cannot get server")
	}

	// TODO: Check for other permission as well
	if server.IsOwner(entities.UserId(params.UserId)) {
		return query.GetInvitationsByServerIdResult{}, entities.NewError(entities.ErrCodeDepFail, "user don't have enough permission", nil)
	}

	invs, err := s.ir.FindByServerId(ctx, entities.ServerId(params.ServerId))
	if err != nil {
		return query.GetInvitationsByServerIdResult{}, entities.GetErrOrDefault(err, entities.ErrCodeDepFail, "cannot get invitations")
	}

	return query.GetInvitationsByServerIdResult{
		Result: arrutil.Map(invs, func(inv *entities.Invitation) (target *common.Invitation, find bool) {
			return mapper.InvitationToResult(inv), true
		}),
	}, nil
}

func (s *InvitationService) InvalidateInvitation(ctx context.Context, params command.InvalidateInvitationCommand) error {
	inv, err := s.ir.Find(ctx, entities.InvitationId(params.InvitationId))
	if err != nil {
		return entities.GetErrOrDefault(err, entities.ErrCodeDepFail, "cannot get invitation")
	}

	server, err := s.sr.Find(ctx, entities.ServerId(inv.ServerId))
	if err != nil {
		return entities.GetErrOrDefault(err, entities.ErrCodeDepFail, "cannot get server")
	}

	// Check for role permission as well
	if server.IsOwner(entities.UserId(params.UserId)) {
		return entities.NewError(entities.ErrCodeForbidden, "user not authorized to perfrom this action", nil)
	}

	inv.Invalidate()
	_, err = s.ir.Save(ctx, inv)
	return entities.GetErrOrDefault(err, entities.ErrCodeDepFail, "cannot save server")
}
