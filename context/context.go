package context

import (
	"fmt"
	"net/http"
)

/*
The function Server takes a Store and returns us a
"http.HandlerFunc". Store is defined as
the returned function calls the store's Fetch method
to get the data and writes it to the response.
*/

type Store interface {
	Fetch() string
}

func Server(store Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, store.Fetch())
	}
}
