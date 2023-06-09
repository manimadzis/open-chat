// Code generated by mockery v2.23.1. DO NOT EDIT.

package mocks

import (
	entities "open-chat/internal/entities"

	mock "github.com/stretchr/testify/mock"
)

// RoleSystem is an autogenerated mock type for the RoleSystem type
type RoleSystem struct {
	mock.Mock
}

// Check provides a mock function with given fields: permission
func (_m *RoleSystem) Check(permission ...entities.PermissionValue) error {
	_va := make([]interface{}, len(permission))
	for _i := range permission {
		_va[_i] = permission[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 error
	if rf, ok := ret.Get(0).(func(...entities.PermissionValue) error); ok {
		r0 = rf(permission...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SetRoles provides a mock function with given fields: _a0
func (_m *RoleSystem) SetRoles(_a0 []entities.Role) {
	_m.Called(_a0)
}

type mockConstructorTestingTNewRoleSystem interface {
	mock.TestingT
	Cleanup(func())
}

// NewRoleSystem creates a new instance of RoleSystem. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewRoleSystem(t mockConstructorTestingTNewRoleSystem) *RoleSystem {
	mock := &RoleSystem{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
