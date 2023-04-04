// Code generated by mockery v2.23.1. DO NOT EDIT.

package mocks

import (
	context "context"
	entities "open-chat/internal/entities"

	mock "github.com/stretchr/testify/mock"
)

// ServerProfileChecker is an autogenerated mock type for the ServerProfileChecker type
type ServerProfileChecker struct {
	mock.Mock
}

// Check provides a mock function with given fields: ctx, userId, serverId, permissions
func (_m *ServerProfileChecker) Check(ctx context.Context, userId entities.UserId, serverId entities.ServerId, permissions ...entities.PermissionValue) error {
	_va := make([]interface{}, len(permissions))
	for _i := range permissions {
		_va[_i] = permissions[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, userId, serverId)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, entities.UserId, entities.ServerId, ...entities.PermissionValue) error); ok {
		r0 = rf(ctx, userId, serverId, permissions...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewServerProfileChecker interface {
	mock.TestingT
	Cleanup(func())
}

// NewServerProfileChecker creates a new instance of ServerProfileChecker. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewServerProfileChecker(t mockConstructorTestingTNewServerProfileChecker) *ServerProfileChecker {
	mock := &ServerProfileChecker{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
