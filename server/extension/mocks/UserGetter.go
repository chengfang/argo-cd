// Code generated by mockery v2.43.2. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// UserGetter is an autogenerated mock type for the UserGetter type
type UserGetter struct {
	mock.Mock
}

// GetGroups provides a mock function with given fields: ctx
func (_m *UserGetter) GetGroups(ctx context.Context) []string {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for GetGroups")
	}

	var r0 []string
	if rf, ok := ret.Get(0).(func(context.Context) []string); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	return r0
}

// GetUser provides a mock function with given fields: ctx
func (_m *UserGetter) GetUser(ctx context.Context) string {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for GetUser")
	}

	var r0 string
	if rf, ok := ret.Get(0).(func(context.Context) string); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// NewUserGetter creates a new instance of UserGetter. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewUserGetter(t interface {
	mock.TestingT
	Cleanup(func())
}) *UserGetter {
	mock := &UserGetter{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}