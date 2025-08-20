package entities

import (
	"time"

	"github.com/google/uuid"
)

type ServerPermissionBits uint64

const (
	// General server perm
	PermViewChannel    ServerPermissionBits = 1 << 0
	PermManageChannel  ServerPermissionBits = 1 << 1
	PermManageRoles    ServerPermissionBits = 1 << 2
	PermCreateEmote    ServerPermissionBits = 1 << 3
	PermManageEmote    ServerPermissionBits = 1 << 4
	PermViewAudit      ServerPermissionBits = 1 << 5
	PermManageServer   ServerPermissionBits = 1 << 6
	PermCreateInvite   ServerPermissionBits = 1 << 7
	PermChangeNickname ServerPermissionBits = 1 << 8
	PermManageNickname ServerPermissionBits = 1 << 9
	PermManageMember   ServerPermissionBits = 1 << 10
	PermBanMember      ServerPermissionBits = 1 << 11
	PermTimeout        ServerPermissionBits = 1 << 12

	// Text channel perm
	PermSendMessage         ServerPermissionBits = 1 << 13
	PermEmbedLinks          ServerPermissionBits = 1 << 14
	PermAttachFiles         ServerPermissionBits = 1 << 15
	PermAddReactions        ServerPermissionBits = 1 << 16
	PermExternalEmote       ServerPermissionBits = 1 << 17
	PermMentionEveryone     ServerPermissionBits = 1 << 18
	PermManageMessages      ServerPermissionBits = 1 << 19
	PermReadMessagesHistory ServerPermissionBits = 1 << 20
	PermManagePermissions   ServerPermissionBits = 1 << 21

	// Voice channel perm
	PermAdministrator ServerPermissionBits = 1 << 22
)

type ServerId uuid.UUID

type Server struct {
	Id           ServerId
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    *time.Time
	Name         string
	Description  string
	IconUrl      string
	BannerUrl    string
	NeedApproval bool

	Owner UserId

	DefaultRole         *RoleId
	AnnouncementChannel *ChannelId
}

func (s *Server) Validate() error {
	if s.Name == "" {
		return NewError(ErrCodeValidationError, "cannot have empty server name", nil)
	}
	if len(s.Name) > 256 {
		return NewError(ErrCodeValidationError, "server name cannot exceed 256 characters", nil)
	}
	if len(s.Description) > 512 {
		return NewError(ErrCodeValidationError, "server description cannot exceed 512 characters", nil)
	}
	// TODO: Allow this when role exist
	// if s.DefaultRole == nil {
	// 	return NewError(ErrCodeValidationError, "server have no @everyone role", nil)
	// }
	if s.IconUrl != "" && !IsValidUrl(s.IconUrl) {
		return NewError(ErrCodeValidationError, "icon invalid url", nil)
	}
	if len(s.IconUrl) > 2048 {
		return NewError(ErrCodeValidationError, "icon url too long", nil)
	}
	if s.BannerUrl != "" && !IsValidUrl(s.BannerUrl) {
		return NewError(ErrCodeValidationError, "banner invalid url", nil)
	}
	if len(s.BannerUrl) > 2048 {
		return NewError(ErrCodeValidationError, "banner url too long", nil)
	}

	return nil
}

// Server's mutators
// # Owner UserId
func (s *Server) UpdateName(newName string) error {
	if newName == "" {
		return NewError(ErrCodeValidationError, "server name cannot be empty", nil)
	}
	if len(newName) > 256 {
		return NewError(ErrCodeValidationError, "server name cannot exceed 256 characters", nil)
	}
	if s.Name != newName { // Only update if changed
		s.Name = newName
		s.UpdatedAt = time.Now()
	}
	return nil
}

func (s *Server) UpdateDescription(newDescription string) error {
	if len(newDescription) > 512 {
		return NewError(ErrCodeValidationError, "server description cannot exceed 512 characters", nil)
	}
	if s.Description != newDescription {
		s.Description = newDescription
		s.UpdatedAt = time.Now()
	}
	return nil
}

func (s *Server) UpdateIconUrl(newIconUrl string) error {
	if newIconUrl != "" && !IsValidUrl(newIconUrl) {
		return NewError(ErrCodeValidationError, "icon invalid url", nil)
	}
	if len(newIconUrl) > 2048 {
		return NewError(ErrCodeValidationError, "icon url too long", nil)
	}
	if s.IconUrl != newIconUrl {
		s.IconUrl = newIconUrl
		s.UpdatedAt = time.Now()
	}
	return nil
}

func (s *Server) UpdateBannerUrl(newBannerUrl string) error {
	if newBannerUrl != "" && !IsValidUrl(newBannerUrl) {
		return NewError(ErrCodeValidationError, "banner invalid url", nil)
	}
	if len(newBannerUrl) > 2048 {
		return NewError(ErrCodeValidationError, "banner url too long", nil)
	}
	if s.BannerUrl != newBannerUrl {
		s.BannerUrl = newBannerUrl
		s.UpdatedAt = time.Now()
	}
	return nil
}

func (s *Server) UpdateNeedApproval(needApproval bool) error {
	if s.NeedApproval != needApproval {
		s.NeedApproval = needApproval
		s.UpdatedAt = time.Now()
	}
	return nil
}

func (s *Server) UpdateAnnouncementChannel(channelId *ChannelId) error {
	if s.AnnouncementChannel == nil || *s.AnnouncementChannel != *channelId {
		s.AnnouncementChannel = channelId
		s.UpdatedAt = time.Now()
	}
	return nil
}

func (s *Server) IsOwner(userId UserId) bool {
	return userId == s.Owner
}

func NewServer(userId UserId, name, description, iconUrl, bannerUrl string, needApproval bool) *Server {
	return &Server{
		Id:           ServerId(uuid.New()),
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
		DeletedAt:    nil,
		Name:         name,
		Description:  description,
		IconUrl:      iconUrl,
		BannerUrl:    bannerUrl,
		NeedApproval: needApproval,

		Owner: userId,

		DefaultRole:         nil,
		AnnouncementChannel: nil,
	}
}

type CategoryId uuid.UUID

type Category struct {
	Id        CategoryId
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
	ServerId  ServerId
	Name      string
	Order     uint16
}

func NewCategory(sid ServerId, name string, order uint16) *Category {
	return &Category{
		Id:        CategoryId(uuid.New()),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		DeletedAt: nil,
		ServerId:  sid,
		Name:      name,
		Order:     order,
	}
}

type JoinRequestId uuid.UUID

type JoinRequest struct {
	Id        JoinRequestId
	CreatedAt time.Time
	ServerId  ServerId
	UserId    UserId
}

func NewJoinRequest(sid ServerId, uid UserId) *JoinRequest {
	return &JoinRequest{
		Id:        JoinRequestId(uuid.New()),
		CreatedAt: time.Now(),
		ServerId:  sid,
		UserId:    uid,
	}
}
