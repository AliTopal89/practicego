package concurrency

import (
	"testing"
	"time"
)

func slowStubWebsiteChecker(_ string) bool {
	time.Sleep(20 * time.Millisecond)
	return true
}

func BenchmarkCheckWebsites(b *testing.B) {
	urls := make([]string, 100)
	for i := 0; i < len(urls); i++ {
		urls[i] = "a url"
	}

	for i := 0; i < b.N; i++ {
		// fmt.Printf("Current Time: %v\n", time.Now())
		CheckWebsites(slowStubWebsiteChecker, urls)
	}
}

/*
  The benchmark tests CheckWebsites using a slice of one hundred
  urls and uses a new fake implementation of "WebsiteChecker".
  "slowStubWebsiteChecker" is deliberately slow. It uses "time.Sleep"
  to wait exactly twenty milliseconds and then it returns true
*/
