// Code generated by mockery. DO NOT EDIT.

package mocks

import (
	echo "github.com/labstack/echo/v4"

	mock "github.com/stretchr/testify/mock"
)

// FranchisesGroup is an autogenerated mock type for the FranchisesGroup type
type FranchisesGroup struct {
	mock.Mock
}

// Resource provides a mock function with given fields: g
func (_m *FranchisesGroup) Resource(g *echo.Group) {
	_m.Called(g)
}

type mockConstructorTestingTNewFranchisesGroup interface {
	mock.TestingT
	Cleanup(func())
}

// NewFranchisesGroup creates a new instance of FranchisesGroup. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewFranchisesGroup(t mockConstructorTestingTNewFranchisesGroup) *FranchisesGroup {
	mock := &FranchisesGroup{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}