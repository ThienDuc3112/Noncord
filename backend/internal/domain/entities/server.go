package entities

import (
	"backend/internal/domain/events"
	"time"

	"github.com/google/uuid"
)

type ServerPermissionBits uint64

const (
	// General server perm
	PermViewChannel ServerPermissionBits = 1 << iota
	PermManageChannel
	PermManageRoles
	PermCreateEmote
	PermManageEmote
	PermViewAudit
	PermManageServer
	PermCreateInvite
	PermChangeNickname
	PermManageNickname
	PermManageMember
	PermBanMember
	PermTimeout

	// Text channel perm
	PermSendMessage
	PermEmbedLinks
	PermAttachFiles
	PermAddReactions
	PermExternalEmote
	PermMentionEveryone
	PermManageMessages
	PermReadMessagesHistory
	PermManagePermissions

	// Voice channel perm
	PermAdministrator
)

func CreatePermission(permissions ...ServerPermissionBits) ServerPermissionBits {
	var res ServerPermissionBits = 0
	for _, p := range permissions {
		res |= p
	}
	return res
}

// Supported many permission check by making check a combination of many permission bits
// Can use `CreatePermission` to create one
func (p ServerPermissionBits) HasAll(check ServerPermissionBits) bool {
	return (p & check) == check
}

// Supported many permission check by making check a combination of many permission bits
// Can use `CreatePermission` to create one
func (p ServerPermissionBits) HasAny(check ServerPermissionBits) bool {
	return (p & check) > 0
}

type ServerId uuid.UUID

type Server struct {
	events.Recorder

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

	DefaultPermission   ServerPermissionBits
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
		old := s.Name
		s.Name = newName
		s.UpdatedAt = time.Now()
		s.Record(NewServerNameUpdated(s, old))
	}
	return nil
}

func (s *Server) UpdateDescription(newDescription string) error {
	if len(newDescription) > 512 {
		return NewError(ErrCodeValidationError, "server description cannot exceed 512 characters", nil)
	}
	if s.Description != newDescription {
		old := s.Description
		s.Description = newDescription
		s.UpdatedAt = time.Now()
		s.Record(NewServerDescriptionUpdated(s, old))
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
		old := s.IconUrl
		s.IconUrl = newIconUrl
		s.UpdatedAt = time.Now()
		s.Record(NewServerIconURLUpdated(s, old))
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
		old := s.BannerUrl
		s.BannerUrl = newBannerUrl
		s.UpdatedAt = time.Now()
		s.Record(NewServerBannerURLUpdated(s, old))
	}
	return nil
}

func (s *Server) UpdateNeedApproval(needApproval bool) error {
	if s.NeedApproval != needApproval {
		old := s.NeedApproval
		s.NeedApproval = needApproval
		s.UpdatedAt = time.Now()
		s.Record(NewServerNeedApprovalChanged(s, old))
	}
	return nil
}

func (s *Server) UpdateAnnouncementChannel(channelId *ChannelId) error {
	changed := (s.AnnouncementChannel == nil && channelId != nil) ||
		(s.AnnouncementChannel != nil && channelId == nil) ||
		(s.AnnouncementChannel != nil && channelId != nil && *s.AnnouncementChannel != *channelId)

	if changed {
		old := s.AnnouncementChannel
		s.AnnouncementChannel = channelId
		s.UpdatedAt = time.Now()
		s.Record(NewServerAnnouncementChannelChanged(s, old))
	}
	return nil
}

func (s *Server) UpdateDefaultPermission(perm ServerPermissionBits) error {
	if s.DefaultPermission != perm {
		old := s.DefaultPermission
		s.DefaultPermission = perm
		s.UpdatedAt = time.Now()
		s.Record(NewServerDefaultPermissionChanged(s, old))
	}
	return nil
}

func (s *Server) IsOwner(userId UserId) bool {
	return userId == s.Owner
}

func (s *Server) Delete() error {
	now := time.Now()
	s.DeletedAt = &now
	s.Record(NewServerDeleted(s))
	return nil
}

func NewServer(userId UserId, name, description, iconUrl, bannerUrl string, needApproval bool) (*Server, error) {
	s := &Server{
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

		DefaultPermission:   CreatePermission(PermViewChannel, PermCreateInvite, PermChangeNickname, PermSendMessage, PermEmbedLinks, PermAttachFiles, PermAddReactions, PermExternalEmote, PermReadMessagesHistory),
		AnnouncementChannel: nil,
	}

	if err := s.Validate(); err != nil {
		return nil, err
	}

	s.Record(NewServerCreated(s))
	return s, nil
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
