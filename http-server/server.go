package main

import (
	"fmt"
	"net/http"
	"strings"
)

type PlayerStore interface {
	GetPlayerScore(name string) int
}

type PlayerServer struct {
	store PlayerStore
}

func (p *PlayerServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	player := strings.TrimPrefix(r.URL.Path, "/players/")

	score := p.store.GetPlayerScore(player)

	if score == 0 {
		w.WriteHeader(http.StatusNotFound)
	}

	fmt.Fprint(w, score)
}

// now we use store.GetPlayerScore to get the score instead of the local function below
// func GetPlayerScore(name string) string {

// 	if name == "Mojojo" {
// 		return "20"
// 	}
// 	if name == "Megatron" {
// 		return "10"
// 	}
// 	return ""

// }

/*
From the DI chapter, we touched on HTTP servers with a Greet function.
We learned that net/http's "ResponseWriter" also implements io Writer
so we can use "fmt.Fprint" to send strings as HTTP responses.
*/

/*
r.URL.Path returns the path of the request which we can then use
`strings.TrimPrefix`` to trim away /players/ to get the requested player.
It's not very robust but will do the trick for now.
*/
