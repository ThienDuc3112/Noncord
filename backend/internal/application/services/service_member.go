package services

import (
	"backend/internal/application/command"
	"backend/internal/application/interfaces"
	"backend/internal/application/mapper"
	"backend/internal/domain/entities"
	"backend/internal/domain/repositories"
	"context"
)

type MemberRepos interface {
	Member() repositories.MemberRepo
	Invitation() repositories.InvitationRepo
	Server() repositories.ServerRepo
	User() repositories.UserRepo
}

type MemberService struct {
	uow repositories.UnitOfWork[MemberRepos]
}

func NewMemberService(uow repositories.UnitOfWork[MemberRepos]) interfaces.MembershipService {
	return &MemberService{uow}
}
func NewMemberQueries(uow repositories.UnitOfWork[MemberRepos]) interfaces.MembershipQueries {
	return &MemberService{uow}
}

func (s *MemberService) JoinServer(ctx context.Context, params command.JoinServerCommand) (res command.JoinServerCommandResult, err error) {
	err = s.uow.Do(ctx, func(ctx context.Context, repos MemberRepos) error {
		user, err := repos.User().Find(ctx, entities.UserId(params.UserId))
		if err != nil {
			return entities.GetErrOrDefault(err, entities.ErrCodeDepFail, "cannot get user")
		}

		inv, err := repos.Invitation().Find(ctx, entities.InvitationId(params.InvitationId))
		if err != nil {
			return entities.GetErrOrDefault(err, entities.ErrCodeDepFail, "cannot get invitation")
		}

		server, err := repos.Server().Find(ctx, inv.ServerId)
		if err != nil {
			return entities.GetErrOrDefault(err, entities.ErrCodeDepFail, "cannot get server, server may be deleted")
		}

		if server.NeedApproval && !inv.BypassApproval {
			return entities.NewError(entities.ErrCodeDepFail, "approval is not supported yet", nil)
		}

		if inv.JoinCount >= inv.JoinLimit && inv.JoinLimit > 0 {
			return entities.NewError(entities.ErrCodeDepFail, "Invitation expired", nil)
		}

		inv.JoinCount++
		membership := entities.NewMembership(server.Id, user.Id, user.DisplayName)
		membership, err = repos.Member().Save(ctx, membership)
		if err != nil {
			return entities.GetErrOrDefault(err, entities.ErrCodeDepFail, "cannot save membership")
		}
		inv, err = repos.Invitation().Save(ctx, inv)
		if err != nil {
			return entities.GetErrOrDefault(err, entities.ErrCodeDepFail, "cannot save invitation")
		}

		res = command.JoinServerCommandResult{
			Result: mapper.MembershipToResult(membership),
		}
		return nil
	})

	return res, err
}

func (s *MemberService) LeaveServer(ctx context.Context, params command.LeaveServerCommand) error {
	return s.uow.Do(ctx, func(ctx context.Context, repos MemberRepos) error {
		m, err := repos.Member().Find(ctx, entities.UserId(params.UserId), entities.ServerId(params.ServerId))
		if err != nil {
			return entities.GetErrOrDefault(err, entities.ErrCodeDepFail, "cannot get membership status")
		}
		if err = m.Delete(); err != nil {
			return err
		}

		// err := repos.Member().Delete(ctx, entities.UserId(params.UserId), entities.ServerId(params.ServerId))
		m, err = repos.Member().Save(ctx, m)
		return entities.GetErrOrDefault(err, entities.ErrCodeDepFail, "Cannot leave server")

	})
}

func (s *MemberService) Kick(context.Context, command.KickCommand) error {
	return entities.NewError(entities.ErrCodeDepFail, "not implemented", nil)
}

func (s *MemberService) Ban(context.Context, command.BanCommand) error {
	return entities.NewError(entities.ErrCodeDepFail, "not implemented", nil)
}

func (s *MemberService) SetNickname(context.Context, command.SetNickname) error {
	return entities.NewError(entities.ErrCodeDepFail, "not implemented", nil)
}
