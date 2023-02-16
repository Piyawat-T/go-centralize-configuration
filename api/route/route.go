package route

import (
	"time"

	"github.com/Piyawat-T/go-centralize-configuration/bootstrap"
	"github.com/gin-gonic/gin"
)

func Setup(env *bootstrap.Env, timeout time.Duration, db bootstrap.Database, route *gin.RouterGroup) {
	publicRouter := route.Group("")
	PropertiesRouter(env, timeout, db, publicRouter)
}
