package repository_test

import (
	"context"
	"errors"
	"testing"

	"github.com/Piyawat-T/go-centralize-configuration/domain"
	"github.com/Piyawat-T/go-centralize-configuration/mysqldb/mocks"
	"github.com/Piyawat-T/go-centralize-configuration/repository"
	_ "github.com/jinzhu/gorm"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestFetch(t *testing.T) {
	var databaseHelper *mocks.Database
	databaseHelper = &mocks.Database{}

	collectionName := domain.CollectionProperties
	t.Run("success", func(t *testing.T) {
		databaseHelper.On("Find", mock.AnythingOfType("*[]domain.Properties")).Run(func(args mock.Arguments) {
			arg := args.Get(0).(*[]domain.Properties)
			properties := []domain.Properties{
				{
					Id: 1,
				},
			}
			*arg = properties
		}).Return(nil)

		repo := repository.NewPropertiesRepository(databaseHelper, collectionName)
		list, err := repo.Fetch(context.Background())
		assert.NoError(t, err)
		assert.NotNil(t, list)

		databaseHelper.AssertExpectations(t)
	})
}

func TestFetch2(t *testing.T) {
	var databaseHelper *mocks.Database
	databaseHelper = &mocks.Database{}

	collectionName := domain.CollectionProperties

	t.Run("error", func(t *testing.T) {
		databaseHelper.On("Find", mock.Anything).Return(errors.New("Unexpected"))
		repo := repository.NewPropertiesRepository(databaseHelper, collectionName)
		list, err := repo.Fetch(context.Background())

		assert.Error(t, err)
		assert.Nil(t, list)
		databaseHelper.AssertExpectations(t)
	})
}

func TestFetchByApplicationAndProfile(t *testing.T) {
	databaseHelper := &mocks.Database{}
	collectionName := domain.CollectionProperties
	t.Run("success", func(t *testing.T) {
		databaseHelper.On("Find", mock.AnythingOfType("*[]domain.Properties"), mock.Anything).Run(func(args mock.Arguments) {
			arg := args.Get(0).(*[]domain.Properties)
			properties := []domain.Properties{
				{
					Id: 1,
				},
			}
			*arg = properties
		}).Return(nil)
		repo := repository.NewPropertiesRepository(databaseHelper, collectionName)
		list, err := repo.FetchByApplicationAndProfile(context.Background(), "application", "profile")
		assert.NoError(t, err)
		assert.NotNil(t, list)

		databaseHelper.AssertExpectations(t)
	})
}

func TestErrorFetchByApplicationAndProfile(t *testing.T) {
	databaseHelper := &mocks.Database{}
	collectionName := domain.CollectionProperties
	t.Run("error", func(t *testing.T) {
		databaseHelper.On("Find", mock.Anything, mock.Anything).Return(errors.New("Unexpected"))
		repo := repository.NewPropertiesRepository(databaseHelper, collectionName)
		list, err := repo.FetchByApplicationAndProfile(context.Background(), "application", "profile")

		assert.Error(t, err)
		assert.Nil(t, list)
		databaseHelper.AssertExpectations(t)
	})
}
