package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGETPlayers(t *testing.T) {
	t.Run("returns Mojojo's score", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/players/Mojojo", nil)
		response := httptest.NewRecorder()

		PlayerServer(response, request)

		got := response.Body.String()
		want := "20"

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})

	t.Run("returns Megatron's score", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/players/Megatron", nil)
		response := httptest.NewRecorder()

		PlayerServer(response, request)

		got := response.Body.String()
		want := "10"

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})
}

/*
We use "http.NewRequest" to create a request.
- The first argument is the request's method and the second is the request's path.
The "nil" argument refers to the request's body, which we don't need to set in this case.

- "net/http/httptest" package has a spy already made for us called "ResponseRecorder" so we can use that.
It has many helpful methods to inspect what has been written as a response.
*/
