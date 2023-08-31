package main

import (
	"context"
	"time"

	route "github.com/Piyawat-T/go-centralize-configuration/api/route"
	routeV1 "github.com/Piyawat-T/go-centralize-configuration/api/route/v1"
	"github.com/Piyawat-T/go-centralize-configuration/bootstrap/tracer"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"github.com/Piyawat-T/go-centralize-configuration/bootstrap"

	"github.com/gin-gonic/gin"
	"github.com/uptrace/opentelemetry-go-extra/otelzap"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
)

func main() {
	ctx := context.Background()
	tp := tracer.NewTraceProvider()
	tc := tp.Tracer("opentelemetry-logger/main")
	tracer.NewContext(ctx, tc)

	app := bootstrap.App()

	env := app.Env

	defer app.CloseDatabaseConnection()

	timeout := time.Duration(env.ContextTimeout) * time.Second

	logger := otelzap.New(zap.Must(zap.NewDevelopment()),
		otelzap.WithTraceIDField(true),
		otelzap.WithMinLevel(zapcore.DebugLevel),
	)
	defer logger.Sync()

	undo := otelzap.ReplaceGlobals(logger)
	defer undo()

	otelzap.L().Info("Replaced zap's global loggers")
	otelzap.Ctx(context.TODO()).Info("... and with context")

	gin.SetMode(env.GinMode)
	r := gin.Default()
	r.Use(otelgin.Middleware("service-name"))

	routerV1 := r.Group("v1")

	routeV1.Setup(env, timeout, app.Database, routerV1)

	routeV0 := r.Group(env.ContextPath)
	route.Setup(env, timeout, app.Database, routeV0)

	r.Run(env.ServerAddress)
}
