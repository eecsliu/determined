// Code generated by mockery v2.20.0. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"

	model "github.com/determined-ai/determined/master/pkg/model"

	projectv1 "github.com/determined-ai/determined/proto/pkg/projectv1"

	workspacev1 "github.com/determined-ai/determined/proto/pkg/workspacev1"
)

// WorkspaceAuthZ is an autogenerated mock type for the WorkspaceAuthZ type
type WorkspaceAuthZ struct {
	mock.Mock
}

// CanArchiveWorkspace provides a mock function with given fields: ctx, curUser, _a2
func (_m *WorkspaceAuthZ) CanArchiveWorkspace(ctx context.Context, curUser model.User, _a2 *workspacev1.Workspace) error {
	ret := _m.Called(ctx, curUser, _a2)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, model.User, *workspacev1.Workspace) error); ok {
		r0 = rf(ctx, curUser, _a2)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CanBindRPWorkspace provides a mock function with given fields: ctx, curUser, workspaceIDs
func (_m *WorkspaceAuthZ) CanBindRPWorkspace(ctx context.Context, curUser model.User, workspaceIDs []int32) error {
	ret := _m.Called(ctx, curUser, workspaceIDs)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, model.User, []int32) error); ok {
		r0 = rf(ctx, curUser, workspaceIDs)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CanCreateWorkspace provides a mock function with given fields: ctx, curUser
func (_m *WorkspaceAuthZ) CanCreateWorkspace(ctx context.Context, curUser model.User) error {
	ret := _m.Called(ctx, curUser)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, model.User) error); ok {
		r0 = rf(ctx, curUser)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CanCreateWorkspaceWithAgentUserGroup provides a mock function with given fields: ctx, curUser
func (_m *WorkspaceAuthZ) CanCreateWorkspaceWithAgentUserGroup(ctx context.Context, curUser model.User) error {
	ret := _m.Called(ctx, curUser)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, model.User) error); ok {
		r0 = rf(ctx, curUser)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CanCreateWorkspaceWithCheckpointStorageConfig provides a mock function with given fields: ctx, curUser
func (_m *WorkspaceAuthZ) CanCreateWorkspaceWithCheckpointStorageConfig(ctx context.Context, curUser model.User) error {
	ret := _m.Called(ctx, curUser)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, model.User) error); ok {
		r0 = rf(ctx, curUser)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CanDeleteWorkspace provides a mock function with given fields: ctx, curUser, _a2
func (_m *WorkspaceAuthZ) CanDeleteWorkspace(ctx context.Context, curUser model.User, _a2 *workspacev1.Workspace) error {
	ret := _m.Called(ctx, curUser, _a2)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, model.User, *workspacev1.Workspace) error); ok {
		r0 = rf(ctx, curUser, _a2)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CanGetWorkspace provides a mock function with given fields: ctx, curUser, _a2
func (_m *WorkspaceAuthZ) CanGetWorkspace(ctx context.Context, curUser model.User, _a2 *workspacev1.Workspace) error {
	ret := _m.Called(ctx, curUser, _a2)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, model.User, *workspacev1.Workspace) error); ok {
		r0 = rf(ctx, curUser, _a2)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CanPinWorkspace provides a mock function with given fields: ctx, curUser, _a2
func (_m *WorkspaceAuthZ) CanPinWorkspace(ctx context.Context, curUser model.User, _a2 *workspacev1.Workspace) error {
	ret := _m.Called(ctx, curUser, _a2)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, model.User, *workspacev1.Workspace) error); ok {
		r0 = rf(ctx, curUser, _a2)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CanSetWorkspacesAgentUserGroup provides a mock function with given fields: ctx, curUser, _a2
func (_m *WorkspaceAuthZ) CanSetWorkspacesAgentUserGroup(ctx context.Context, curUser model.User, _a2 *workspacev1.Workspace) error {
	ret := _m.Called(ctx, curUser, _a2)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, model.User, *workspacev1.Workspace) error); ok {
		r0 = rf(ctx, curUser, _a2)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CanSetWorkspacesCheckpointStorageConfig provides a mock function with given fields: ctx, curUser, _a2
func (_m *WorkspaceAuthZ) CanSetWorkspacesCheckpointStorageConfig(ctx context.Context, curUser model.User, _a2 *workspacev1.Workspace) error {
	ret := _m.Called(ctx, curUser, _a2)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, model.User, *workspacev1.Workspace) error); ok {
		r0 = rf(ctx, curUser, _a2)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CanSetWorkspacesName provides a mock function with given fields: ctx, curUser, _a2
func (_m *WorkspaceAuthZ) CanSetWorkspacesName(ctx context.Context, curUser model.User, _a2 *workspacev1.Workspace) error {
	ret := _m.Called(ctx, curUser, _a2)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, model.User, *workspacev1.Workspace) error); ok {
		r0 = rf(ctx, curUser, _a2)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CanUnBindRPWorkspace provides a mock function with given fields: ctx, curUser, workspaceIDs
func (_m *WorkspaceAuthZ) CanUnBindRPWorkspace(ctx context.Context, curUser model.User, workspaceIDs []int32) error {
	ret := _m.Called(ctx, curUser, workspaceIDs)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, model.User, []int32) error); ok {
		r0 = rf(ctx, curUser, workspaceIDs)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CanUnarchiveWorkspace provides a mock function with given fields: ctx, curUser, _a2
func (_m *WorkspaceAuthZ) CanUnarchiveWorkspace(ctx context.Context, curUser model.User, _a2 *workspacev1.Workspace) error {
	ret := _m.Called(ctx, curUser, _a2)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, model.User, *workspacev1.Workspace) error); ok {
		r0 = rf(ctx, curUser, _a2)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CanUnpinWorkspace provides a mock function with given fields: ctx, curUser, _a2
func (_m *WorkspaceAuthZ) CanUnpinWorkspace(ctx context.Context, curUser model.User, _a2 *workspacev1.Workspace) error {
	ret := _m.Called(ctx, curUser, _a2)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, model.User, *workspacev1.Workspace) error); ok {
		r0 = rf(ctx, curUser, _a2)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// FilterWorkspaceIDs provides a mock function with given fields: ctx, curUser, workspaces
func (_m *WorkspaceAuthZ) FilterWorkspaceIDs(ctx context.Context, curUser model.User, workspaces []int32) ([]int32, error) {
	ret := _m.Called(ctx, curUser, workspaces)

	var r0 []int32
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, model.User, []int32) ([]int32, error)); ok {
		return rf(ctx, curUser, workspaces)
	}
	if rf, ok := ret.Get(0).(func(context.Context, model.User, []int32) []int32); ok {
		r0 = rf(ctx, curUser, workspaces)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]int32)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, model.User, []int32) error); ok {
		r1 = rf(ctx, curUser, workspaces)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FilterWorkspaceProjects provides a mock function with given fields: ctx, curUser, projects
func (_m *WorkspaceAuthZ) FilterWorkspaceProjects(ctx context.Context, curUser model.User, projects []*projectv1.Project) ([]*projectv1.Project, error) {
	ret := _m.Called(ctx, curUser, projects)

	var r0 []*projectv1.Project
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, model.User, []*projectv1.Project) ([]*projectv1.Project, error)); ok {
		return rf(ctx, curUser, projects)
	}
	if rf, ok := ret.Get(0).(func(context.Context, model.User, []*projectv1.Project) []*projectv1.Project); ok {
		r0 = rf(ctx, curUser, projects)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*projectv1.Project)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, model.User, []*projectv1.Project) error); ok {
		r1 = rf(ctx, curUser, projects)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FilterWorkspaces provides a mock function with given fields: ctx, curUser, workspaces
func (_m *WorkspaceAuthZ) FilterWorkspaces(ctx context.Context, curUser model.User, workspaces []*workspacev1.Workspace) ([]*workspacev1.Workspace, error) {
	ret := _m.Called(ctx, curUser, workspaces)

	var r0 []*workspacev1.Workspace
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, model.User, []*workspacev1.Workspace) ([]*workspacev1.Workspace, error)); ok {
		return rf(ctx, curUser, workspaces)
	}
	if rf, ok := ret.Get(0).(func(context.Context, model.User, []*workspacev1.Workspace) []*workspacev1.Workspace); ok {
		r0 = rf(ctx, curUser, workspaces)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*workspacev1.Workspace)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, model.User, []*workspacev1.Workspace) error); ok {
		r1 = rf(ctx, curUser, workspaces)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewWorkspaceAuthZ interface {
	mock.TestingT
	Cleanup(func())
}

// NewWorkspaceAuthZ creates a new instance of WorkspaceAuthZ. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewWorkspaceAuthZ(t mockConstructorTestingTNewWorkspaceAuthZ) *WorkspaceAuthZ {
	mock := &WorkspaceAuthZ{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
