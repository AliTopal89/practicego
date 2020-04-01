package concurrency

import "time"

type WebsiteChecker func(string) bool

func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)

	for _, url := range urls {
		// Anonymous function
		go func(u string) {
			results[u] = wc(u)

		}(url) // expression in go must be function call

		time.Sleep(time.Second)
	}

	return results
}
