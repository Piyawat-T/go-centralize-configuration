package controller

import (
	"net/http"

	"github.com/Piyawat-T/go-centralize-configuration/bootstrap"
	"github.com/Piyawat-T/go-centralize-configuration/domain"
	"github.com/gin-gonic/gin"
)

type PropertiesController struct {
	PropertiesUsecase domain.PropertiesUseCase
	Env               *bootstrap.Env
}

func (controller *PropertiesController) GetProperties(c *gin.Context) {
	properties, err := controller.PropertiesUsecase.GetProperties(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, properties)
}

func (controller *PropertiesController) GetConfiguration(c *gin.Context) {
	application := c.Param("application")
	profile := c.Param("profile")

	properties, err := controller.PropertiesUsecase.GetByApplicationAndProfile(c, application, profile)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, properties)
}
