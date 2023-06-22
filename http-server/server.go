package main

import (
	"fmt"
	"net/http"
	"strings"
)

// PlayerStore stores score information about players.
type PlayerStore interface {
	GetPlayerScore(name string) int
	RecordWin(name string)
}

// PlayerServer is a HTTP interface for player information.
type PlayerServer struct {
	store PlayerStore
}

// So for our new endpoint, we use http.HandlerFunc and an anonymous function
//to w.WriteHeader(http.StatusOK) when '/league' is requested to make
// our new test pass.
func (p *PlayerServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	router := http.NewServeMux()
	router.Handle("/league", http.HandlerFunc(p.leagueHandler))
	router.Handle("/players/", http.HandlerFunc(p.playersHandler))

	router.ServeHTTP(w, r)
}

func (p *PlayerServer) leagueHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func (p *PlayerServer) playersHandler(w http.ResponseWriter, r *http.Request) {
	player := strings.TrimPrefix(r.URL.Path, "/players/")

	switch r.Method {
	case http.MethodPost:
		p.processWin(w, player)
	case http.MethodGet:
		p.showScore(w, player)
	}
}

func (p *PlayerServer) showScore(w http.ResponseWriter, player string) {
	score := p.store.GetPlayerScore(player)

	if score == 0 {
		w.WriteHeader(http.StatusNotFound)
	}

	fmt.Fprint(w, score)
}

func (p *PlayerServer) processWin(w http.ResponseWriter, player string) {
	p.store.RecordWin(player)
	w.WriteHeader(http.StatusAccepted)
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

// func (p *PlayerServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {

// 	router := http.NewServeMux()
// 	router.Handle("/league", http.HandlerFunc(p.leagueHandler))
// 	router.Handle("/players/", http.HandlerFunc(p.playersHandler))

// 	router.ServeHTTP(w, r)
// }

// func (p *PlayerServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {

// 	router := http.NewServeMux()
// 	router.Handle("/league", http.HandlerFunc(p.leagueHandler))
// 	router.Handle("/players/", http.HandlerFunc(p.playersHandler))

// 	router.ServeHTTP(w, r)
// }
