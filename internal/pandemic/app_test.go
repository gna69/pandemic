package pandemic_test

import (
	"context"
	"testing"

	"github.com/gna69/pandemic/internal/pandemic"
	"github.com/stretchr/testify/suite"
	"go.uber.org/zap"
)

type ApplicationSuite struct {
	suite.Suite

	logger *zap.SugaredLogger
}

func TestApplicationSuite(t *testing.T) {
	suite.Run(t, new(ApplicationSuite))
}

func (s *ApplicationSuite) SetupSuite() {
	s.logger = zap.NewNop().Sugar()
}

func (s *ApplicationSuite) TestSuccessInitApplication() {
	// arrange
	config := pandemic.Config{}
	app := pandemic.NewApplication(&config, s.logger)

	// act
	err := app.Init()

	// assert
	s.NoError(err)
}

func (s *ApplicationSuite) TestSuccessRunApplication() {
	// arrange
	ctx, cancel := context.WithCancel(context.Background())
	config := pandemic.Config{}
	app := pandemic.NewApplication(&config, s.logger)
	err := app.Init()
	s.Require().NoError(err)

	// act
	cancel()
	err = app.Run(ctx)

	// assert
	s.NoError(err)
}
