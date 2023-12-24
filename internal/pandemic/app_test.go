package pandemic_test

import (
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
	config := pandemic.Config{}
	app := pandemic.NewApplication(&config, s.logger)
	err := app.Init()
	s.Require().NoError(err)

	// act
	err = app.Run()

	// assert
	s.NoError(err)
}
