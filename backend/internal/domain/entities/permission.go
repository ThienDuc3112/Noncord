package entities

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
