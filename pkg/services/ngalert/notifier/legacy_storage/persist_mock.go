// Code generated by mockery v2.44.1. DO NOT EDIT.

package legacy_storage

import (
	context "context"

	models "github.com/grafana/grafana/pkg/services/ngalert/models"
	mock "github.com/stretchr/testify/mock"
)

// MockAMConfigStore is an autogenerated mock type for the AMConfigStore type
type MockAMConfigStore struct {
	mock.Mock
}

type MockAMConfigStore_Expecter struct {
	mock *mock.Mock
}

func (_m *MockAMConfigStore) EXPECT() *MockAMConfigStore_Expecter {
	return &MockAMConfigStore_Expecter{mock: &_m.Mock}
}

// GetLatestAlertmanagerConfiguration provides a mock function with given fields: ctx, orgID
func (_m *MockAMConfigStore) GetLatestAlertmanagerConfiguration(ctx context.Context, orgID int64) (*models.AlertConfiguration, error) {
	ret := _m.Called(ctx, orgID)

	if len(ret) == 0 {
		panic("no return value specified for GetLatestAlertmanagerConfiguration")
	}

	var r0 *models.AlertConfiguration
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int64) (*models.AlertConfiguration, error)); ok {
		return rf(ctx, orgID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int64) *models.AlertConfiguration); ok {
		r0 = rf(ctx, orgID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.AlertConfiguration)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, int64) error); ok {
		r1 = rf(ctx, orgID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockAMConfigStore_GetLatestAlertmanagerConfiguration_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetLatestAlertmanagerConfiguration'
type MockAMConfigStore_GetLatestAlertmanagerConfiguration_Call struct {
	*mock.Call
}

// GetLatestAlertmanagerConfiguration is a helper method to define mock.On call
//   - ctx context.Context
//   - orgID int64
func (_e *MockAMConfigStore_Expecter) GetLatestAlertmanagerConfiguration(ctx interface{}, orgID interface{}) *MockAMConfigStore_GetLatestAlertmanagerConfiguration_Call {
	return &MockAMConfigStore_GetLatestAlertmanagerConfiguration_Call{Call: _e.mock.On("GetLatestAlertmanagerConfiguration", ctx, orgID)}
}

func (_c *MockAMConfigStore_GetLatestAlertmanagerConfiguration_Call) Run(run func(ctx context.Context, orgID int64)) *MockAMConfigStore_GetLatestAlertmanagerConfiguration_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(int64))
	})
	return _c
}

func (_c *MockAMConfigStore_GetLatestAlertmanagerConfiguration_Call) Return(_a0 *models.AlertConfiguration, _a1 error) *MockAMConfigStore_GetLatestAlertmanagerConfiguration_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockAMConfigStore_GetLatestAlertmanagerConfiguration_Call) RunAndReturn(run func(context.Context, int64) (*models.AlertConfiguration, error)) *MockAMConfigStore_GetLatestAlertmanagerConfiguration_Call {
	_c.Call.Return(run)
	return _c
}

// UpdateAlertmanagerConfiguration provides a mock function with given fields: ctx, cmd
func (_m *MockAMConfigStore) UpdateAlertmanagerConfiguration(ctx context.Context, cmd *models.SaveAlertmanagerConfigurationCmd) error {
	ret := _m.Called(ctx, cmd)

	if len(ret) == 0 {
		panic("no return value specified for UpdateAlertmanagerConfiguration")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *models.SaveAlertmanagerConfigurationCmd) error); ok {
		r0 = rf(ctx, cmd)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockAMConfigStore_UpdateAlertmanagerConfiguration_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UpdateAlertmanagerConfiguration'
type MockAMConfigStore_UpdateAlertmanagerConfiguration_Call struct {
	*mock.Call
}

// UpdateAlertmanagerConfiguration is a helper method to define mock.On call
//   - ctx context.Context
//   - cmd *models.SaveAlertmanagerConfigurationCmd
func (_e *MockAMConfigStore_Expecter) UpdateAlertmanagerConfiguration(ctx interface{}, cmd interface{}) *MockAMConfigStore_UpdateAlertmanagerConfiguration_Call {
	return &MockAMConfigStore_UpdateAlertmanagerConfiguration_Call{Call: _e.mock.On("UpdateAlertmanagerConfiguration", ctx, cmd)}
}

func (_c *MockAMConfigStore_UpdateAlertmanagerConfiguration_Call) Run(run func(ctx context.Context, cmd *models.SaveAlertmanagerConfigurationCmd)) *MockAMConfigStore_UpdateAlertmanagerConfiguration_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*models.SaveAlertmanagerConfigurationCmd))
	})
	return _c
}

func (_c *MockAMConfigStore_UpdateAlertmanagerConfiguration_Call) Return(_a0 error) *MockAMConfigStore_UpdateAlertmanagerConfiguration_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockAMConfigStore_UpdateAlertmanagerConfiguration_Call) RunAndReturn(run func(context.Context, *models.SaveAlertmanagerConfigurationCmd) error) *MockAMConfigStore_UpdateAlertmanagerConfiguration_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockAMConfigStore creates a new instance of MockAMConfigStore. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockAMConfigStore(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockAMConfigStore {
	mock := &MockAMConfigStore{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
