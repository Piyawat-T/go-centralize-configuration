package repository

import (
	"context"

	"github.com/Piyawat-T/go-centralize-configuration/bootstrap"
	"github.com/Piyawat-T/go-centralize-configuration/domain"
)

type propertiesRepository struct {
	database   bootstrap.Database
	collection string
}

func NewPropertiesRepository(db bootstrap.Database, collection string) domain.PropertiesRepository {
	return &propertiesRepository{
		database:   db,
		collection: collection,
	}
}

func (ur *propertiesRepository) Fetch(c context.Context) ([]domain.Properties, error) {
	var properties []domain.Properties
	err := ur.database.Find(&properties)
	return properties, err
}

func (repository *propertiesRepository) FetchByApplicationAndProfile(c context.Context, application string, profile string) ([]domain.Properties, error) {
	var properties []domain.Properties
	err := repository.database.Find(&properties, "application = ? AND profile = ?", application, profile)
	return properties, err
}
