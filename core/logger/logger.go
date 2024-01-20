package logger

import (
	"core/system"
	"runtime/debug"
)

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
	stack := debug.Stack()
	if len(stack) != 0 {
		system.GetLogger().Errorf("%v", string(stack))
	}
	system.GetLogger().Errorf(format, args...)
}

func Fatal(format string, args ...interface{}) {
	system.GetLogger().Fatalf(format, args...)
}

func Panic(format string, args ...interface{}) {
	system.GetLogger().Panicf(format, args...)
}
