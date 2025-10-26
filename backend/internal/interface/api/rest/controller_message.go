package rest

import (
	"backend/internal/application/interfaces"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type MessageController struct {
	messageService interfaces.MessageService
}

func NewMessageController(service interfaces.MessageService) *MessageController {
	return &MessageController{service}
}

func (ac *MessageController) RegisterRoute(r chi.Router) {
	r.Route("/message", func(r chi.Router) {
	})
}

// register 		godoc
//
//	@Summary		Send a message
//	@Description	Send a message
//	@Tags			Message
//	@Accept			json
//	@Produce		json
//	@Param			payload	body		request.Register	true	"Message content"
//	@Param			Authorization	header		string						true	"Bearer token"
//	@Success		201		{object}	nil					"Message sent"
//	@Failure		400		{object}	response.ErrorResponse "Invalid request body"
//	@Failure		401		{object}	response.ErrorResponse "Unauthorized"
//	@Failure		403		{object}	response.ErrorResponse "User not allowed to send message in the request channel/group"
//	@Failure		404		{object}	response.ErrorResponse "Target channel/group not found"
//	@Failure		500		{object}	response.ErrorResponse
//	@Router			/api/v1/message [post]
func (ac *MessageController) CreateMessageController(w http.ResponseWriter, r *http.Request) {
	log.Println("[CreateMessageController] Create message")
}

// register 		godoc
//
//	@Summary		Get a message
//	@Description	Get a message details
//	@Tags			Message
//	@Accept			json
//	@Produce		json
//	@Param			Authorization	header		string						true	"Bearer token"
//	@Param			message_id		path		string	true	"message id to fetch"
//	@Success		200		{object}	nil
//	@Failure		400		{object}	response.ErrorResponse "Invalid message id"
//	@Failure		401		{object}	response.ErrorResponse "Unauthorized"
//	@Failure		404		{object}	response.ErrorResponse "Message not found"
//	@Failure		500		{object}	response.ErrorResponse
//	@Router			/api/v1/message/{message_id} [get]
func (ac *MessageController) GetMessageController(w http.ResponseWriter, r *http.Request) {
	log.Println("[GetMessageController] Get message")
}

// register 		godoc
//
//	@Summary		Get messages by channel id
//	@Description	Get messages in a channel using channel id, default limit to 100
//	@Tags			Message
//	@Accept			json
//	@Produce		json
//	@Param			Authorization	header		string						true	"Bearer token"
//	@Param			channel_id		path		string	true	"channel id to fetch messages"
//	@Param			limit		query		int	false	"Message limit" minimum(1) maximum(500) default(100)
//	@Param			before		query		int	false	"Time in unix"
//	@Success		200		{object}	nil
//	@Failure		400		{object}	response.ErrorResponse "Invalid channel id"
//	@Failure		401		{object}	response.ErrorResponse "Unauthorized"
//	@Failure		404		{object}	response.ErrorResponse "Channel not found"
//	@Failure		500		{object}	response.ErrorResponse
//	@Router			/api/v1/message/channel/{channel_id} [get]
func (ac *MessageController) GetMessagesByChannelIdController(w http.ResponseWriter, r *http.Request) {
	log.Println("[GetMessagesByChannelIdController] Getting messages by channel id")
}

// register 		godoc
//
//	@Summary		Get messages by group id
//	@Description	Get messages in a group using group id, default limit to 100
//	@Tags			Message
//	@Accept			json
//	@Produce		json
//	@Param			Authorization	header		string						true	"Bearer token"
//	@Param			group_id		path		string	true	"channel id to fetch messages"
//	@Param			limit		query		int	false	"Message limit" minimum(1) maximum(500) default(100)
//	@Param			before		query		int	false	"Time in unix"
//	@Success		200		{object}	nil
//	@Failure		400		{object}	response.ErrorResponse "Invalid group id"
//	@Failure		401		{object}	response.ErrorResponse "Unauthorized"
//	@Failure		404		{object}	response.ErrorResponse "Group not found"
//	@Failure		500		{object}	response.ErrorResponse
//	@Router			/api/v1/message/group/{group_id} [get]
func (ac *MessageController) GetMessagesByGroupIdController(w http.ResponseWriter, r *http.Request) {
	log.Println("[GetMessagesByGroupIdController] Getting messages by group id")
}
