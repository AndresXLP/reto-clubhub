// Code generated by mockery. DO NOT EDIT.

package mocks

import (
	echo "github.com/labstack/echo/v4"

	mock "github.com/stretchr/testify/mock"
)

// CompaniesGroup is an autogenerated mock type for the CompaniesGroup type
type CompaniesGroup struct {
	mock.Mock
}

// Resource provides a mock function with given fields: g
func (_m *CompaniesGroup) Resource(g *echo.Group) {
	_m.Called(g)
}

type mockConstructorTestingTNewCompaniesGroup interface {
	mock.TestingT
	Cleanup(func())
}

// NewCompaniesGroup creates a new instance of CompaniesGroup. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewCompaniesGroup(t mockConstructorTestingTNewCompaniesGroup) *CompaniesGroup {
	mock := &CompaniesGroup{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
