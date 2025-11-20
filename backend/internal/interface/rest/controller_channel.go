package rest

import (
	"backend/internal/application/command"
	"backend/internal/application/interfaces"
	"backend/internal/application/query"
	"backend/internal/interface/rest/dto/request"
	"backend/internal/interface/rest/dto/response"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/google/uuid"
)

type ChannelController struct {
	authService    interfaces.AuthService
	channelService interfaces.ChannelService
}

func NewChannelController(authService interfaces.AuthService, channelService interfaces.ChannelService) *ChannelController {
	return &ChannelController{
		authService:    authService,
		channelService: channelService,
	}
}

func (c *ChannelController) RegisterRoute(r chi.Router) {
	r.Route("/channels", func(r chi.Router) {
		r.Use(authMiddleware(c.authService))

		r.Post("/", c.CreateChannelController)

		r.Get("/{channel_id}", c.GetChannelController)
		r.Put("/{channel_id}", c.UpdateChannelController)
		r.Patch("/{channel_id}", c.UpdateChannelController)
		r.Delete("/{channel_id}", c.DeleteChannelController)
	})
}

// register 		godoc
//
//	@Summary		Get channel details
//	@Description	Get channel details
//	@Tags			Channel
//	@Produce		json
//	@Param			Authorization	header		string						true	"Bearer token"
//	@Param			channel_id	path		string	true	"channel id to fetch"
//	@Success		200			{object}	response.Channel
//	@Failure		400			{object}	response.ErrorResponse	"Invalid channel id"
//	@Failure		404			{object}	response.ErrorResponse	"Channel not found"
//	@Failure		500			{object}	response.ErrorResponse
//	@Router			/api/v1/channels/{channel_id} [get]
func (c *ChannelController) GetChannelController(w http.ResponseWriter, r *http.Request) {
	log.Println("[GetChannelController] Getting channel")

	channelId, err := uuid.Parse(chi.URLParam(r, "channel_id"))
	if err != nil {
		render.Render(w, r, response.ParseErrorResponse("Invalid server id", http.StatusBadRequest, err))
		return
	}

	userId := extractUserId(r.Context())
	if userId == nil {
		render.Render(w, r, response.ParseErrorResponse("Cannot authenticate user", http.StatusUnauthorized, nil))
		return
	}

	channel, err := c.channelService.Get(r.Context(), query.GetChannel{
		ChannelId: channelId,
		UserId:    *userId,
	})

	if err != nil {
		render.Render(w, r, response.ParseErrorResponse("Cannot get channel", 500, err))
		return
	}

	render.Status(r, 200)
	render.JSON(w, r, response.Channel{
		Id:             channel.Result.Id,
		CreatedAt:      channel.Result.CreatedAt,
		UpdatedAt:      channel.Result.UpdatedAt,
		Name:           channel.Result.Name,
		Description:    channel.Result.Description,
		ServerId:       channel.Result.ServerId,
		Order:          channel.Result.Order,
		ParentCategory: channel.Result.ParentCategory,
	})
}

// register 		godoc
//
//	@Summary		Create channel
//	@Description	Create channel
//	@Tags			Channel
//	@Produce		json
//	@Param			Authorization	header		string						true	"Bearer token"
//	@Param			channel_id	path		string						true	"channel id to fetch"
//	@Param			payload			body		request.CreateChannel	true	"payload"
//	@Success		200				{object}	response.Channel			"Created channel"
//	@Failure		400				{object}	response.ErrorResponse		"Invalid body or invalid channel id"
//	@Failure		401				{object}	response.ErrorResponse		"Cannot authenticate user"
//	@Failure		403				{object}	response.ErrorResponse		"Forbidden action"
//	@Failure		404				{object}	response.ErrorResponse		"Channel not found"
//	@Failure		500				{object}	response.ErrorResponse
//	@Router			/api/v1/channels [post]
func (c *ChannelController) CreateChannelController(w http.ResponseWriter, r *http.Request) {
	log.Println("[CreateChannelController] Creating channel")

	userId := extractUserId(r.Context())
	if userId == nil {
		render.Render(w, r, response.ParseErrorResponse("Cannot authenticate user", http.StatusUnauthorized, nil))
		return
	}

	body := request.CreateChannel{}
	if err := render.Bind(r, &body); err != nil {
		render.Render(w, r, response.ParseErrorResponse("Invalid body", http.StatusBadRequest, err))
		return
	}

	channel, err := c.channelService.Create(r.Context(), command.CreateChannelCommand{
		Name:           body.Name,
		Description:    body.Description,
		ServerId:       body.ServerId,
		ParentCategory: body.ParentCategory,
		UserId:         *userId,
	})
	if err != nil {
		render.Render(w, r, response.ParseErrorResponse("Cannot create channel", http.StatusInternalServerError, err))
		return
	}

	render.Status(r, 200)
	render.JSON(w, r, response.Channel{
		Id:             channel.Result.Id,
		CreatedAt:      channel.Result.CreatedAt,
		UpdatedAt:      channel.Result.UpdatedAt,
		Name:           channel.Result.Name,
		Description:    channel.Result.Description,
		ServerId:       channel.Result.ServerId,
		Order:          channel.Result.Order,
		ParentCategory: channel.Result.ParentCategory,
	})
}

// register 		godoc
//
//	@Summary		Update channel
//	@Description	Update channel
//	@Tags			Channel
//	@Produce		json
//	@Param			Authorization	header		string						true	"Bearer token"
//	@Param			channel_id	path		string						true	"channel id"
//	@Param			payload			body		request.UpdateChannel	true	"payload"
//	@Success		200				{object}	response.Channel			"New updated channel"
//	@Failure		400				{object}	response.ErrorResponse		"Invalid body or invalid channel id"
//	@Failure		401				{object}	response.ErrorResponse		"Cannot authenticate user"
//	@Failure		403				{object}	response.ErrorResponse		"Forbidden action"
//	@Failure		404				{object}	response.ErrorResponse		"Channel not found"
//	@Failure		500				{object}	response.ErrorResponse
//	@Router			/api/v1/channels/{channel_id} [put]
//	@Router			/api/v1/channels/{channel_id} [patch]
func (c *ChannelController) UpdateChannelController(w http.ResponseWriter, r *http.Request) {
	log.Println("[UpdateChannelController] Updating channel")
}

// register 		godoc
//
//	@Summary		Delete channel
//	@Description	Delete channel
//	@Tags			Channel
//	@Produce		json
//	@Param			Authorization	header		string						true	"Bearer token"
//	@Param			channel_id	path		string						true	"channel id"
//	@Success		204				{object}	nil
//	@Failure		400				{object}	response.ErrorResponse		"invalid channel id"
//	@Failure		401				{object}	response.ErrorResponse		"Cannot authenticate user"
//	@Failure		403				{object}	response.ErrorResponse		"Forbidden action"
//	@Failure		404				{object}	response.ErrorResponse		"Channel not found"
//	@Failure		500				{object}	response.ErrorResponse
//	@Router			/api/v1/invitations/{channel_id} [delete]
func (c *ChannelController) DeleteChannelController(w http.ResponseWriter, r *http.Request) {
	log.Println("[DeleteChannelController] Deleting channel")
}
