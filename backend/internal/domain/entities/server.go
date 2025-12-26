package entities

import (
	"backend/internal/domain/events"
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

	// Admin
	PermAdministrator ServerPermissionBits = 1 << 63
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
func (p ServerPermissionBits) HasAll(check ...ServerPermissionBits) bool {
	for _, c := range check {
		if (p & c) != c {
			return false
		}
	}
	return true
}

// Supported many permission check by making check a combination of many permission bits
// Can use `CreatePermission` to create one
// Return true if it have a single common permission with a single check bits
func (p ServerPermissionBits) HasAny(check ...ServerPermissionBits) bool {
	for _, c := range check {
		if (p & c) > 0 {
			return true
		}
	}
	return false
}

func (p ServerPermissionBits) ToFlagArray() []string {
	res := make([]string, 0)

	if PermViewChannel&p > 0 {
		res = append(res, "ViewChannel")
	}
	if PermManageChannel&p > 0 {
		res = append(res, "ManageChannel")
	}
	if PermManageRoles&p > 0 {
		res = append(res, "ManageRoles")
	}
	if PermCreateEmote&p > 0 {
		res = append(res, "CreateEmote")
	}
	if PermManageEmote&p > 0 {
		res = append(res, "ManageEmote")
	}
	if PermViewAudit&p > 0 {
		res = append(res, "ViewAudit")
	}
	if PermManageServer&p > 0 {
		res = append(res, "ManageServer")
	}
	if PermCreateInvite&p > 0 {
		res = append(res, "CreateInvite")
	}
	if PermChangeNickname&p > 0 {
		res = append(res, "ChangeNickname")
	}
	if PermManageNickname&p > 0 {
		res = append(res, "ManageNickname")
	}
	if PermManageMember&p > 0 {
		res = append(res, "ManageMember")
	}
	if PermBanMember&p > 0 {
		res = append(res, "BanMember")
	}
	if PermTimeout&p > 0 {
		res = append(res, "Timeout")
	}
	if PermSendMessage&p > 0 {
		res = append(res, "SendMessage")
	}
	if PermEmbedLinks&p > 0 {
		res = append(res, "EmbedLinks")
	}
	if PermAttachFiles&p > 0 {
		res = append(res, "AttachFiles")
	}
	if PermAddReactions&p > 0 {
		res = append(res, "AddReactions")
	}
	if PermExternalEmote&p > 0 {
		res = append(res, "ExternalEmote")
	}
	if PermMentionEveryone&p > 0 {
		res = append(res, "MentionEveryone")
	}
	if PermManageMessages&p > 0 {
		res = append(res, "ManageMessages")
	}
	if PermReadMessagesHistory&p > 0 {
		res = append(res, "ReadMessagesHistory")
	}
	if PermManagePermissions&p > 0 {
		res = append(res, "ManagePermissions")
	}
	if PermAdministrator&p > 0 {
		res = append(res, "Administrator")
	}

	return res
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

	roleDirty bool

	Owner UserId

	DefaultRole         RoleId
	AnnouncementChannel *ChannelId

	Roles map[RoleId]*Role
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
	if r, ok := s.Roles[s.DefaultRole]; !ok || r == nil || r.DeletedAt != nil {
		return NewError(ErrCodeValidationError, "server have no @everyone role", nil)
	}
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

func (s *Server) IsRoleDirty() bool { return s.roleDirty }

// Server's mutators
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

		AnnouncementChannel: nil,

		Roles: make(map[RoleId]*Role),

		roleDirty: true,
	}

	everyoneRole := &Role{
		Id:           RoleId(uuid.New()),
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
		DeletedAt:    nil,
		Name:         "everyone",
		Color:        0x808080,
		Priority:     0,
		AllowMention: false,
		Permissions:  CreatePermission(PermViewChannel, PermCreateInvite, PermChangeNickname, PermSendMessage, PermEmbedLinks, PermAttachFiles, PermAddReactions, PermExternalEmote, PermReadMessagesHistory),
		ServerId:     s.Id,
		dirty:        true,
	}

	s.Roles[everyoneRole.Id] = everyoneRole
	s.DefaultRole = everyoneRole.Id

	s.Record(NewServerCreated(s))

	if err := s.Validate(); err != nil {
		return nil, err
	}

	return s, nil
}

