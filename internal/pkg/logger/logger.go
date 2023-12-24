package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type (
	SyncFunc func()
	Option   func(*zap.Config)
)

func NewLogger(options ...Option) (*zap.SugaredLogger, SyncFunc, error) {
	zapConfig := zap.Config{
		Level:            zap.NewAtomicLevelAt(zapcore.InfoLevel),
		Development:      true,
		Encoding:         "json",
		EncoderConfig:    zap.NewProductionEncoderConfig(),
		OutputPaths:      []string{"stdout"},
		ErrorOutputPaths: []string{"stderr"},
	}

	for _, option := range options {
		option(&zapConfig)
	}

	logger, err := zapConfig.Build()
	if err != nil {
		return nil, nil, err
	}

	syncFunc := func() {
		_ = logger.Sync()
	}
	return logger.Sugar(), syncFunc, nil
}
