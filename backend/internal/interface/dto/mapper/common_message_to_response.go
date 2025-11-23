package mapper

import (
	"backend/internal/application/common"
	"backend/internal/application/query"
	"backend/internal/interface/dto/response"
)

func ParseEnrichedMessage(m query.EnrichedMessage) response.Message {
	return response.Message{
		Id:          m.Id,
		CreatedAt:   m.CreatedAt,
		UpdatedAt:   m.UpdatedAt,
		ChannelId:   m.ChannelId,
		GroupId:     m.GroupId,
		Author:      m.Author,
		AuthorType:  m.AuthorType,
		Message:     m.Message.Message,
		DisplayName: m.Nickname,
		AvatarUrl:   m.AvatarUrl,
	}
}

func ParseCommonMessage(m *common.Message) response.Message {
	return response.Message{
		Id:         m.Id,
		CreatedAt:  m.CreatedAt,
		UpdatedAt:  m.UpdatedAt,
		ChannelId:  m.ChannelId,
		GroupId:    m.GroupId,
		Author:     m.Author,
		AuthorType: m.AuthorType,
		Message:    m.Message,
	}
}
