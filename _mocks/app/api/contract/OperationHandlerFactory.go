// Code generated by mockery v2.43.2. DO NOT EDIT.

package contract

import (
	factory "myapp/app/operation/factory"

	mock "github.com/stretchr/testify/mock"
)

// OperationHandlerFactory is an autogenerated mock type for the OperationHandlerFactory type
type OperationHandlerFactory struct {
	mock.Mock
}

// CommandHandler provides a mock function with given fields: _a0
func (_m *OperationHandlerFactory) CommandHandler(_a0 factory.CommandHandlers) interface{} {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for CommandHandler")
	}

	var r0 interface{}
	if rf, ok := ret.Get(0).(func(factory.CommandHandlers) interface{}); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(interface{})
		}
	}

	return r0
}

// QueryHandler provides a mock function with given fields: _a0
func (_m *OperationHandlerFactory) QueryHandler(_a0 factory.QueryHandlers) interface{} {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for QueryHandler")
	}

	var r0 interface{}
	if rf, ok := ret.Get(0).(func(factory.QueryHandlers) interface{}); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(interface{})
		}
	}

	return r0
}

// NewOperationHandlerFactory creates a new instance of OperationHandlerFactory. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewOperationHandlerFactory(t interface {
	mock.TestingT
	Cleanup(func())
}) *OperationHandlerFactory {
	mock := &OperationHandlerFactory{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
