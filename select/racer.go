package racer

import (
	"net/http"
	"time"
)

func Racer(slowURL, fastURL string) (winner string) {

	startA := time.Now()
	http.Get(slowURL)
	aDuration := time.Since(startA)

	startB := time.Now()
	http.Get(fastURL)
	bDuration := time.Since(startB)

	if aDuration < bDuration {
		return slowURL
	}

	return fastURL
}
