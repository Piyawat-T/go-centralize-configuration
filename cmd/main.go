package main

import (
	"time"

	route "github.com/Piyawat-T/go-centralize-configuration/api/route"
	routeV1 "github.com/Piyawat-T/go-centralize-configuration/api/route/v1"

	"github.com/Piyawat-T/go-centralize-configuration/bootstrap"

	"github.com/gin-gonic/gin"
)

func main() {
	app := bootstrap.App()

	env := app.Env

	defer app.CloseDatabaseConnection()

	timeout := time.Duration(env.ContextTimeout) * time.Second

	r := gin.Default()

	routerV1 := r.Group("v1")

	routeV1.Setup(env, timeout, app.Database, routerV1)

	routeV0 := r.Group(env.ContextPath)
	route.Setup(env, timeout, app.Database, routeV0)

	r.Run(env.ServerAddress)
}
