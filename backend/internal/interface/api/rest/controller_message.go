package rest

import (
	"backend/internal/application/command"
	"backend/internal/application/interfaces"
	"backend/internal/application/query"
	"backend/internal/interface/api/rest/dto/request"
	"backend/internal/interface/api/rest/dto/response"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/google/uuid"
)

type MessageController struct {
	messageService interfaces.MessageService
	authService    interfaces.AuthService
}

func NewMessageController(service interfaces.MessageService, authService interfaces.AuthService) *MessageController {
	return &MessageController{service, authService}
}

func (ac *MessageController) RegisterRoute(r chi.Router) {
	r.Route("/message", func(r chi.Router) {
		r.Use(authMiddleware(ac.authService))

		r.Post("/", ac.CreateMessageController)
		r.Get("/{message_id}", ac.GetMessageController)
		r.Get("/channel/{channel_id}", ac.GetMessagesByChannelIdController)
		r.Get("/group/{group_id}", ac.GetMessagesByGroupIdController)
	})
}

// register 		godoc
//
//	@Summary		Send a message
//	@Description	Send a message
//	@Tags			Message
//	@Accept			json
//	@Produce		json
//	@Param			payload	body		request.CreateMessage	true	"Message content"
//	@Param			Authorization	header		string						true	"Bearer token"
//	@Success		201		{object}	response.Message					"Message sent"
//	@Failure		400		{object}	response.ErrorResponse "Invalid request body"
//	@Failure		401		{object}	response.ErrorResponse "Unauthorized"
//	@Failure		403		{object}	response.ErrorResponse "User not allowed to send message in the request channel/group"
//	@Failure		404		{object}	response.ErrorResponse "Target channel/group not found"
//	@Failure		500		{object}	response.ErrorResponse
//	@Router			/api/v1/message [post]
func (ac *MessageController) CreateMessageController(w http.ResponseWriter, r *http.Request) {
	log.Println("[CreateMessageController] Create message")

	user := extractUser(r.Context())
	if user == nil {
		render.Render(w, r, response.ParseErrorResponse("Cannot authenticate user", http.StatusUnauthorized, nil))
		return
	}

	var body request.CreateMessage
	if err := render.Bind(r, &body); err != nil {
		render.Render(w, r, response.ParseErrorResponse("Invalid body", http.StatusBadRequest, err))
		return
	}

	msg, err := ac.messageService.Create(r.Context(), command.CreateMessageCommand{
		UserId:          user.Id,
		TargetId:        body.TargetId,
		Content:         body.Content,
		IsTargetChannel: body.IsTargetChannel,
	})
	if err != nil {
		render.Render(w, r, response.ParseErrorResponse("Cannot create message", 500, err))
		return
	}

	render.Status(r, 201)
	render.JSON(w, r, response.Message{
		Id:        msg.Result.Id,
		CreatedAt: msg.Result.CreatedAt,
		UpdatedAt: msg.Result.UpdatedAt,
		ChannelId: msg.Result.ChannelId,
		GroupId:   msg.Result.GroupId,
		Author:    msg.Result.Author,
		Message:   msg.Result.Message,
	})
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
//	@Success		200		{object}	response.Message
//	@Failure		400		{object}	response.ErrorResponse "Invalid message id"
//	@Failure		401		{object}	response.ErrorResponse "Unauthorized"
//	@Failure		404		{object}	response.ErrorResponse "Message not found"
//	@Failure		500		{object}	response.ErrorResponse
//	@Router			/api/v1/message/{message_id} [get]
func (ac *MessageController) GetMessageController(w http.ResponseWriter, r *http.Request) {
	log.Println("[GetMessageController] Get message")

	user := extractUser(r.Context())
	if user == nil {
		render.Render(w, r, response.ParseErrorResponse("Cannot authenticate user", http.StatusUnauthorized, nil))
		return
	}

	messageId, err := uuid.Parse(chi.URLParam(r, "message_id"))
	if err != nil {
		render.Render(w, r, response.ParseErrorResponse("Invalid message id", http.StatusBadRequest, err))
		return
	}

	msg, err := ac.messageService.Get(r.Context(), query.GetMessage{
		MessageId: messageId,
		UserId:    user.Id,
	})
	if err != nil {
		render.Render(w, r, response.ParseErrorResponse("Cannot get message", 500, err))
		return
	}

	render.JSON(w, r, response.Message{
		Id:        msg.Result.Id,
		CreatedAt: msg.Result.CreatedAt,
		UpdatedAt: msg.Result.UpdatedAt,
		ChannelId: msg.Result.ChannelId,
		GroupId:   msg.Result.GroupId,
		Author:    msg.Result.Author,
		Message:   msg.Result.Message,
	})
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
