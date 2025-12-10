package ws

import (
	"backend/internal/application/ports"
	"backend/internal/domain/entities"
	"backend/internal/domain/events"
	"backend/internal/interface/dto/response"
	"context"
	"fmt"
	"log/slog"
)

const (
	incomingMessage = "incoming_message"
)

func (h *Hub) messageCreatedHandler(ctx context.Context, event ports.EventMessage) error {
	slog.Default().Info("Incoming message", "event", slog.GroupValue(
		slog.Attr{
			Key:   "aggregatedId",
			Value: slog.AnyValue(event.AggregateId),
		},
		slog.Attr{
			Key:   "eventType",
			Value: slog.AnyValue(event.EventType),
		},
		slog.Attr{
			Key:   "headers",
			Value: slog.AnyValue(event.Headers),
		},
	))

	e, err := events.ParseSpecificEvent[entities.MessageCreated](event.Payload, entities.EventMessageCreated, entities.MessageCreatedSchemaVersion)
	if err != nil {
		slog.Default().Warn("Unabled to parse event", "error", err)
		return err
	}

	slog.Default().Info("Parsed event", "event", e)
	h.m.RLock()
	defer h.m.RUnlock()
	message := response.Message{
		Id:          e.AggregateID,
		CreatedAt:   e.OccurredAt,
		UpdatedAt:   e.OccurredAt,
		ChannelId:   e.ChannelID,
		GroupId:     e.GroupID,
		Author:      e.AuthorID,
		AuthorType:  e.AuthorType,
		Message:     e.Content,
		DisplayName: "",
		AvatarUrl:   "",
	}

	key := ""
	if e.AuthorID != nil && e.ChannelID != nil {
		key = fmt.Sprintf("user_enrichment.channel.%s.%s", e.AuthorID, e.ChannelID)
	} else if e.AuthorID != nil {
		key = fmt.Sprintf("user_enrichment.user.%s", e.AuthorID)
	}

	if key != "" {
		data, exist := h.nicknameCache.Get(key)
		enrichment, ok := data.(ports.UserEnrichment)
		if exist && ok {
			message.AvatarUrl = enrichment.AvatarUrl
			message.DisplayName = enrichment.Nickname
		} else {
			var hydrate ports.UserEnrichment
			if e.ChannelID != nil {
				hydrate, err = h.userResolver.FromUserChannel(ctx, *e.AuthorID, *e.ChannelID)
			} else {
				hydrate, err = h.userResolver.FromUser(ctx, *e.AuthorID)
			}
			if err != nil {
				slog.Warn("Cannot get user enrichment details", "error", err)
			} else {
				h.nicknameCache.Set(key, hydrate)
				message.AvatarUrl = hydrate.AvatarUrl
				message.DisplayName = hydrate.Nickname
			}
		}
	}

	if e.ChannelID != nil {
		if _, ok := h.channelSub[*e.ChannelID]; !ok {
			slog.Default().Info("no listener on channel", "channel_id", e.ChannelID.String())
			return nil
		}

		for uId := range h.channelSub[*e.ChannelID] {
			for _, c := range h.userConn[uId] {
				c.Write(incomingMessage, message)
			}
		}
	}

	return nil
}

func (h *Hub) messageEditedHandler(ctx context.Context, event ports.EventMessage) error {
	return nil
}

func (h *Hub) messageDeletedHandler(ctx context.Context, event ports.EventMessage) error {
	return nil
}
