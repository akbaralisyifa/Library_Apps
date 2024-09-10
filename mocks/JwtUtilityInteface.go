// Code generated by mockery v2.43.2. DO NOT EDIT.

package mocks

import (
	jwt "github.com/golang-jwt/jwt/v5"
	mock "github.com/stretchr/testify/mock"
)

// JwtUtilityInteface is an autogenerated mock type for the JwtUtilityInteface type
type JwtUtilityInteface struct {
	mock.Mock
}

// DecodToken provides a mock function with given fields: token
func (_m *JwtUtilityInteface) DecodToken(token *jwt.Token) float64 {
	ret := _m.Called(token)

	if len(ret) == 0 {
		panic("no return value specified for DecodToken")
	}

	var r0 float64
	if rf, ok := ret.Get(0).(func(*jwt.Token) float64); ok {
		r0 = rf(token)
	} else {
		r0 = ret.Get(0).(float64)
	}

	return r0
}

// GenerateJwt provides a mock function with given fields: id
func (_m *JwtUtilityInteface) GenerateJwt(id uint) (string, error) {
	ret := _m.Called(id)

	if len(ret) == 0 {
		panic("no return value specified for GenerateJwt")
	}

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(uint) (string, error)); ok {
		return rf(id)
	}
	if rf, ok := ret.Get(0).(func(uint) string); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(uint) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewJwtUtilityInteface creates a new instance of JwtUtilityInteface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewJwtUtilityInteface(t interface {
	mock.TestingT
	Cleanup(func())
}) *JwtUtilityInteface {
	mock := &JwtUtilityInteface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
