package pandemic

import (
	"context"
	"fmt"
	"net/http"

	v1 "github.com/gna69/pandemic/internal/pandemic/api/v1"
	"github.com/go-chi/chi/v5"
	"go.uber.org/zap"
)

type Application struct {
	config     *Config
	httpRouter http.Handler
	logger     *zap.SugaredLogger
}

func NewApplication(config *Config, logger *zap.SugaredLogger) *Application {
	return &Application{
		config: config,
		logger: logger,
	}
}

func (app *Application) Init() error {
	app.initAPI()
	return nil
}

func (app *Application) initAPI() {
	router := chi.NewRouter()
	router.Mount(v1.APIPattern, v1.NewRouter())

	app.httpRouter = router
}

func (app *Application) Run(ctx context.Context) error {
	app.logger.Infow("start pandemic server", "port", app.config.Port)
	defer app.logger.Info("pandemic server was stopped")

	go func() {
		if err := http.ListenAndServe(fmt.Sprintf(":%d", app.config.Port), app.httpRouter); err != nil {
			app.logger.Errorw("failed listen and serve application", err)
		}
	}()

	<-ctx.Done()

	return nil
}
