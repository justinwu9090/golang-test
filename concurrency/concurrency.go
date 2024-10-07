package concurrency

// websitechecker is a function that takes a single URL and returns a bool
type WebsiteChecker func(string) bool

// result is associated with the return value of WebsiteChecker
type result struct {
	string //anonymous values
	bool
}

func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)
	resultChannel := make(chan result) // create a channel of type result

	for _, url := range urls {
		// anonymous function that launches go routine, both declared and called -> () at the end of the anon function definition
		go func(u string) {
			// send statement
			resultChannel <- result{u, wc(u)}
		}(url) // pass in url as value so that they don't have the same reference to url!
	}
	for i := 0; i < len(urls); i++ {
		// receive statement coming out of the channel
		r := <-resultChannel
		results[r.string] = r.bool
	}
	return results
}
