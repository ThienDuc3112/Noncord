package mapper

import (
	"backend/internal/application/common"
	"backend/internal/interface/rest/dto/response"
)

func ParseCommonMessage(m *common.Message) response.Message {
	if m == nil {
		return response.Message{}
	}
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
