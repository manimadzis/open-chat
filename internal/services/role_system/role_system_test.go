package role_system

import (
	"github.com/stretchr/testify/require"
	"open-chat/internal/entities"
	"open-chat/internal/services"
	"testing"
)

func TestRoleSystem_Check(t *testing.T) {
	tests := []struct {
		name                string
		role                []entities.Role
		requiredPermissions []entities.PermissionValue
		result              error
	}{
		{
			"no required permissions",
			[]entities.Role{entities.NewRoleByPermissionValues(PERM_ADD_FILE)},
			[]entities.PermissionValue{},
			nil,
		},
		{
			"all permission satisfied",
			[]entities.Role{entities.NewRoleByPermissionValues(PERM_ADD_FILE)},
			[]entities.PermissionValue{PERM_ADD_FILE},
			nil,
		},
		{
			"all permission satisfied",
			[]entities.Role{
				entities.NewRoleByPermissionValues(
					PERM_DELETE_MESSAGE,
					PERM_DELETE_CHANNEL,
				),
			},
			[]entities.PermissionValue{
				PERM_DELETE_MESSAGE,
				PERM_DELETE_CHANNEL,
			},
			nil,
		},
		{
			"all permission satisfied",
			[]entities.Role{
				entities.NewRoleByPermissionValues(
					PERM_DELETE_MESSAGE,
					PERM_INVITE_USER,
				),
				entities.NewRoleByPermissionValues(
					PERM_ADD_STICKER,
					PERM_CHANGE_ROLE,
				),
			},
			[]entities.PermissionValue{PERM_CHANGE_ROLE},
			nil,
		},
		{
			"not all permission satisfied",
			[]entities.Role{
				entities.NewRoleByPermissionValues(),
				entities.NewRoleByPermissionValues(
					PERM_DELETE_MESSAGE,
					PERM_INVITE_USER,
				),
			},
			[]entities.PermissionValue{PERM_ADD_FILE},
			services.ErrNotEnoughPermissions,
		},
		{
			"not all permission satisfied",
			[]entities.Role{
				entities.NewRoleByPermissionValues(PERM_DELETE_MESSAGE),
				entities.NewRoleByPermissionValues(
					PERM_DELETE_MESSAGE,
					PERM_INVITE_USER,
				),
			},
			[]entities.PermissionValue{
				PERM_DELETE_MESSAGE,
				PERM_DELETE_CHANNEL,
			},
			services.ErrNotEnoughPermissions,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			err := NewRoleSystem(test.role).Check(test.requiredPermissions...)
			require.Equalf(t, test.result, err, "")
		},
		)
	}
}
