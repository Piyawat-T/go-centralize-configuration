package route

import (
	"time"

	"github.com/Piyawat-T/go-centralize-configuration/api/controller"
	"github.com/Piyawat-T/go-centralize-configuration/bootstrap"
	"github.com/Piyawat-T/go-centralize-configuration/domain"
	"github.com/Piyawat-T/go-centralize-configuration/repository"
	"github.com/Piyawat-T/go-centralize-configuration/usecase"
	"github.com/gin-gonic/gin"
)

func PropertiesRouter(env *bootstrap.Env, timeout time.Duration, db bootstrap.Database, group *gin.RouterGroup) {
	repo := repository.NewPropertiesRepository(db, domain.CollectionProperties)
	cont := controller.PropertiesController{
		PropertiesUsecase: usecase.NewPropertiesUsecase(repo, timeout),
		Env:               env,
	}
	group.GET("/properties", cont.GetProperties)
}
