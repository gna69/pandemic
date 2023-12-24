package logger_test

import (
	"testing"

	"github.com/gna69/pandemic/internal/pkg/logger"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func TestWithLogLevel(t *testing.T) {
	// arrange
	zapConfig := zap.Config{}
	expectedLogLevel := zap.NewAtomicLevelAt(zapcore.DebugLevel)

	logLevelOption := logger.WithLogLevel(zapcore.DebugLevel)

	// act
	logLevelOption(&zapConfig)

	// assert
	assert.Equal(t, expectedLogLevel, zapConfig.Level)
}