func (s *Server) CreateRole(name string, color uint32, priority uint16, allowMention bool, perm ServerPermissionBits) (*Role, error) {
	r := &Role{
		Id:           RoleId(uuid.New()),
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
		DeletedAt:    nil,
		Name:         name,
		Color:        color,
		Priority:     priority,
		AllowMention: allowMention,
		Permissions:  perm,
		ServerId:     s.Id,
	}
	if err := r.Validate(); err != nil {
		return nil, err
	}

	s.roleDirty = true
	r.dirty = true

	s.Record(NewRoleCreated(r))
	s.Roles[r.Id] = r
	return r, nil
}

func (s *Server) DeleteRole(id RoleId) error {
	r, ok := s.Roles[id]
	if !ok || r == nil {
		return nil
	}

	s.roleDirty = true
	r.dirty = true

	now := time.Now()
	r.DeletedAt = &now
	s.Record(NewRoleDeleted(r))
	return nil
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

type RoleId uuid.UUID

type Role struct {
	Id           RoleId
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    *time.Time
	Name         string
	Color        uint32
	Priority     uint16
	AllowMention bool
	Permissions  ServerPermissionBits
	ServerId     ServerId
	dirty        bool
}

func (r *Role) Validate() error {
	if r.Name == "" {
		return NewError(ErrCodeValidationError, "name cannot be empty", nil)
	}
	if len(r.Name) > 64 {
		return NewError(ErrCodeValidationError, "name cannot exceed 64 characters", nil)
	}
	return nil
}

func (r *Role) IsDirty() bool {
	return r.dirty
}

func (s *Server) UpdateRoleName(rid RoleId, name string) error {
	role, ok := s.Roles[rid]
	if !ok {
		return NewError(ErrCodeValidationError, "role don't exist", nil)
	}
	if name == role.Name {
		return nil
	}
	s.roleDirty = true
	role.dirty = true
	role.UpdatedAt = time.Now()

	role.Name, name = name, role.Name
	s.Record(NewRoleNameUpdated(role, name))

	return role.Validate()
}

func (s *Server) UpdateRoleColor(rid RoleId, color uint32) error {
	role, ok := s.Roles[rid]
	if !ok {
		return NewError(ErrCodeValidationError, "role don't exist", nil)
	}

	if color == role.Color {
		return nil
	}
	s.roleDirty = true
	role.dirty = true
	role.UpdatedAt = time.Now()

	role.Color, color = color, role.Color
	s.Record(NewRoleColorUpdated(role, color))

	return role.Validate()
}

func (s *Server) UpdateRoleAllowMention(rid RoleId, allow bool) error {
	role, ok := s.Roles[rid]
	if !ok {
		return NewError(ErrCodeValidationError, "role don't exist", nil)
	}

	if allow == role.AllowMention {
		return nil
	}
	s.roleDirty = true
	role.dirty = true
	role.UpdatedAt = time.Now()

	role.AllowMention, allow = allow, role.AllowMention
	s.Record(NewRoleAllowMentionUpdated(role, allow))

	return role.Validate()
}

func (s *Server) UpdateRolePermissions(rid RoleId, perms ServerPermissionBits) error {
	role, ok := s.Roles[rid]
	if !ok {
		return NewError(ErrCodeValidationError, "role don't exist", nil)
	}

	s.roleDirty = true
	role.dirty = true
	role.UpdatedAt = time.Now()

	role.Permissions, perms = perms, role.Permissions
	s.Record(NewRolePermissionsUpdated(role, perms))

	return role.Validate()
}

func (s *Server) ReordereRole() {}
