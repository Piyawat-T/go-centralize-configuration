package controller

import (
	"errors"
	"net/http"

	"github.com/Piyawat-T/go-centralize-configuration/bootstrap"
	"github.com/Piyawat-T/go-centralize-configuration/domain"
	"github.com/gin-gonic/gin"
	"github.com/uptrace/opentelemetry-go-extra/otelzap"
	"go.uber.org/zap"
)

type PropertiesController struct {
	PropertiesUsecase domain.PropertiesUseCase
	Env               *bootstrap.Env
}

func (controller *PropertiesController) GetProperties(c *gin.Context) {
	ctx := c.Request.Context()
	log := otelzap.L()
	log.Ctx(ctx).Error("Hello from zap",
		zap.Error(errors.New("Hello World!")),
		zap.String("foo", "bar"))
	log.ErrorContext(ctx, "Hello from zap",
		zap.Error(errors.New("Hello World!")),
		zap.String("foo", "bar"))

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

	ctx := c.Request.Context()
	log := otelzap.L()
	log.Ctx(ctx).Error("Hello from zap",
		zap.Error(errors.New("Hello World!")),
		zap.String("foo", "bar"))
	log.DebugContext(ctx, "hi")

	properties, err := controller.PropertiesUsecase.GetByApplicationAndProfile(c, application, profile)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, properties)
}
