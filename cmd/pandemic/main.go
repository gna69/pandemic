package main

import (
	"context"
	"fmt"
	"log"

	"github.com/caarlos0/env/v9"
	"github.com/gna69/pandemic/internal/pandemic"
	"github.com/gna69/pandemic/internal/pkg/logger"
	"go.uber.org/zap/zapcore"
)

func main() {
	var cfg pandemic.Config
	if err := env.Parse(&cfg); err != nil {
		log.Fatal(fmt.Errorf("failed read envs: %w", err))
	}

	var loggerOptions []logger.Option
	if cfg.Debug {
		loggerOptions = append(loggerOptions, logger.WithLogLevel(zapcore.DebugLevel))
	}

	appLogger, syncLogger, err := logger.NewLogger(loggerOptions...)
	if err != nil {
		log.Fatal(fmt.Errorf("failed to init logger: %w", err))
	}
	defer syncLogger()

	app := pandemic.NewApplication(&cfg, appLogger)
	if err := app.Init(); err != nil {
		appLogger.Fatalw("failed init application", "err", err)
	}

	if err := app.Run(context.Background()); err != nil {
		appLogger.Fatalw("failed run application", "err", err)
	}
}
