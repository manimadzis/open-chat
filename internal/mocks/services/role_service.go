// Code generated by mockery v2.23.1. DO NOT EDIT.

package mocks

import (
	context "context"
	entities "open-chat/internal/entities"

	mock "github.com/stretchr/testify/mock"
)

// RoleService is an autogenerated mock type for the RoleService type
type RoleService struct {
	mock.Mock
}

// Change provides a mock function with given fields: ctx, role, userId, serverId
func (_m *RoleService) Change(ctx context.Context, role entities.Role, userId entities.UserId, serverId entities.ServerId) error {
	ret := _m.Called(ctx, role, userId, serverId)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, entities.Role, entities.UserId, entities.ServerId) error); ok {
		r0 = rf(ctx, role, userId, serverId)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Create provides a mock function with given fields: ctx, role
func (_m *RoleService) Create(ctx context.Context, role entities.Role) (entities.RoleId, error) {
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

// Delete provides a mock function with given fields: ctx, roleId, userId, serverId
func (_m *RoleService) Delete(ctx context.Context, roleId entities.RoleId, userId entities.UserId, serverId entities.ServerId) error {
	ret := _m.Called(ctx, roleId, userId, serverId)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, entities.RoleId, entities.UserId, entities.ServerId) error); ok {
		r0 = rf(ctx, roleId, userId, serverId)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// FindByServer provides a mock function with given fields: ctx, serverId, userId
func (_m *RoleService) FindByServer(ctx context.Context, serverId entities.ServerId, userId entities.UserId) ([]entities.Role, error) {
	ret := _m.Called(ctx, serverId, userId)

	var r0 []entities.Role
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, entities.ServerId, entities.UserId) ([]entities.Role, error)); ok {
		return rf(ctx, serverId, userId)
	}
	if rf, ok := ret.Get(0).(func(context.Context, entities.ServerId, entities.UserId) []entities.Role); ok {
		r0 = rf(ctx, serverId, userId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]entities.Role)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, entities.ServerId, entities.UserId) error); ok {
		r1 = rf(ctx, serverId, userId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewRoleService interface {
	mock.TestingT
	Cleanup(func())
}

// NewRoleService creates a new instance of RoleService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewRoleService(t mockConstructorTestingTNewRoleService) *RoleService {
	mock := &RoleService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
