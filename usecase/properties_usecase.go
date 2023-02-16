package usecase

import (
	"context"
	"time"

	"github.com/Piyawat-T/go-centralize-configuration/domain"
)

type propertiesUsecase struct {
	propertiesRepository domain.PropertiesRepository
	contextTimeout       time.Duration
}

func NewPropertiesUsecase(propertiesRepository domain.PropertiesRepository, timeout time.Duration) domain.PropertiesUseCase {
	return &propertiesUsecase{
		propertiesRepository: propertiesRepository,
		contextTimeout:       timeout,
	}
}

func (usecase *propertiesUsecase) GetProperties(c context.Context) ([]domain.Properties, error) {
	ctx, cancel := context.WithTimeout(c, usecase.contextTimeout)
	defer cancel()

	properties, err := usecase.propertiesRepository.Fetch(ctx)
	if err != nil {
		return nil, err
	}
	return properties, nil
}

func (usecase *propertiesUsecase) GetByApplicationAndProfile(c context.Context, application string, profile string) ([]domain.Properties, error) {
	ctx, cancel := context.WithTimeout(c, usecase.contextTimeout)
	defer cancel()
	properties, err := usecase.propertiesRepository.FetchByApplicationAndProfile(ctx, application, profile)
	if err != nil {
		return nil, err
	}
	return properties, nil
}
