package route

import (
	"time"

	"github.com/Piyawat-T/go-centralize-configuration/bootstrap"
	"github.com/gin-gonic/gin"
)

func Setup(env *bootstrap.Env, timeout time.Duration, db bootstrap.Database, routerV1 *gin.RouterGroup) {
	publicRouterV1 := routerV1.Group("")
	// All Public APIs
	HomeRouter(env, timeout, publicRouterV1)
	PropertiesRouter(env, timeout, db, publicRouterV1)
}
