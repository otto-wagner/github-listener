// Code generated by mockery v2.40.3. DO NOT EDIT.

package mocks

import (
	core "github-observer/internal/core"

	github "github.com/google/go-github/v61/github"

	mock "github.com/stretchr/testify/mock"
)

// IExecutor is an autogenerated mock type for the IExecutor type
type IExecutor struct {
	mock.Mock
}

// EventPullRequest provides a mock function with given fields: _a0
func (_m *IExecutor) EventPullRequest(_a0 github.PullRequestEvent) {
	_m.Called(_a0)
}

// EventPullRequestReview provides a mock function with given fields: _a0
func (_m *IExecutor) EventPullRequestReview(_a0 github.PullRequestReviewEvent) {
	_m.Called(_a0)
}

// EventRun provides a mock function with given fields: _a0
func (_m *IExecutor) EventRun(_a0 github.CheckRunEvent) {
	_m.Called(_a0)
}

// LastWorkflows provides a mock function with given fields: _a0, _a1
func (_m *IExecutor) LastWorkflows(_a0 core.Repository, _a1 []*github.WorkflowRun) {
	_m.Called(_a0, _a1)
}

// PullRequests provides a mock function with given fields: _a0, _a1
func (_m *IExecutor) PullRequests(_a0 core.Repository, _a1 []*github.PullRequest) {
	_m.Called(_a0, _a1)
}

// NewIExecutor creates a new instance of IExecutor. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewIExecutor(t interface {
	mock.TestingT
	Cleanup(func())
}) *IExecutor {
	mock := &IExecutor{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
