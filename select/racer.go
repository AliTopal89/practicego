package racer

import (
	"net/http"
)

func Racer(a, b string) (winner string) {
	select {
	case <-ping(a):
		return a
	case <-ping(b):
		return b
	}
}

// aDuration := measureResponseTime(a)
// bDuration := measureResponseTime(b)

// if aDuration < bDuration {
// 	return a
// }

// return b

func ping(url string) chan struct{} {
	ch := make(chan struct{})
	// Anonymous function
	go func() {
		http.Get(url)
		close(ch)
	}()
	return ch
}

// func measureResponseTime(url string) time.Duration {
// 	start := time.Now()
// 	http.Get(url)
// 	return time.Since(start)
// }
