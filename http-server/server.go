package main

import (
	"fmt"
	"net/http"
	"strings"
)

func PlayerServer(w http.ResponseWriter, r *http.Request) {
	player := strings.TrimPrefix(r.URL.Path, "/players/")

	if player == "Mojojo" {
		fmt.Fprintf(w, "20")
		return
	}
	if player == "Megatron" {
		fmt.Fprintf(w, "10")
		return
	}
}

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
