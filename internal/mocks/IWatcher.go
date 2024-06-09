// Code generated by mockery v2.40.3. DO NOT EDIT.

package mocks

import (
	core "github-observer/internal/core"

	mock "github.com/stretchr/testify/mock"
)

// IWatcher is an autogenerated mock type for the IWatcher type
type IWatcher struct {
	mock.Mock
}

// CheckRateLimit provides a mock function with given fields:
func (_m *IWatcher) CheckRateLimit() {
	_m.Called()
}

// PullRequests provides a mock function with given fields: _a0
func (_m *IWatcher) PullRequests(_a0 core.Repository) {
	_m.Called(_a0)
}

// Watch provides a mock function with given fields:
func (_m *IWatcher) Watch() {
	_m.Called()
}

// WorkflowRuns provides a mock function with given fields: _a0
func (_m *IWatcher) WorkflowRuns(_a0 core.Repository) {
	_m.Called(_a0)
}

// NewIWatcher creates a new instance of IWatcher. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewIWatcher(t interface {
	mock.TestingT
	Cleanup(func())
}) *IWatcher {
	mock := &IWatcher{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
