// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import (
	context "context"

	users "cinemo.com/shoping-cart/internal/users"
	mock "github.com/stretchr/testify/mock"
)

// Service is an autogenerated mock type for the Service type
type Service struct {
	mock.Mock
}

// CreateUser provides a mock function with given fields: ctx, username, password, firstName, lastName
func (_m *Service) CreateUser(ctx context.Context, username string, password string, firstName *string, lastName *string) (*users.User, error) {
	ret := _m.Called(ctx, username, password, firstName, lastName)

	var r0 *users.User
	if rf, ok := ret.Get(0).(func(context.Context, string, string, *string, *string) *users.User); ok {
		r0 = rf(ctx, username, password, firstName, lastName)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*users.User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, string, *string, *string) error); ok {
		r1 = rf(ctx, username, password, firstName, lastName)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RetrieveUserByUsername provides a mock function with given fields: ctx, username
func (_m *Service) RetrieveUserByUsername(ctx context.Context, username string) (*users.User, error) {
	ret := _m.Called(ctx, username)

	var r0 *users.User
	if rf, ok := ret.Get(0).(func(context.Context, string) *users.User); ok {
		r0 = rf(ctx, username)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*users.User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, username)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Validate provides a mock function with given fields: ctx, username, password
func (_m *Service) Validate(ctx context.Context, username string, password string) (*users.User, error) {
	ret := _m.Called(ctx, username, password)

	var r0 *users.User
	if rf, ok := ret.Get(0).(func(context.Context, string, string) *users.User); ok {
		r0 = rf(ctx, username, password)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*users.User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, string) error); ok {
		r1 = rf(ctx, username, password)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
