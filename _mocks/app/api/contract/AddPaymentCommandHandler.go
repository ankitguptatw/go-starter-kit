// Code generated by mockery v2.43.2. DO NOT EDIT.

package contract

import (
	context "context"
	command "myapp/app/operation/command"

	mock "github.com/stretchr/testify/mock"

	operation "myapp/app/operation"
)

// AddPaymentCommandHandler is an autogenerated mock type for the AddPaymentCommandHandler type
type AddPaymentCommandHandler struct {
	mock.Mock
}

// Handle provides a mock function with given fields: _a0, _a1
func (_m *AddPaymentCommandHandler) Handle(_a0 context.Context, _a1 command.AddPaymentCommand) (operation.AddPaymentResponse, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for Handle")
	}

	var r0 operation.AddPaymentResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, command.AddPaymentCommand) (operation.AddPaymentResponse, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, command.AddPaymentCommand) operation.AddPaymentResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Get(0).(operation.AddPaymentResponse)
	}

	if rf, ok := ret.Get(1).(func(context.Context, command.AddPaymentCommand) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewAddPaymentCommandHandler creates a new instance of AddPaymentCommandHandler. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewAddPaymentCommandHandler(t interface {
	mock.TestingT
	Cleanup(func())
}) *AddPaymentCommandHandler {
	mock := &AddPaymentCommandHandler{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
