package pandemic

import "go.uber.org/zap"

type Application struct {
	config *Config
	logger *zap.SugaredLogger
}

func NewApplication(config *Config, logger *zap.SugaredLogger) *Application {
	return &Application{
		config: config,
		logger: logger,
	}
}

func (app *Application) Init() error {
	return nil
}

func (app *Application) Run() error {
	return nil
}
