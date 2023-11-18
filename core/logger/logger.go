package logger

import "core/system"

func Trace(format string, args ...interface{}) {
	system.GetLogger().Tracef(format, args...)
}

func Debug(format string, args ...interface{}) {
	system.GetLogger().Debugf(format, args...)
}

func Info(format string, args ...interface{}) {
	system.GetLogger().Infof(format, args...)
}

func Warn(format string, args ...interface{}) {
	system.GetLogger().Warnf(format, args...)
}

func Error(format string, args ...interface{}) {
	system.GetLogger().Errorf(format, args...)
}

func Fatal(format string, args ...interface{}) {
	system.GetLogger().Fatalf(format, args...)
}

func Panic(format string, args ...interface{}) {
	system.GetLogger().Panicf(format, args...)
}
