package main

type WebSiteChecker func(string) bool
type result struct {
	string
	bool
}

// CheckWebsites returns the results of the WebSiteChecker
func CheckWebsites(wc WebSiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)
	resultChannel := make(chan result)

	for _, url := range urls {
		go func(u string) {
			// Send statement to the resultChannel
			resultChannel <- result{u, wc(u)}
		}(url)
	}

	// Recieve expression will assign the value from resultChannel to a variable
	for i := 0; i < len(urls); i++ {
		result := <-resultChannel
		results[result.string] = result.bool
	}
	return results
}
