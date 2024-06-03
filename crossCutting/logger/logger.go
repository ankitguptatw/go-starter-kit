package logger

import (
	"context"
	"myapp/crossCutting/logger/config"
	"myapp/crossCutting/util"
	"sync"

	"github.com/uptrace/opentelemetry-go-extra/otelzap"
	"go.uber.org/zap"
)

var once sync.Once
var instance *otelzap.Logger

/*
application logger is an abstraction over otelzap logger. This way internal logger can be replaced easily as required
otelzap is a thin wrapper for zap.Logger that adds Ctx to logger methods.This is useful for tracking trace_id.
Therefore, logs can be indexed by trace_id in log aggregator tools like ELK stack.

application logger is a singleton, and we can obtain the logger object using `GetLogger` method wherever required.
*/

type applicationLogger struct {
	lgr otelzap.LoggerWithCtx
}

func InitApplicationLogger(config config.LoggerConfig) {

	once.Do(func() {
		instance = otelzap.New(newZapLogger(
			config.GetLevel()), otelzap.WithTraceIDField(true), otelzap.WithMinLevel(config.GetLevel()))

		defer func(zapLogger *otelzap.Logger) {
			_ = zapLogger.Sync()
		}(instance)
	})
}

func GetLogger(ctx context.Context) *applicationLogger {
	if instance == nil {
		conf := config.LoggerConfig{
			Level: "info",
		}
		InitApplicationLogger(conf)
	}

	if ctx == nil {
		ctx = context.Background()
	}
	return &applicationLogger{lgr: instance.Ctx(util.GetTraceContext(ctx))}
}

func (logger *applicationLogger) BaseLogger() *zap.Logger {
	return logger.lgr.ZapLogger()
}

func (logger *applicationLogger) Debug(msg string, args ...interface{}) {
	logger.lgr.Debug(util.Format(msg, args...))
}

func (logger *applicationLogger) Info(msg string, args ...interface{}) {
	logger.lgr.Info(util.Format(msg, args...))
}

func (logger *applicationLogger) Warn(msg string, args ...interface{}) {
	logger.lgr.Warn(util.Format(msg, args...))
}
func (logger *applicationLogger) Error(msg string, args ...interface{}) {
	logger.lgr.Error(util.Format(msg, args...))
}
