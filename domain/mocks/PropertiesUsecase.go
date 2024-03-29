// Code generated by mockery v2.23.1. DO NOT EDIT.

package mocks

import (
	context "context"

	domain "github.com/Piyawat-T/go-centralize-configuration/domain"
	mock "github.com/stretchr/testify/mock"
)

// PropertiesUseCase is an autogenerated mock type for the PropertiesUseCase type
type PropertiesUseCase struct {
	mock.Mock
}

// GetByApplicationAndProfile provides a mock function with given fields: c, application, profile
func (_m *PropertiesUseCase) GetByApplicationAndProfile(c context.Context, application string, profile string) ([]domain.Properties, error) {
	ret := _m.Called(c, application, profile)

	var r0 []domain.Properties
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string) ([]domain.Properties, error)); ok {
		return rf(c, application, profile)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, string) []domain.Properties); ok {
		r0 = rf(c, application, profile)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.Properties)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, string) error); ok {
		r1 = rf(c, application, profile)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetProperties provides a mock function with given fields: c
func (_m *PropertiesUseCase) GetProperties(c context.Context) ([]domain.Properties, error) {
	ret := _m.Called(c)

	var r0 []domain.Properties
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) ([]domain.Properties, error)); ok {
		return rf(c)
	}
	if rf, ok := ret.Get(0).(func(context.Context) []domain.Properties); ok {
		r0 = rf(c)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.Properties)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(c)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewPropertiesUseCase interface {
	mock.TestingT
	Cleanup(func())
}

// NewPropertiesUseCase creates a new instance of PropertiesUseCase. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewPropertiesUseCase(t mockConstructorTestingTNewPropertiesUseCase) *PropertiesUseCase {
	mock := &PropertiesUseCase{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
