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

type InvitationRepos interface {
	Invitation() repositories.InvitationRepo
	Server() repositories.ServerRepo
}

type InvitationService struct {
	uow repositories.UnitOfWork[InvitationRepos]
}

func NewInvitationService(uow repositories.UnitOfWork[InvitationRepos]) interfaces.InviteService {
	return &InvitationService{uow}
}

func (s *InvitationService) CreateInvitation(ctx context.Context, param command.CreateInvitationCommand) (res command.CreateInvitationCommandResult, err error) {
	err = s.uow.Do(ctx, func(ctx context.Context, repos InvitationRepos) error {
		server, err := repos.Server().Find(ctx, entities.ServerId(param.ServerId))
		if err != nil {
			return entities.GetErrOrDefault(err, entities.ErrCodeDepFail, "cannot get server")
		}

		if !server.IsOwner(entities.UserId(param.UserId)) {
			return entities.NewError(entities.ErrCodeForbidden, "not enough permission to create invitation", nil)
		}
		// TODO: Check permission in role as well

		invitation := entities.NewInvitation(entities.ServerId(param.ServerId), param.ExpiresAt, param.BypassApproval, param.JoinLimit)
		newInv, err := repos.Invitation().Save(ctx, invitation)
		if err != nil {
			return entities.GetErrOrDefault(err, entities.ErrCodeDepFail, "cannot create new invitation")
		}

		res = command.CreateInvitationCommandResult{
			Result: mapper.InvitationToResult(newInv),
		}
		return nil
	})

	return res, err
}

func (s *InvitationService) UpdateInvitation(ctx context.Context, params command.UpdateInvitationCommand) (res command.UpdateInvitationCommandResult, err error) {
	err = s.uow.Do(ctx, func(ctx context.Context, repos InvitationRepos) error {
		server, err := repos.Server().FindByInvitationId(ctx, entities.InvitationId(params.InvitationId))
		if err != nil {
			return entities.GetErrOrDefault(err, entities.ErrCodeDepFail, "cannot get server")
		}

		if !server.IsOwner(entities.UserId(params.UserId)) {
			return entities.NewError(entities.ErrCodeForbidden, "not enough permission to create invitation", nil)
		}
		// TODO: Check permission in role as well

		inv, err := repos.Invitation().Find(ctx, entities.InvitationId(params.InvitationId))
		if err != nil {
			return entities.GetErrOrDefault(err, entities.ErrCodeDepFail, "cannot get invitation")
		}

		if params.Updates.BypassApproval != nil {
			if err = inv.UpdateBypassApproval(*params.Updates.BypassApproval); err != nil {
				return err
			}
		}
		if params.Updates.JoinLimit != nil {
			if err = inv.UpdateJoinLimit(*params.Updates.JoinLimit); err != nil {
				return err
			}
		}

		newInv, err := repos.Invitation().Save(ctx, inv)
		if err != nil {
			return entities.GetErrOrDefault(err, entities.ErrCodeDepFail, "cannot update invitation")
		}

		res = command.UpdateInvitationCommandResult{
			Result: mapper.InvitationToResult(newInv),
		}
		return nil
	})

	return res, err
}

func (s *InvitationService) GetInvitationById(ctx context.Context, params query.GetInvitation) (res query.GetInvitationResult, err error) {
	err = s.uow.Do(ctx, func(ctx context.Context, repos InvitationRepos) error {
		inv, err := repos.Invitation().Find(ctx, entities.InvitationId(params.InvitationId))
		if err != nil {
			return entities.GetErrOrDefault(err, entities.ErrCodeDepFail, "cannot get invitation")
		}

		res = query.GetInvitationResult{
			Result: mapper.InvitationToResult(inv),
		}
		return nil
	})

	return res, err
}

func (s *InvitationService) GetInvitationsByServerId(ctx context.Context, params query.GetInvitationsByServerId) (res query.GetInvitationsByServerIdResult, err error) {
	err = s.uow.Do(ctx, func(ctx context.Context, repos InvitationRepos) error {
		server, err := repos.Server().Find(ctx, entities.ServerId(params.ServerId))
		if err != nil {
			return entities.GetErrOrDefault(err, entities.ErrCodeDepFail, "cannot get server")
		}

		// TODO: Check for other permission as well
		if !server.IsOwner(entities.UserId(params.UserId)) {
			return entities.NewError(entities.ErrCodeForbidden, "user don't have enough permission", nil)
		}

		invs, err := repos.Invitation().FindByServerId(ctx, entities.ServerId(params.ServerId))
		if err != nil {
			return entities.GetErrOrDefault(err, entities.ErrCodeDepFail, "cannot get invitations")
		}

		res = query.GetInvitationsByServerIdResult{
			Result: arrutil.Map(invs, func(inv *entities.Invitation) (target *common.Invitation, find bool) {
				return mapper.InvitationToResult(inv), true
			}),
		}
		return nil
	})

	return res, err
}

func (s *InvitationService) InvalidateInvitation(ctx context.Context, params command.InvalidateInvitationCommand) error {
	return s.uow.Do(ctx, func(ctx context.Context, repos InvitationRepos) error {
		inv, err := repos.Invitation().Find(ctx, entities.InvitationId(params.InvitationId))
		if err != nil {
			return entities.GetErrOrDefault(err, entities.ErrCodeDepFail, "cannot get invitation")
		}

		server, err := repos.Server().Find(ctx, entities.ServerId(inv.ServerId))
		if err != nil {
			return entities.GetErrOrDefault(err, entities.ErrCodeDepFail, "cannot get server")
		}

		// Check for role permission as well
		if !server.IsOwner(entities.UserId(params.UserId)) {
			return entities.NewError(entities.ErrCodeForbidden, "user not authorized to perfrom this action", nil)
		}

		inv.Invalidate()
		_, err = repos.Invitation().Save(ctx, inv)
		return entities.GetErrOrDefault(err, entities.ErrCodeDepFail, "cannot save server")
	})
}
