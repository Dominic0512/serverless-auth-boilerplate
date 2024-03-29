// Code generated by mockery v2.41.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// PasswordHelper is an autogenerated mock type for the PasswordHelper type
type PasswordHelper struct {
	mock.Mock
}

// Hash provides a mock function with given fields: password
func (_m *PasswordHelper) Hash(password string) (string, error) {
	ret := _m.Called(password)

	if len(ret) == 0 {
		panic("no return value specified for Hash")
	}

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (string, error)); ok {
		return rf(password)
	}
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(password)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(password)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// IsMatch provides a mock function with given fields: hashedPassword, password
func (_m *PasswordHelper) IsMatch(hashedPassword string, password string) bool {
	ret := _m.Called(hashedPassword, password)

	if len(ret) == 0 {
		panic("no return value specified for IsMatch")
	}

	var r0 bool
	if rf, ok := ret.Get(0).(func(string, string) bool); ok {
		r0 = rf(hashedPassword, password)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// NewPasswordHelper creates a new instance of PasswordHelper. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewPasswordHelper(t interface {
	mock.TestingT
	Cleanup(func())
}) *PasswordHelper {
	mock := &PasswordHelper{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
