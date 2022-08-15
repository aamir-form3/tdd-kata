package server

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

func (s *ServerTestSuite) SetupTest() {
	store := StubPlayerStore{
		map[string]int{
			"Pepper": 20,
			"Floyd":  10,
		},
	}
	s.server = &PlayerServer{Store: &store}
}

func (s *ServerTestSuite) newGetRequest(name string) *http.Request {
	request, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/players/%s", name), nil)
	return request
}

func (s *ServerTestSuite) TestReturnPeppersScore() {
	request := s.newGetRequest("Pepper")
	response := httptest.NewRecorder()

	s.server.ServeHTTP(response, request)

	s.Equal(http.StatusOK, response.Code)
	s.Equal("20", response.Body.String())
}

func (s *ServerTestSuite) TestReturnFloydsScore() {
	request := s.newGetRequest("Floyd")
	response := httptest.NewRecorder()

	s.server.ServeHTTP(response, request)

	s.Equal(http.StatusOK, response.Code)
	s.Equal("10", response.Body.String())
}

func (s *ServerTestSuite) TestMissingPlayer_returns_404() {
	request := s.newGetRequest("Apollo")
	response := httptest.NewRecorder()

	s.server.ServeHTTP(response, request)

	s.Equal(http.StatusNotFound, response.Code)
}

func (s *ServerTestSuite) TestPostPlayer_returnAccepted() {
	request, _ := http.NewRequest(http.MethodPost, "/players/Pepper", nil)
	response := httptest.NewRecorder()

	s.server.ServeHTTP(response, request)

	s.Equal(http.StatusAccepted, response.Code)
}

type StubPlayerStore struct {
	scores map[string]int
}

func (s *StubPlayerStore) GetPlayerScore(name string) int {
	return s.scores[name]
}
