// Code generated by mockery. DO NOT EDIT.

package mocks

import (
	context "context"
	http "net/http"

	mock "github.com/stretchr/testify/mock"
)

// HttpClient is an autogenerated mock type for the HttpClient type
type HttpClient struct {
	mock.Mock
}

// Get provides a mock function with given fields: c, url
func (_m *HttpClient) Get(c context.Context, url string) (*http.Response, error) {
	ret := _m.Called(c, url)

	var r0 *http.Response
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*http.Response, error)); ok {
		return rf(c, url)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *http.Response); ok {
		r0 = rf(c, url)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*http.Response)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(c, url)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewHttpClient interface {
	mock.TestingT
	Cleanup(func())
}

// NewHttpClient creates a new instance of HttpClient. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewHttpClient(t mockConstructorTestingTNewHttpClient) *HttpClient {
	mock := &HttpClient{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}