package log

import (
	"context"
	"github.com/rs/zerolog"
	"runtime"
	"strconv"
)

type logContext struct {
	log    *Logger
	Level  zerolog.Level
	fields map[string]interface{}
}

func (l *logContext) Field(k string, v interface{}) *logContext {
	l.fields[k] = v
	return l
}

func (l *logContext) Fields(fields map[string]interface{}) *logContext {
	for k, v := range fields {
		l.fields[k] = v
	}
	return l
}

func (l *logContext) Stack(depth int) *logContext {
	stack := make([]string, 0, depth)
	for i := 1; i <= depth; i++ {
		_, file, line, ok := runtime.Caller(i)
		if ok {
			stack = append(stack, file+":"+strconv.Itoa(line))
		}
	}
	l.fields["stack"] = stack
	return l
}

func (l *logContext) TraceID(ctx context.Context) *logContext {
	if l.log.traceIdFunc != nil {
		l.fields["traceId"] = l.log.traceIdFunc(ctx)
	}
	return l
}

func (l *logContext) Info(args ...interface{}) {
	l.log.msg(zerolog.InfoLevel, l.fields, args...)
}

func (l *logContext) Warn(args ...interface{}) {
	l.log.msg(zerolog.WarnLevel, l.fields, args...)
}

func (l *logContext) Error(args ...interface{}) {
	l.log.msg(zerolog.ErrorLevel, l.fields, args...)
}

func (l *logContext) Fatal(args ...interface{}) {
	l.log.msg(zerolog.FatalLevel, l.fields, args...)
}

func (l *logContext) Debug(args ...interface{}) {
	l.log.msg(zerolog.DebugLevel, l.fields, args...)
}

func (l *logContext) Panic(args ...interface{}) {
	l.log.msg(zerolog.PanicLevel, l.fields, args...)
}

func (l *logContext) Infof(format string, args ...interface{}) {
	l.log.msgf(zerolog.InfoLevel, l.fields, format, args...)
}

func (l *logContext) Warnf(format string, args ...interface{}) {
	l.log.msgf(zerolog.WarnLevel, l.fields, format, args...)
}

func (l *logContext) Errorf(format string, args ...interface{}) {
	l.log.msgf(zerolog.ErrorLevel, l.fields, format, args...)
}

func (l *logContext) Fatalf(format string, args ...interface{}) {
	l.log.msgf(zerolog.FatalLevel, l.fields, format, args...)
}

func (l *logContext) Debugf(format string, args ...interface{}) {
	l.log.msgf(zerolog.DebugLevel, l.fields, format, args...)
}

func (l *logContext) Panicf(format string, args ...interface{}) {
	l.log.msgf(zerolog.PanicLevel, l.fields, format, args...)
}
