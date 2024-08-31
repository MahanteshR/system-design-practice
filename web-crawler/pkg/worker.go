package webcrawler

import (
	"fmt"
	"log"
	"strings"
	"sync"
	"time"
)

func Worker(frontier *Frontier, wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		url := frontier.GetURL()
		fmt.Println("Crawling:", url)

		body, err := Fetch(url)
		fmt.Print(body)
		if err != nil {
			log.Println("Error fetching:", err)
			continue
		}

		links := Parse(body)
		for _, link := range links {
			if strings.HasPrefix(link, "/") {
				link = url + link
			}

			frontier.AddURL(link)
		}

		time.Sleep(500 * time.Millisecond)
	}
}

func StartWorkerPool(n int, frontier *Frontier) {
	var wg sync.WaitGroup
	for i := 0; i < n; i++ {
		wg.Add(1)
		go Worker(frontier, &wg)
	}
	wg.Wait()
}
