package domain

import "context"

const (
	CollectionProperties = "properties"
)

type Properties struct {
	Id          uint   `json:"id"`
	Application string `json:"application"`
	Profile     string `json:"profile"`
	Key         string `json:"key"`
	Value       string `json:"value"`
}

type PropertiesRepository interface {
	Fetch(c context.Context) ([]Properties, error)
	FetchByApplicationAndProfile(c context.Context, application string, profile string) ([]Properties, error)
}

type PropertiesUseCase interface {
	GetProperties(c context.Context) ([]Properties, error)
	GetByApplicationAndProfile(c context.Context, application string, profile string) ([]Properties, error)
}
