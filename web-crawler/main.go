package main

import webcrawler "web-crawler/pkg"

func main() {
	frontier := webcrawler.NewFrontier(100)
	frontier.AddURL("https://web-scraping.dev/")

	webcrawler.StartWorkerPool(4, frontier)
}
