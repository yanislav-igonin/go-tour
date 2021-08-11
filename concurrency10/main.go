// Exercise: Web Crawler
// In this exercise you'll use Go's concurrency
// features to parallelize a web crawler.
//
// Modify the Crawl function to fetch URLs in
// parallel without fetching the same URL twice.
//
// Hint: you can keep a cache of the URLs that
// have been fetched on a map, but maps alone
// are not safe for concurrent use!

package main

import (
	"fmt"
	"sync"
)

type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string) (body string, urls []string, err error)
}

type FetchedUrls struct {
	mu sync.Mutex
	v  map[string]bool
}

func (f *FetchedUrls) fetched(url string) {
	f.mu.Lock()
	f.v[url] = true
	f.mu.Unlock()
}

func (f *FetchedUrls) isFetched(url string) bool {
	f.mu.Lock()
	defer f.mu.Unlock()
	_, isFetched := f.v[url]
	return isFetched
}

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(
	url string,
	depth int,
	fetcher Fetcher,
	fUrls *FetchedUrls,
	ch chan string,
) {
	defer close(ch)

	if depth <= 0 {
		return
	}

	isFetched := fUrls.isFetched(url)
	if isFetched {
		return
	}

	body, urls, err := fetcher.Fetch(url)
	fUrls.fetched(url)
	if err != nil {
		ch <- err.Error()
		return
	}

	ch <- fmt.Sprintf("found: %s %q", url, body)

	result := make([]chan string, len(urls))
	for i, u := range urls {
		result[i] = make(chan string)
		go Crawl(u, depth-1, fetcher, fUrls, result[i])
	}

	for i := range result {
		for s := range result[i] {
			ch <- s
		}
	}
}

func main() {
	fUrls := FetchedUrls{v: make(map[string]bool)}
	ch := make(chan string)

	go Crawl("https://golang.org/", 4, fetcher, &fUrls, ch)

	for resp := range ch {
		fmt.Println(resp)
	}
}

// fakeFetcher is Fetcher that returns canned results.
type fakeFetcher map[string]*fakeResult

type fakeResult struct {
	body string
	urls []string
}

func (f fakeFetcher) Fetch(url string) (string, []string, error) {
	if res, ok := f[url]; ok {
		return res.body, res.urls, nil
	}
	return "", nil, fmt.Errorf("not found: %s", url)
}

// fetcher is a populated fakeFetcher.
var fetcher = fakeFetcher{
	"https://golang.org/": &fakeResult{
		"The Go Programming Language",
		[]string{
			"https://golang.org/pkg/",
			"https://golang.org/cmd/",
		},
	},
	"https://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"https://golang.org/",
			"https://golang.org/cmd/",
			"https://golang.org/pkg/fmt/",
			"https://golang.org/pkg/os/",
		},
	},
	"https://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
	"https://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
}
