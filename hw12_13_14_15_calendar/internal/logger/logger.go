package logger

import "github.com/VictoriaMetrics/VictoriaMetrics/lib/logger"

type Logger struct {
	level string
}

func New(level string) *Logger {
	return &Logger{
		level: level,
	}
}

func (l Logger) Info(msg string) {
	if l.level == "info" || l.level == "debug" {
		logger.Infof(msg)
	}
}

func (l Logger) Warn(msg string) {
	if l.level == "warn" || l.level == "debug" {
		logger.Warnf(msg)
	}
}

func (l Logger) Error(msg string) {
	logger.Errorf("%s", msg)
}
