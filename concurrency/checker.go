package main

type WebsiteChecker func(string) bool

type result struct {
	string
	bool
}

func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)
	resultsChanal := make(chan result)
	for _, url := range urls {
		go func(s string) {
			resultsChanal <- result{s, wc(s)}
		}(url)
	}
	for i := 0; i < len(urls); i++ {
		r := <-resultsChanal
		results[r.string] = r.bool
	}
	return results
}
