package main

import (
	"net/http"
	"time"
)

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

func Racer(a, b string) (winner string) {
	startA := time.Now()
	http.Get(a)
	aDuration := time.Since(startA)

	startB := time.Now()
	http.Get(b)
	bDuration := time.Since(startB)
	if aDuration < bDuration {
		return a
	}
	return b
}
