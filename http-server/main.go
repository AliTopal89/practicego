package main

import (
	"log"
	"net/http"
)

func main() {
	server := &PlayerServer{NewInMemoryPlayerStore()}
	log.Fatal(http.ListenAndServe(":5000", server))
}

/*
From the documentation, we see that type `HandlerFunc` has already
implemented the `ServeHTTP` method. By type casting our PlayerServer
function with it, we have now implemented the required Handler.
*/
