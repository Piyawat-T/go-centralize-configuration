package usecase_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/Piyawat-T/go-centralize-configuration/domain"
	"github.com/Piyawat-T/go-centralize-configuration/domain/mocks"
	"github.com/Piyawat-T/go-centralize-configuration/usecase"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetProperties(t *testing.T) {
	mockPropertiesRepository := new(mocks.PropertiesRepository)

	t.Run("success", func(t *testing.T) {
		mockProperties := []domain.Properties{
			{
				Id: 1,
			},
		}
		mockPropertiesRepository.On("Fetch", mock.Anything).Return(mockProperties, nil).Once()

		u := usecase.NewPropertiesUsecase(mockPropertiesRepository, time.Second*2)
		list, err := u.GetProperties(context.Background())

		assert.NoError(t, err)
		assert.NotNil(t, list)
		assert.Len(t, list, len(mockProperties))

		mockPropertiesRepository.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		mockPropertiesRepository.On("Fetch", mock.Anything).Return(nil, errors.New("Unexpected")).Once()

		u := usecase.NewPropertiesUsecase(mockPropertiesRepository, time.Second*2)
		list, err := u.GetProperties(context.Background())

		assert.Error(t, err)
		assert.Nil(t, list)

		mockPropertiesRepository.AssertExpectations(t)
	})
}

func TestGetByApplicationAndProfile(t *testing.T) {
	mockPropertiesRepository := new(mocks.PropertiesRepository)
	t.Run("success", func(t *testing.T) {
		mockProperties := []domain.Properties{
			{
				Id: 1,
			},
		}
		mockPropertiesRepository.On("FetchByApplicationAndProfile", mock.Anything, mock.Anything, mock.Anything).Return(mockProperties, nil).Once()

		u := usecase.NewPropertiesUsecase(mockPropertiesRepository, time.Second*2)
		list, err := u.GetByApplicationAndProfile(context.Background(), "application", "profile")

		assert.NoError(t, err)
		assert.NotNil(t, list)
		assert.Len(t, list, len(mockProperties))

		mockPropertiesRepository.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		mockPropertiesRepository.On(
			"FetchByApplicationAndProfile",
			mock.Anything, mock.Anything,
			mock.Anything,
		).Return(nil, errors.New("Unexpected")).Once()
		u := usecase.NewPropertiesUsecase(mockPropertiesRepository, time.Second*2)
		list, err := u.GetByApplicationAndProfile(context.Background(), "application", "profile")

		assert.Error(t, err)
		assert.Nil(t, list)

		mockPropertiesRepository.AssertExpectations(t)
	})
}
