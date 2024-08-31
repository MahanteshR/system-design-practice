package webcrawler

import (
	"fmt"
	"sync"
)

type Frontier struct {
	queue   chan string
	visited map[string]bool
	mu      sync.Mutex
}

func NewFrontier(limit int) *Frontier {
	return &Frontier{
		queue:   make(chan string, limit),
		visited: make(map[string]bool),
	}
}

func (f *Frontier) AddURL(url string) {
	f.mu.Lock()
	defer f.mu.Unlock()

	if !f.visited[url] {
		f.visited[url] = true
		f.queue <- url
	}

	fmt.Println(f.queue)
}

func (f *Frontier) GetURL() string {
	return <-f.queue
}
