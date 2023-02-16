package mocks

import (
	"context"

	"github.com/Piyawat-T/go-centralize-configuration/domain"
	"github.com/stretchr/testify/mock"
)

type PropertiesUsecase struct {
	mock.Mock
}

func (u *PropertiesUsecase) GetProperties(c context.Context) ([]domain.Properties, error) {
	ret := u.Called(c)

	var r0 []domain.Properties
	if rf, ok := ret.Get(0).(func(context.Context) []domain.Properties); ok {
		r0 = rf(c)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.Properties)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(c)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

func (u *PropertiesUsecase) GetByApplicationAndProfile(c context.Context, application string, profile string) ([]domain.Properties, error) {
	ret := u.Called(c, application, profile)

	var r0 []domain.Properties
	if rf, ok := ret.Get(0).(func(context.Context, string, string) []domain.Properties); ok {
		r0 = rf(c, application, profile)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.Properties)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, string) error); ok {
		r1 = rf(c, application, profile)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
