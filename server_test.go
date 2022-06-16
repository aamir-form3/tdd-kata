package tdd_kata

import (
	"fmt"
	"github.com/stretchr/testify/suite"
	"net/http"
	"net/http/httptest"
	"testing"
)

type ServerTestSuite struct {
	suite.Suite
	server *PlayerServer
}

func TestServerSuite(t *testing.T) {
	suite.Run(t, new(ServerTestSuite))
}

func (suite *ServerTestSuite) SetupTest() {
	store := StubPlayerStore{
		map[string]int{
			"Pepper": 20,
			"Floyd":  10,
		},
	}
	suite.server = &PlayerServer{Store: &store}
}

func (suite *ServerTestSuite) newGetRequest(name string) *http.Request {
	request, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/players/%s", name), nil)
	return request
}

func (suite *ServerTestSuite) TestReturnPeppersScore() {
	request := suite.newGetRequest("Pepper")
	response := httptest.NewRecorder()

	suite.server.ServeHTTP(response, request)

	suite.Equal(http.StatusOK, response.Code)
	suite.Equal("20", response.Body.String())
}

func (suite *ServerTestSuite) TestReturnFloydsScore() {
	request := suite.newGetRequest("Floyd")
	response := httptest.NewRecorder()

	suite.server.ServeHTTP(response, request)

	suite.Equal(http.StatusOK, response.Code)
	suite.Equal("10", response.Body.String())
}

func (suite *ServerTestSuite) TestMissingPlayer_returns_404() {
	request := suite.newGetRequest("Apollo")
	response := httptest.NewRecorder()

	suite.server.ServeHTTP(response, request)

	suite.Equal(http.StatusNotFound, response.Code)
}

func (suite *ServerTestSuite) TestPostPlayer_returnAccepted() {
	request, _ := http.NewRequest(http.MethodPost, "/players/Pepper", nil)
	response := httptest.NewRecorder()

	suite.server.ServeHTTP(response, request)

	suite.Equal(http.StatusAccepted, response.Code)
}

type StubPlayerStore struct {
	scores map[string]int
}

func (s *StubPlayerStore) GetPlayerScore(name string) int {
	return s.scores[name]
}
