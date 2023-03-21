package entities

import "time"

type Role struct {
	Id              RoleId
	Name            string
	PermissionValue PermissionValue
	CreatedAt       time.Time
	ServerId        ServerId
	CreatorId       UserId
	Permissions     []Permission
}

type RoleView struct {
	Id              RoleId          `json:"id"`
	Name            string          `json:"name"`
	PermissionValue PermissionValue `json:"permission-value"`
	CreatedAt       time.Time       `json:"create-at"`
	Permissions     []Permission    `json:"permissions"`
	Creator         *User           `json:"creator"`
}

func NewRoleByPermissions(permissions ...Permission) Role {
	r := Role{}
	for _, perm := range permissions {
		r.PermissionValue |= perm.Value
	}
	return r
}

func NewRoleByPermissionValues(permissionValues ...PermissionValue) Role {
	r := Role{}
	for _, permValue := range permissionValues {
		r.PermissionValue |= permValue
		r.Permissions = append(r.Permissions, Permission{Value: permValue})
	}
	return r
}
