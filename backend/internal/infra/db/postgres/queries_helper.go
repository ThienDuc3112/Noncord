package postgres

import (
	"github.com/google/uuid"
)

type getChannelsHelperParams struct {
	Channels      []channel
	RolePerms     []rolePerm
	RoleOverwrite []overwrite
	UserOverwrite []overwrite
}

type channel struct {
	id       uuid.UUID
	serverId uuid.UUID
}

type rolePerm struct {
	serverId   uuid.UUID
	roleId     uuid.UUID
	permission int64
}

type overwrite struct {
	allow     int64
	deny      int64
	channelId uuid.UUID
	targetId  uuid.UUID
}

func effectivePermPerChannel(params getChannelsHelperParams) map[uuid.UUID]int64 {
	rolesPermMap := make(map[uuid.UUID][]rolePerm)
	for _, role := range params.RolePerms {
		if _, ok := rolesPermMap[role.serverId]; !ok {
			rolesPermMap[role.serverId] = make([]rolePerm, 0)
		}
		rolesPermMap[role.serverId] = append(rolesPermMap[role.serverId], role)
	}

	rolesOverwriteMap := make(map[uuid.UUID]map[uuid.UUID]overwrite)
	for _, ro := range params.RoleOverwrite {
		if !in(rolesOverwriteMap, ro.channelId) {
			rolesOverwriteMap[ro.channelId] = make(map[uuid.UUID]overwrite)
		}
		rolesOverwriteMap[ro.channelId][ro.targetId] = ro
	}

	userOverwriteMap := make(map[uuid.UUID]overwrite)
	for _, uo := range params.UserOverwrite {
		userOverwriteMap[uo.channelId] = uo
	}

	res := make(map[uuid.UUID]int64)
	for _, c := range params.Channels {
		effPerm := int64(0)

		for _, role := range rolesPermMap[c.serverId] {
			effPerm |= role.permission
			if in(rolesOverwriteMap, c.id) && in(rolesOverwriteMap[c.id], role.roleId) {
				effPerm |= rolesOverwriteMap[c.id][role.roleId].allow
				effPerm &= (^rolesOverwriteMap[c.id][role.roleId].deny)
			}
		}

		if in(userOverwriteMap, c.id) {
			effPerm |= userOverwriteMap[c.id].allow
			effPerm &= (^userOverwriteMap[c.id].deny)
		}

		res[c.id] = effPerm
	}

	return res
}
