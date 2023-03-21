package role_system

import (
	"github.com/stretchr/testify/require"
	"open-chat/internal/entities"
	"testing"
)

func TestRoleSystem_Check(t *testing.T) {
	tests := []struct {
		role                []entities.Role
		requiredPermissions []entities.PermissionValue
		result              error
	}{
		{
			[]entities.Role{entities.NewRoleByPermissionValues(PERM_ADD_FILE)},
			[]entities.PermissionValue{},
			nil,
		},
		{
			[]entities.Role{entities.NewRoleByPermissionValues(PERM_ADD_FILE)},
			[]entities.PermissionValue{PERM_ADD_FILE},
			nil,
		},
		{
			[]entities.Role{entities.NewRoleByPermissionValues(PERM_DELETE_MESSAGE, PERM_DELETE_CHANNEL)},
			[]entities.PermissionValue{PERM_DELETE_MESSAGE, PERM_DELETE_CHANNEL},
			nil},
		{[]entities.Role{
			entities.NewRoleByPermissionValues(PERM_DELETE_MESSAGE, PERM_INVITE_USER),
			entities.NewRoleByPermissionValues(PERM_ADD_STICKER, PERM_CHANGE_ROLE),
		},
			[]entities.PermissionValue{PERM_CHANGE_ROLE},
			nil,
		},
		{
			[]entities.Role{
				entities.NewRoleByPermissionValues(),
				entities.NewRoleByPermissionValues(PERM_DELETE_MESSAGE, PERM_INVITE_USER),
			},
			[]entities.PermissionValue{PERM_ADD_FILE},
			ErrNotEnoughPermissions},
		{
			[]entities.Role{entities.NewRoleByPermissionValues(PERM_DELETE_MESSAGE),
				entities.NewRoleByPermissionValues(PERM_DELETE_MESSAGE, PERM_INVITE_USER),
			},
			[]entities.PermissionValue{PERM_DELETE_MESSAGE, PERM_DELETE_CHANNEL},
			ErrNotEnoughPermissions,
		},
	}

	for _, test := range tests {
		t.Run("", func(t *testing.T) {
			err := NewRoleSystem(test.role).Check(test.requiredPermissions...)
			require.Equalf(t, test.result, err, "")
		})
	}
}
