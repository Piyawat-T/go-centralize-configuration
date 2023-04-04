package controller_test

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Piyawat-T/go-centralize-configuration/api/controller"
	"github.com/Piyawat-T/go-centralize-configuration/domain"
	"github.com/Piyawat-T/go-centralize-configuration/domain/mocks"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetProperties(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		mockProperties := []domain.Properties{
			{
				Id:          1,
				Application: "deposit",
				Profile:     "default",
				Key:         "PORT",
				Value:       "8101",
			},
		}

		mockPropertiesUsecase := new(mocks.PropertiesUseCase)
		mockPropertiesUsecase.On("GetProperties", mock.Anything).Return(mockProperties, nil)

		cont := &controller.PropertiesController{
			PropertiesUsecase: mockPropertiesUsecase,
		}

		body, err := json.Marshal(mockProperties)
		assert.NoError(t, err)
		bodyString := string(body)

		r := gin.Default()
		r.GET("/properties", cont.GetProperties)

		rec := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodGet, "/properties", nil)
		r.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, bodyString, rec.Body.String())
	})

	t.Run("error", func(t *testing.T) {
		mockPropertiesUsecase := new(mocks.PropertiesUseCase)
		customErr := errors.New("Unexpected")
		mockPropertiesUsecase.On("GetProperties", mock.Anything).Return(nil, customErr)

		cont := &controller.PropertiesController{
			PropertiesUsecase: mockPropertiesUsecase,
		}

		body, err := json.Marshal(domain.ErrorResponse{Message: customErr.Error()})
		assert.NoError(t, err)
		bodyString := string(body)

		r := gin.Default()
		r.GET("/properties", cont.GetProperties)

		rec := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodGet, "/properties", nil)
		r.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusInternalServerError, rec.Code)
		assert.Equal(t, bodyString, rec.Body.String())
	})
}

func TestGetConfiguration(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		mockProperties := []domain.Properties{
			{
				Id:          1,
				Application: "deposit",
				Profile:     "default",
				Key:         "PORT",
				Value:       "8101",
			},
		}

		mockPropertiesUsecase := new(mocks.PropertiesUseCase)
		mockPropertiesUsecase.On(
			"GetByApplicationAndProfile",
			mock.Anything,
			"deposit",
			"default",
		).Return(mockProperties, nil)

		cont := &controller.PropertiesController{
			PropertiesUsecase: mockPropertiesUsecase,
		}

		body, err := json.Marshal(mockProperties)
		assert.NoError(t, err)
		bodyString := string(body)

		r := gin.Default()
		r.GET("/:application/:profile", cont.GetConfiguration)

		rec := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodGet, "/deposit/default", nil)
		r.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, bodyString, rec.Body.String())
	})

	t.Run("error", func(t *testing.T) {
		mockPropertiesUsecase := new(mocks.PropertiesUseCase)
		customErr := errors.New("Unexpected")
		mockPropertiesUsecase.On(
			"GetByApplicationAndProfile",
			mock.Anything,
			"deposit",
			"default",
		).Return(nil, customErr)

		cont := &controller.PropertiesController{
			PropertiesUsecase: mockPropertiesUsecase,
		}

		body, err := json.Marshal(domain.ErrorResponse{Message: customErr.Error()})
		assert.NoError(t, err)
		bodyString := string(body)

		r := gin.Default()
		r.GET("/:application/:profile", cont.GetConfiguration)

		rec := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodGet, "/deposit/default", nil)
		r.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusInternalServerError, rec.Code)
		assert.Equal(t, bodyString, rec.Body.String())
	})
}
