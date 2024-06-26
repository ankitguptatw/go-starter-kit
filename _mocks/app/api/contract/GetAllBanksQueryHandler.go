// Code generated by mockery v2.43.2. DO NOT EDIT.

package contract

import (
	context "context"
	operation "myapp/app/operation"

	mock "github.com/stretchr/testify/mock"
)

// GetAllBanksQueryHandler is an autogenerated mock type for the GetAllBanksQueryHandler type
type GetAllBanksQueryHandler struct {
	mock.Mock
}

// Handle provides a mock function with given fields: ctx
func (_m *GetAllBanksQueryHandler) Handle(ctx context.Context) (operation.BanksResponse, error) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for Handle")
	}

	var r0 operation.BanksResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) (operation.BanksResponse, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) operation.BanksResponse); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Get(0).(operation.BanksResponse)
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewGetAllBanksQueryHandler creates a new instance of GetAllBanksQueryHandler. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewGetAllBanksQueryHandler(t interface {
	mock.TestingT
	Cleanup(func())
}) *GetAllBanksQueryHandler {
	mock := &GetAllBanksQueryHandler{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
