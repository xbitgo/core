package log

import (
	"context"
	"fmt"
	"io"
	"strings"
	"time"

	"github.com/rs/zerolog"
)

// Logger ...
type Logger struct {
	zLogger     zerolog.Logger
	traceIdFunc func(ctx context.Context) string
}

func (l *Logger) msg(level zerolog.Level, fields map[string]interface{}, msg ...interface{}) {
	switch len(msg) {
	case 0:
		return
	case 1:
		switch v := msg[0].(type) {
		case string:
			l.msgf(level, fields, v)
		default:
			l.msgf(level, fields, "%v", v)
		}
	default:
		format := strings.Repeat("%v, ", len(msg))
		l.msgf(level, fields, format[:len(format)-2], msg)
	}
}

func (l *Logger) msgf(level zerolog.Level, fields map[string]interface{}, format string, v ...interface{}) {
	fields["_ts"] = fmt.Sprintf("%d.%d", time.Now().Unix(), time.Now().Nanosecond()/1000)
	if len(v) == 0 {
		l.zLogger.WithLevel(level).Fields(fields).Msg(format)
		return
	}
	l.zLogger.WithLevel(level).Fields(fields).Msgf(format, v...)
}

func (l *Logger) Info(args ...interface{}) {
	l.msg(zerolog.InfoLevel, map[string]interface{}{}, args...)
}

func (l *Logger) Warn(args ...interface{}) {
	l.msg(zerolog.WarnLevel, map[string]interface{}{}, args...)
}

func (l *Logger) Error(args ...interface{}) {
	l.msg(zerolog.ErrorLevel, map[string]interface{}{}, args...)
}

func (l *Logger) Fatal(args ...interface{}) {
	l.msg(zerolog.FatalLevel, map[string]interface{}{}, args...)
}

func (l *Logger) Debug(args ...interface{}) {
	l.msg(zerolog.DebugLevel, map[string]interface{}{}, args...)
}

func (l *Logger) Panic(args ...interface{}) {
	l.msg(zerolog.PanicLevel, map[string]interface{}{}, args...)
}

func (l *Logger) Infof(format string, args ...interface{}) {
	l.msgf(zerolog.InfoLevel, map[string]interface{}{}, format, args...)
}

func (l *Logger) Warnf(format string, args ...interface{}) {
	l.msgf(zerolog.WarnLevel, map[string]interface{}{}, format, args...)
}

func (l *Logger) Errorf(format string, args ...interface{}) {
	l.msgf(zerolog.ErrorLevel, map[string]interface{}{}, format, args...)
}

func (l *Logger) Fatalf(format string, args ...interface{}) {
	l.msgf(zerolog.FatalLevel, map[string]interface{}{}, format, args...)
}

func (l *Logger) Debugf(format string, args ...interface{}) {
	l.msgf(zerolog.DebugLevel, map[string]interface{}{}, format, args...)
}

func (l *Logger) Panicf(format string, args ...interface{}) {
	l.msgf(zerolog.PanicLevel, map[string]interface{}{}, format, args...)
}

func (l *Logger) With() *logContext {
	return &logContext{
		log:    l,
		fields: map[string]interface{}{},
	}
}

func (l *Logger) SetTraceIdFunc(fun func(ctx context.Context) string) {
	l.traceIdFunc = fun
}

func (l *Logger) SetLevel(level Level) {
	l.zLogger = l.zLogger.Level(zerolog.Level(level))
}

func NewLogger(w io.Writer) *Logger {
	return &Logger{
		zLogger: zerolog.New(w),
	}
}
