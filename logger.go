package logger

import (
	"github.com/timandy/routine"
	"go.uber.org/zap"
)

type logWithAdditional struct {
	*zap.SugaredLogger
	threadLocal routine.ThreadLocal[[]any]
}

func (l *logWithAdditional) Infow(msg string, keysAndValues ...interface{}) {
	keysAndValues = append(keysAndValues, l.threadLocal.Get()...)
	l.SugaredLogger.WithOptions(zap.AddCallerSkip(1)).Infow(msg, keysAndValues...)
}

func (l *logWithAdditional) Debugw(msg string, keysAndValues ...interface{}) {
	keysAndValues = append(keysAndValues, l.threadLocal.Get()...)
	l.SugaredLogger.WithOptions(zap.AddCallerSkip(1)).Debugw(msg, keysAndValues...)
}

func (l *logWithAdditional) Errorw(msg string, keysAndValues ...interface{}) {
	keysAndValues = append(keysAndValues, l.threadLocal.Get()...)
	l.SugaredLogger.WithOptions(zap.AddCallerSkip(1)).Errorw(msg, keysAndValues...)
}

func (l *logWithAdditional) Warnw(msg string, keysAndValues ...interface{}) {
	keysAndValues = append(keysAndValues, l.threadLocal.Get()...)
	l.SugaredLogger.WithOptions(zap.AddCallerSkip(1)).Warnw(msg, keysAndValues...)
}

func (l *logWithAdditional) AddAdditionalInfo(k, v any) {
	l.threadLocal.Set(append(l.threadLocal.Get(), k, v))
}

func (l *logWithAdditional) DeleteAdditionalInfo(layer int) {
	if layer < 0 {
		l.threadLocal.Set([]any{})
		return
	}
	oldKv := l.threadLocal.Get()
	if len(oldKv) < layer*2 {
		l.threadLocal.Set([]any{})
		return
	}
	l.threadLocal.Set(oldKv[:len(oldKv)-2*layer])
}

func NewDefaultLogger() *logWithAdditional {
	log, _ := zap.NewDevelopment()
	return &logWithAdditional{
		SugaredLogger: log.Sugar(),
		threadLocal:   routine.NewThreadLocal[[]any](),
	}
}
