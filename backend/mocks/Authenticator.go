// Code generated by mockery v2.41.0. DO NOT EDIT.

package mocks

import (
	authenticator "github.com/Dominic0512/serverless-auth-boilerplate/infra/authenticator"
	mock "github.com/stretchr/testify/mock"
)

// Authenticator is an autogenerated mock type for the Authenticator type
type Authenticator struct {
	mock.Mock
}

// ExchangeMetaDataByCode provides a mock function with given fields: code
func (_m *Authenticator) ExchangeMetaDataByCode(code string) (*authenticator.AuthMetaData, error) {
	ret := _m.Called(code)

	if len(ret) == 0 {
		panic("no return value specified for ExchangeMetaDataByCode")
	}

	var r0 *authenticator.AuthMetaData
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (*authenticator.AuthMetaData, error)); ok {
		return rf(code)
	}
	if rf, ok := ret.Get(0).(func(string) *authenticator.AuthMetaData); ok {
		r0 = rf(code)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*authenticator.AuthMetaData)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(code)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GenerateAuthCodeURL provides a mock function with given fields:
func (_m *Authenticator) GenerateAuthCodeURL() (string, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GenerateAuthCodeURL")
	}

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func() (string, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewAuthenticator creates a new instance of Authenticator. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewAuthenticator(t interface {
	mock.TestingT
	Cleanup(func())
}) *Authenticator {
	mock := &Authenticator{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
