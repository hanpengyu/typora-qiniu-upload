package logger

import (
	"go.uber.org/zap"
)

var _log *Logger

type Logger struct {
	zapLogger *zap.Logger
}

func setLogger(logger *Logger) {
	_log = logger
}

func GetLog() *zap.Logger {
	return _log.zapLogger
}

func Info(message string, fields interface{}) {
	GetLog().Info(message, zap.Any("INFO", fields))
}
