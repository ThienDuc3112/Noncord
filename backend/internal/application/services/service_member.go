package services

import (
	"backend/internal/application/command"
	"backend/internal/application/interfaces"
	"backend/internal/application/mapper"
	"backend/internal/domain/entities"
	"backend/internal/domain/repositories"
	"context"
)

type MemberService struct {
	mr repositories.MemberRepo
	ir repositories.InvitationRepo
	sr repositories.ServerRepo
	ur repositories.UserRepo
}

func NewMemberService(mr repositories.MemberRepo, ir repositories.InvitationRepo, sr repositories.ServerRepo, ur repositories.UserRepo) interfaces.MembershipService {
	return &MemberService{
		mr: mr,
		ir: ir,
		sr: sr,
		ur: ur,
	}
}

func (s *MemberService) JoinServer(ctx context.Context, params command.JoinServerCommand) (command.JoinServerCommandResult, error) {
	user, err := s.ur.Find(ctx, entities.UserId(params.UserId))
	if err != nil {
		return command.JoinServerCommandResult{}, entities.GetErrOrDefault(err, entities.ErrCodeDepFail, "cannot get user")
	}

	inv, err := s.ir.Find(ctx, entities.InvitationId(params.InvitationId))
	if err != nil {
		return command.JoinServerCommandResult{}, entities.GetErrOrDefault(err, entities.ErrCodeDepFail, "cannot get invitation")
	}

	server, err := s.sr.Find(ctx, inv.ServerId)
	if err != nil {
		return command.JoinServerCommandResult{}, entities.GetErrOrDefault(err, entities.ErrCodeDepFail, "cannot get server, server may be deleted")
	}

	if server.NeedApproval && !inv.BypassApproval {
		return command.JoinServerCommandResult{}, entities.NewError(entities.ErrCodeDepFail, "approval is not supported yet", nil)
	}

	if inv.JoinCount >= inv.JoinLimit && inv.JoinLimit > 0 {
		return command.JoinServerCommandResult{}, entities.NewError(entities.ErrCodeDepFail, "Invitation expired", nil)
	}

	inv.JoinCount++
	membership := entities.NewMembership(server.Id, user.Id, user.DisplayName)
	membership, err = s.mr.Save(ctx, membership)
	if err != nil {
		return command.JoinServerCommandResult{}, entities.GetErrOrDefault(err, entities.ErrCodeDepFail, "cannot save membership")
	}
	inv, err = s.ir.Save(ctx, inv)
	if err != nil {
		return command.JoinServerCommandResult{}, entities.GetErrOrDefault(err, entities.ErrCodeDepFail, "cannot save invitation")
	}

	return command.JoinServerCommandResult{
		Result: mapper.MembershipToResult(membership),
	}, nil
}

func (s *MemberService) LeaveServer(ctx context.Context, params command.LeaveServerCommand) error {
	err := s.mr.Delete(ctx, entities.UserId(params.UserId), entities.ServerId(params.ServerId))
	return entities.GetErrOrDefault(err, entities.ErrCodeDepFail, "Cannot leave server")
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
