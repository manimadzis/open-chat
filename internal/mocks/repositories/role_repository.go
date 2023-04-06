// Code generated by mockery v2.23.1. DO NOT EDIT.

package mocks

import (
	context "context"
	entities "open-chat/internal/entities"

	mock "github.com/stretchr/testify/mock"
)

// RoleRepository is an autogenerated mock type for the RoleRepository type
type RoleRepository struct {
	mock.Mock
}

// Change provides a mock function with given fields: ctx, role
func (_m *RoleRepository) Change(ctx context.Context, role entities.Role) error {
	ret := _m.Called(ctx, role)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, entities.Role) error); ok {
		r0 = rf(ctx, role)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Create provides a mock function with given fields: ctx, role
func (_m *RoleRepository) Create(ctx context.Context, role entities.Role) (entities.RoleId, error) {
	ret := _m.Called(ctx, role)

	var r0 entities.RoleId
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, entities.Role) (entities.RoleId, error)); ok {
		return rf(ctx, role)
	}
	if rf, ok := ret.Get(0).(func(context.Context, entities.Role) entities.RoleId); ok {
		r0 = rf(ctx, role)
	} else {
		r0 = ret.Get(0).(entities.RoleId)
	}

	if rf, ok := ret.Get(1).(func(context.Context, entities.Role) error); ok {
		r1 = rf(ctx, role)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: ctx, roleId
func (_m *RoleRepository) Delete(ctx context.Context, roleId entities.RoleId) error {
	ret := _m.Called(ctx, roleId)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, entities.RoleId) error); ok {
		r0 = rf(ctx, roleId)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// FindPermissionsByValue provides a mock function with given fields: ctx, permission
func (_m *RoleRepository) FindPermissionsByValue(ctx context.Context, permission []entities.PermissionValue) ([]entities.Permission, error) {
	ret := _m.Called(ctx, permission)

	var r0 []entities.Permission
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, []entities.PermissionValue) ([]entities.Permission, error)); ok {
		return rf(ctx, permission)
	}
	if rf, ok := ret.Get(0).(func(context.Context, []entities.PermissionValue) []entities.Permission); ok {
		r0 = rf(ctx, permission)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]entities.Permission)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, []entities.PermissionValue) error); ok {
		r1 = rf(ctx, permission)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindRolesByServerId provides a mock function with given fields: ctx, serverId
func (_m *RoleRepository) FindRolesByServerId(ctx context.Context, serverId entities.ServerId) ([]entities.Role, error) {
	ret := _m.Called(ctx, serverId)

	var r0 []entities.Role
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, entities.ServerId) ([]entities.Role, error)); ok {
		return rf(ctx, serverId)
	}
	if rf, ok := ret.Get(0).(func(context.Context, entities.ServerId) []entities.Role); ok {
		r0 = rf(ctx, serverId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]entities.Role)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, entities.ServerId) error); ok {
		r1 = rf(ctx, serverId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewRoleRepository interface {
	mock.TestingT
	Cleanup(func())
}

// NewRoleRepository creates a new instance of RoleRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewRoleRepository(t mockConstructorTestingTNewRoleRepository) *RoleRepository {
	mock := &RoleRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
