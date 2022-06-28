package log

import (
	"context"
	"io"
	"os"
)

var defaultLogger = NewLogger(os.Stderr)

func InitLogger(w io.Writer) {
	defaultLogger = NewLogger(w)
}

func DefaultLogger() *Logger {
	return defaultLogger
}

func Info(args ...interface{}) {
	defaultLogger.Info(args...)
}

func Warn(args ...interface{}) {
	defaultLogger.Warn(args...)
}

func Error(args ...interface{}) {
	defaultLogger.Error(args...)
}

func Fatal(args ...interface{}) {
	defaultLogger.Fatal(args...)
}

func Debug(args ...interface{}) {
	defaultLogger.Debug(args...)
}

func Panic(args ...interface{}) {
	defaultLogger.Panic(args...)
}

func Infof(format string, args ...interface{}) {
	defaultLogger.Infof(format, args...)
}

func Warnf(format string, args ...interface{}) {
	defaultLogger.Warnf(format, args...)
}

func Errorf(format string, args ...interface{}) {
	defaultLogger.Errorf(format, args...)
}

func Fatalf(format string, args ...interface{}) {
	defaultLogger.Fatalf(format, args...)
}

func Debugf(format string, args ...interface{}) {
	defaultLogger.Debugf(format, args...)
}

func Panicf(format string, args ...interface{}) {
	defaultLogger.Panicf(format, args...)
}

func SetTraceIdFunc(fun func(ctx context.Context) string) {
	defaultLogger.SetTraceIdFunc(fun)
}

func SetLevel(level Level) {
	defaultLogger.SetLevel(level)
}

func With() *logContext {
	return defaultLogger.With()
}
