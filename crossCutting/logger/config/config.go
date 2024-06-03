package config

import (
	"go.uber.org/zap/zapcore"
)

var LevelMap = map[string]zapcore.Level{
	"debug": zapcore.DebugLevel,
	"info":  zapcore.InfoLevel,
	"warn":  zapcore.WarnLevel,
	"error": zapcore.ErrorLevel,
}

type LoggerConfig struct {
	Level string
}

func (lc LoggerConfig) GetLevel() zapcore.Level {
	return LevelMap[lc.Level]
}
