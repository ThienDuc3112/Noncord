package services

import (
	"backend/internal/application/command"
	"backend/internal/application/interfaces"
	"backend/internal/application/mapper"
	"backend/internal/domain/entities"
	"backend/internal/domain/repositories"
	"context"
	"log/slog"

	"github.com/google/uuid"
)

type ServerRepos interface {
	Channel() repositories.ChannelRepo
	Server() repositories.ServerRepo
	Member() repositories.MemberRepo
	Role() repositories.RoleRepo
}

type ServerService struct {
	uow repositories.UnitOfWork[ServerRepos]
}

func NewServerService(uow repositories.UnitOfWork[ServerRepos]) interfaces.ServerService {
	return &ServerService{uow}
}

func (s *ServerService) Create(ctx context.Context, params command.CreateServerCommand) (res command.CreateServerCommandResult, err error) {
	server, err := entities.NewServer(entities.UserId(params.UserId), params.Name, "", "", "", false)
	if err != nil {
		return command.CreateServerCommandResult{}, err
	}
	role, err := entities.NewRole("everyone", 0x808080, 0, false,
		entities.CreatePermission(
			entities.PermViewChannel,
			entities.PermCreateInvite,
			entities.PermChangeNickname,
			entities.PermSendMessage,
			entities.PermEmbedLinks,
			entities.PermAttachFiles,
			entities.PermAddReactions,
			entities.PermExternalEmote,
			entities.PermReadMessagesHistory,
		),
		server.Id,
	)
	if err != nil {
		return command.CreateServerCommandResult{}, err
	}

	err = s.uow.Do(ctx, func(ctx context.Context, repos ServerRepos) error {
		slog.Info("creating server")
		server, err = repos.Server().Save(ctx, server)
		if err != nil {
			return entities.GetErrOrDefault(err, entities.ErrCodeDepFail, "cannot save server")
		}

		slog.Info("creating channel")
		channel := entities.NewChannel("text channel", "Your first channel", server.Id, 1, nil)
		channel, err = repos.Channel().Save(ctx, channel)
		if err != nil {
			return entities.GetErrOrDefault(err, entities.ErrCodeDepFail, "cannot save channel")
		}

		if err = server.UpdateAnnouncementChannel(&channel.Id); err != nil {
			return err
		}

		slog.Info("saving server", "server", server)
		server, err = repos.Server().Save(ctx, server)
		if err != nil {
			return entities.GetErrOrDefault(err, entities.ErrCodeDepFail, "cannot save server")
		}

		slog.Info("creating everyone role")
		role, err = repos.Role().Save(ctx, role)
		if err != nil {
			return entities.GetErrOrDefault(err, entities.ErrCodeDepFail, "cannot save role")
		}

		slog.Info("creating membership")
		membership := entities.NewMembership(server.Id, entities.UserId(params.UserId), params.UserNickname)
		membership.AssignRole(role.Id)
		membership, err = repos.Member().Save(ctx, membership)
		if err != nil {
			return entities.GetErrOrDefault(err, entities.ErrCodeDepFail, "cannot save membership")
		}

		res = command.CreateServerCommandResult{
			Result: mapper.ServerToResult(server),
		}
		return nil
	})

	return res, err
}

func (s *ServerService) Update(ctx context.Context, params command.UpdateServerCommand) (res command.UpdateServerCommandResult, err error) {
	err = s.uow.Do(ctx, func(ctx context.Context, repos ServerRepos) error {
		server, err := repos.Server().Find(ctx, entities.ServerId(params.ServerId))
		if err != nil {
			return entities.GetErrOrDefault(err, entities.ErrCodeDepFail, "cannot get server")
		}

		if params.Updates.Name != nil {
			if err = server.UpdateName(*params.Updates.Name); err != nil {
				return err
			}
		}
		if params.Updates.Description != nil {
			if err = server.UpdateDescription(*params.Updates.Description); err != nil {
				return err
			}
		}
		if params.Updates.IconUrl != nil {
			if err = server.UpdateIconUrl(*params.Updates.IconUrl); err != nil {
				return err
			}
		}
		if params.Updates.BannerUrl != nil {
			if err = server.UpdateBannerUrl(*params.Updates.BannerUrl); err != nil {
				return err
			}
		}
		if params.Updates.NeedApproval != nil {
			if err = server.UpdateNeedApproval(*params.Updates.NeedApproval); err != nil {
				return err
			}
		}
		if params.Updates.AnnouncementChannel.Valid {
			if err = server.UpdateAnnouncementChannel((*entities.ChannelId)(&params.Updates.AnnouncementChannel.UUID)); err != nil {
				return err
			}
		}
		if params.Updates.DefaultRole != nil {
			newRole := (*uuid.UUID)(nil)
			if params.Updates.DefaultRole.Valid {
				newRole = (*uuid.UUID)(&params.Updates.DefaultRole.UUID)
			}
			if err = server.UpdateDefaultRole((*entities.RoleId)(newRole)); err != nil {
				return err
			}
		}

		server, err = repos.Server().Save(ctx, server)
		if err != nil {
			return entities.GetErrOrDefault(err, entities.ErrCodeDepFail, "cannot update server")
		}

		res = command.UpdateServerCommandResult{
			Result: mapper.ServerToResult(server),
		}
		return nil
	})

	return res, err
}

func (s *ServerService) Delete(ctx context.Context, param command.DeleteServerCommand) error {
	return s.uow.Do(ctx, func(ctx context.Context, repos ServerRepos) error {
		server, err := repos.Server().Find(ctx, entities.ServerId(param.ServerId))
		if err != nil {
			return entities.GetErrOrDefault(err, entities.ErrCodeDepFail, "cannot get server")
		}

		if server.Owner != entities.UserId(param.UserId) {
			return entities.NewError(entities.ErrCodeForbidden, "user is not the owner of the server", nil)
		}

		server.Delete()
		server, err = repos.Server().Save(ctx, server)
		return entities.GetErrOrDefault(err, entities.ErrCodeDepFail, "cannot delete server")
	})
}
