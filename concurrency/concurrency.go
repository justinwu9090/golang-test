package concurrency

import "time"

// websitechecker is a function that takes a single URL and returns a bool
type WebsiteChecker func(string) bool

func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)

	for _, url := range urls {
		// anonymous function that launches go routine, both declared and called -> () at the end of the anon function definition
		go func(u string) {
			results[url] = wc(url) // this will result in a race condition!
		}(url) // pass in url as value so that they don't have the same reference to url!
	}
	time.Sleep(2 * time.Second)
	return results
}
