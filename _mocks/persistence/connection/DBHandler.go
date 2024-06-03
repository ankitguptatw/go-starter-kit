// Code generated by mockery v2.43.2. DO NOT EDIT.

package connection

import (
	mock "github.com/stretchr/testify/mock"
	gorm "gorm.io/gorm"
)

// DBHandler is an autogenerated mock type for the DBHandler type
type DBHandler struct {
	mock.Mock
}

// GetDB provides a mock function with given fields:
func (_m *DBHandler) GetDB() (*gorm.DB, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetDB")
	}

	var r0 *gorm.DB
	var r1 error
	if rf, ok := ret.Get(0).(func() (*gorm.DB, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() *gorm.DB); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*gorm.DB)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewDBHandler creates a new instance of DBHandler. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewDBHandler(t interface {
	mock.TestingT
	Cleanup(func())
}) *DBHandler {
	mock := &DBHandler{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}