package main

import (
	"fmt"
	"sync"
)

type Fetcher interface {
	Fetch(url string) (body string, urls []string, err error)
}

type SafeMap struct {
	visited map[string]bool
	mu      sync.Mutex
}

func (sm *SafeMap) IsVisited(url string) bool {
	sm.mu.Lock()
	defer sm.mu.Unlock()
	return sm.visited[url]
}

func (sm *SafeMap) MarkVisited(url string) {
	sm.mu.Lock()
	defer sm.mu.Unlock()
	sm.visited[url] = true
}

func Crawl(url string, depth int, fetcher Fetcher, sm *SafeMap, wg *sync.WaitGroup) {
	defer wg.Done()

	if depth <= 0 || sm.IsVisited(url) {
		return
	}

	sm.MarkVisited(url)

	body, urls, err := fetcher.Fetch(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("found: %s %q\n", url, body)
	for _, u := range urls {
		wg.Add(1)
		go Crawl(u, depth-1, fetcher, sm, wg)
	}
}

func main() {
	var wg sync.WaitGroup // To manage goroutines
	sm := &SafeMap{visited: make(map[string]bool)}

	wg.Add(1)
	go Crawl("https://golang.org/", 4, fetcher, sm, &wg)

	wg.Wait()
}
