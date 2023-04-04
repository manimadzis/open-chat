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
