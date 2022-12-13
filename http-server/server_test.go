package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

// A map is a quick and easy way of making a stub key/value store for our tests
type StubPlayerStore struct {
	scores map[string]int
}

func (s *StubPlayerStore) GetPlayerScore(name string) int {
	score := s.scores[name]
	return score
}

func TestGETPlayers(t *testing.T) {
	store := StubPlayerStore{
		map[string]int{
			"Mojojo":   20,
			"Megatron": 10,
		},
	}
	server := &PlayerServer{&store}
	t.Run("returns Mojojo's score", func(t *testing.T) {
		request := NewGetScoreRequest("Mojojo")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertResponseBody(t, response.Body.String(), "20")
	})

	t.Run("returns Megatron's score", func(t *testing.T) {
		request := NewGetScoreRequest("Megatron")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertResponseBody(t, response.Body.String(), "10")
	})
}

func NewGetScoreRequest(name string) *http.Request {
	req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/players/%s", name), nil)
	return req
}

func assertResponseBody(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("response of the body is wrong got %q want %q", got, want)
	}
}

/*
We use "http.NewRequest" to create a request.
- The first argument is the request's method and the second is the request's path.
The "nil" argument refers to the request's body, which we don't need to set in this case.

- "net/http/httptest" package has a spy already made for us called "ResponseRecorder" so we can use that.
It has many helpful methods to inspect what has been written as a response.
*/
