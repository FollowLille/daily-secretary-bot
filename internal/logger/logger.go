package logger

import "go.uber.org/zap"

func (l *CompositeLogger) Info(msg string, fields ...zap.Field) {
	l.commonLog.Info(msg, fields...)
	l.outputLog.Info(msg, fields...)
	if l.console != nil {
		l.console.Info(msg, fields...)
	}
}

func (l *CompositeLogger) Error(msg string, fields ...zap.Field) {
	l.commonLog.Error(msg, fields...)
	l.errorLog.Error(msg, fields...)
	if l.console != nil {
		l.console.Error(msg, fields...)
	}
}

func (l *CompositeLogger) Debug(msg string, fields ...zap.Field) {
	l.outputLog.Debug(msg, fields...)
	if l.console != nil {
		l.console.Debug(msg, fields...)
	}
}

func (l *CompositeLogger) Sync() error {
	_ = l.commonLog.Sync()
	_ = l.errorLog.Sync()
	_ = l.outputLog.Sync()
	if l.console != nil {
		_ = l.console.Sync()
	}
	return nil
}
