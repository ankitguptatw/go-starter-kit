package logger

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func newZapLogger(level zapcore.Level) *zap.Logger {
	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderConfig()),
		zapcore.NewMultiWriteSyncer(os.Stdout),
		zap.NewAtomicLevelAt(level),
	)

	return zap.New(core, zap.AddCaller())
}

func encoderConfig() zapcore.EncoderConfig {
	return zapcore.EncoderConfig{
		TimeKey:        "@timestamp",
		LevelKey:       "level",
		CallerKey:      "caller",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		EncodeCaller:   zapcore.ShortCallerEncoder,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.NanosDurationEncoder,
	}
}
