package v1_test

import (
	_ "embed"
	"net/http"
	"net/http/httptest"
	"testing"

	v1 "github.com/gna69/pandemic/internal/pandemic/api/v1"
	"github.com/stretchr/testify/suite"
)

var (
	//go:embed testdata/map.json
	expectedMapBody []byte
)

type MapHandlerSuite struct {
	suite.Suite
}

func TestMapHandlerSuite(t *testing.T) {
	suite.Run(t, &MapHandlerSuite{})
}

func (s *MapHandlerSuite) TestGetMapHandler() {
	// arrange
	request, err := http.NewRequest(http.MethodGet, "/api/v1/map", nil)
	s.Require().NoError(err)

	response := httptest.NewRecorder()

	handler := http.HandlerFunc(v1.GetMapHandler())

	// act
	handler.ServeHTTP(response, request)

	// assert
	s.Equal(http.StatusOK, response.Code)
	s.Equal("application/json", response.Header().Get("Content-Type"))
	s.Equal(expectedMapBody, response.Body.Bytes())
}
