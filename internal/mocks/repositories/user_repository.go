// Code generated by mockery v2.23.1. DO NOT EDIT.

package mocks

import (
	context "context"
	entities "open-chat/internal/entities"

	mock "github.com/stretchr/testify/mock"
)

// UserRepository is an autogenerated mock type for the UserRepository type
type UserRepository struct {
	mock.Mock
}

// Create provides a mock function with given fields: ctx, user
func (_m *UserRepository) Create(ctx context.Context, user entities.User) (entities.UserId, error) {
	ret := _m.Called(ctx, user)

	var r0 entities.UserId
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, entities.User) (entities.UserId, error)); ok {
		return rf(ctx, user)
	}
	if rf, ok := ret.Get(0).(func(context.Context, entities.User) entities.UserId); ok {
		r0 = rf(ctx, user)
	} else {
		r0 = ret.Get(0).(entities.UserId)
	}

	if rf, ok := ret.Get(1).(func(context.Context, entities.User) error); ok {
		r1 = rf(ctx, user)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindById provides a mock function with given fields: ctx, userId
func (_m *UserRepository) FindById(ctx context.Context, userId entities.UserId) (*entities.User, error) {
	ret := _m.Called(ctx, userId)

	var r0 *entities.User
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, entities.UserId) (*entities.User, error)); ok {
		return rf(ctx, userId)
	}
	if rf, ok := ret.Get(0).(func(context.Context, entities.UserId) *entities.User); ok {
		r0 = rf(ctx, userId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entities.User)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, entities.UserId) error); ok {
		r1 = rf(ctx, userId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindByLogin provides a mock function with given fields: ctx, login
func (_m *UserRepository) FindByLogin(ctx context.Context, login string) (*entities.User, error) {
	ret := _m.Called(ctx, login)

	var r0 *entities.User
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*entities.User, error)); ok {
		return rf(ctx, login)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *entities.User); ok {
		r0 = rf(ctx, login)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entities.User)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, login)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewUserRepository interface {
	mock.TestingT
	Cleanup(func())
}

// NewUserRepository creates a new instance of UserRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewUserRepository(t mockConstructorTestingTNewUserRepository) *UserRepository {
	mock := &UserRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
