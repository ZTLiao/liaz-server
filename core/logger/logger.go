package logger

import (
	"core/application"
)

func Trace(format string, args ...interface{}) {
	var logger = application.GetLogger()
	logger.Tracef(format, args...)
}

func Debug(format string, args ...interface{}) {
	var logger = application.GetLogger()
	logger.Debugf(format, args...)
}

func Info(format string, args ...interface{}) {
	var logger = application.GetLogger()
	logger.Infof(format, args...)
}

func Warn(format string, args ...interface{}) {
	var logger = application.GetLogger()
	logger.Warnf(format, args...)
}

func Error(format string, args ...interface{}) {
	var logger = application.GetLogger()
	logger.Errorf(format, args...)
}

func Fatal(format string, args ...interface{}) {
	var logger = application.GetLogger()
	logger.Fatalf(format, args...)
}

func Panic(format string, args ...interface{}) {
	var logger = application.GetLogger()
	logger.Panicf(format, args...)
}
