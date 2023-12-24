package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func WithLogLevel(level zapcore.Level) Option {
	return func(cfg *zap.Config) {
		cfg.Level = zap.NewAtomicLevelAt(level)
	}
}
