package role_system

import (
	"github.com/stretchr/testify/require"
	"open-chat/internal/entities"
	"testing"
)

func TestRoleSystem_Check(t *testing.T) {
	tests := []struct {
		role                []entities.Role
		requiredPermissions []entities.Permission
		result              error
	}{
		{[]entities.Role{entities.NewRole(PERM_ADD_FILE)}, []entities.Permission{}, nil},
		{[]entities.Role{entities.NewRole(PERM_ADD_FILE)}, []entities.Permission{PERM_ADD_FILE}, nil},
		{[]entities.Role{entities.NewRole(PERM_DELETE_MESSAGE, PERM_DELETE_CHANNEL)}, []entities.Permission{PERM_DELETE_MESSAGE, PERM_DELETE_CHANNEL}, nil},
		{[]entities.Role{entities.NewRole(PERM_DELETE_MESSAGE, PERM_INVITE_USER), entities.NewRole(PERM_ADD_STICKER, PERM_CHANGE_ROLE)}, []entities.Permission{PERM_CHANGE_ROLE}, nil},
		{[]entities.Role{entities.NewRole(), entities.NewRole(PERM_DELETE_MESSAGE, PERM_INVITE_USER)}, []entities.Permission{PERM_ADD_FILE}, ErrNotEnoughPermissions},
		{[]entities.Role{entities.NewRole(PERM_DELETE_MESSAGE), entities.NewRole(PERM_DELETE_MESSAGE, PERM_INVITE_USER)}, []entities.Permission{PERM_DELETE_MESSAGE, PERM_DELETE_CHANNEL}, ErrNotEnoughPermissions},
	}

	for _, test := range tests {
		t.Run("", func(t *testing.T) {
			err := NewRoleSystem(test.role).Check(test.requiredPermissions...)
			require.Equalf(t, test.result, err, "")
		})
	}
}
