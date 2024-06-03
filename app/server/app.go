package server

import (
	"context"
	"myapp/app/api/contract"
	"myapp/app/api/router"
	"myapp/app/server/config"
	eh "myapp/crossCutting/api/middleware/errorhandler"
	"myapp/crossCutting/api/middleware/requestlogger"
	"myapp/crossCutting/api/middleware/securityheaders"
	"myapp/crossCutting/api/validator"
	"myapp/crossCutting/logger"
	"time"

	"github.com/gin-contrib/cors"
	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/penglongli/gin-metrics/ginmetrics"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
)

/*
Adds required middlewares and register routes
By default we have, otelgin for tracing, request logging, security headers, cors and more can be added as required.
*/

func setUpApp(cfg config.ServerConfig, factory contract.OperationHandlerFactory) *gin.Engine {
	gin.SetMode(cfg.App.GinMode)
	engine := gin.New()
	binding.Validator = validator.NewStructValidator()

	engine.
		Use(otelgin.Middleware("payment-service-api")).
		Use(ginzap.RecoveryWithZap(logger.GetLogger(context.TODO()).BaseLogger(), true)).
		Use(requestlogger.Handler(&requestlogger.Config{
			TimeFormat: time.RFC3339, UTC: true, SkipPaths: []string{"/health", "/metrics"}})).
		Use(setUpCors(cfg)).
		Use(securityheaders.Add()).
		Use(eh.NewErrorHandlerMiddleware().HandleResponse)

	setupReqResPrometheus(engine)
	return router.RegisterRoutes(factory, engine)
}

func setUpCors(cfg config.ServerConfig) gin.HandlerFunc {
	// TODO: cors configs
	conf := cors.DefaultConfig()
	conf.AllowOrigins = cfg.App.CorsConfig
	return cors.New(conf)
}

func setupReqResPrometheus(engine *gin.Engine) {
	m := ginmetrics.GetMonitor()
	m.SetMetricPath("/metrics")
	m.Use(engine)
}
