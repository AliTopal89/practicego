package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

// trying to integrate with: InMemoryPlayerStore and PlayerServer.
// store a variable response
// because we are going to try and get the player's score.
func TestRecordingWinsAndRetrievingThem(t *testing.T) {
	store := NewInMemoryPlayerStore()
	server := PlayerServer{
		store:  store,
		router: &http.ServeMux{},
	}
	player := "Mojojo"

	server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))
	server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))
	server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))

	newFunction(server, player, t)
}

func newFunction(server PlayerServer, player string, t *testing.T) {
	response := httptest.NewRecorder()
	server.ServeHTTP(response, NewGetScoreRequest(player))
	assertStatus(t, response.Code, http.StatusNotFound)

	assertResponseBody(t, response.Body.String(), "404 page not found\n")
}
