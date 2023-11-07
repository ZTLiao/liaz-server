package logger

import (
	"core/application"
)

type ILogger interface {
	Trace(string, ...interface{})
	Debug(string, ...interface{})
	Info(string, ...interface{})
	Warn(string, ...interface{})
	Error(string, ...interface{})
	Fatal(string, ...interface{})
	Panic(string, ...interface{})
}

func Trace(format string, args ...interface{}) {
	var logger = application.GetApp().GetLogger()
	logger.Tracef(format, args...)
}

func Debug(format string, args ...interface{}) {
	var logger = application.GetApp().GetLogger()
	logger.Debugf(format, args...)
}

func Info(format string, args ...interface{}) {
	var logger = application.GetApp().GetLogger()
	logger.Infof(format, args...)
}

func Warn(format string, args ...interface{}) {
	var logger = application.GetApp().GetLogger()
	logger.Warnf(format, args...)
}

func Error(format string, args ...interface{}) {
	var logger = application.GetApp().GetLogger()
	logger.Errorf(format, args...)
}

func Fatal(format string, args ...interface{}) {
	var logger = application.GetApp().GetLogger()
	logger.Fatalf(format, args...)
}

func Panic(format string, args ...interface{}) {
	var logger = application.GetApp().GetLogger()
	logger.Panicf(format, args...)
}
