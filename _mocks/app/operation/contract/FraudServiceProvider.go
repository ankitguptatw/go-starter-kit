// Code generated by mockery v2.43.2. DO NOT EDIT.

package contract

import (
	context "context"

	mock "github.com/stretchr/testify/mock"

	serviceprovider "myapp/app/serviceprovider"
)

// FraudServiceProvider is an autogenerated mock type for the FraudServiceProvider type
type FraudServiceProvider struct {
	mock.Mock
}

// IsFraud provides a mock function with given fields: ctx, request
func (_m *FraudServiceProvider) IsFraud(ctx context.Context, request serviceprovider.FraudClientRequest) (bool, error) {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for IsFraud")
	}

	var r0 bool
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, serviceprovider.FraudClientRequest) (bool, error)); ok {
		return rf(ctx, request)
	}
	if rf, ok := ret.Get(0).(func(context.Context, serviceprovider.FraudClientRequest) bool); ok {
		r0 = rf(ctx, request)
	} else {
		r0 = ret.Get(0).(bool)
	}

	if rf, ok := ret.Get(1).(func(context.Context, serviceprovider.FraudClientRequest) error); ok {
		r1 = rf(ctx, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewFraudServiceProvider creates a new instance of FraudServiceProvider. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewFraudServiceProvider(t interface {
	mock.TestingT
	Cleanup(func())
}) *FraudServiceProvider {
	mock := &FraudServiceProvider{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
