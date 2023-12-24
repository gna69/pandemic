package logger_test

import (
	"testing"

	"github.com/gna69/pandemic/internal/pkg/logger"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap/zapcore"
)

func TestNewLogger(t *testing.T) {
	// act
	logger, closeFunc, err := logger.NewLogger()

	// assert
	assert.NotNil(t, logger)
	assert.NotNil(t, closeFunc)
	assert.NoError(t, err)
}

func TestNewLoggerWithDebugMode(t *testing.T) {
	// act
	logger, closeFunc, err := logger.NewLogger(logger.WithLogLevel(zapcore.DebugLevel))

	// assert
	assert.NotNil(t, logger)
	assert.NotNil(t, closeFunc)
	assert.NoError(t, err)
}
