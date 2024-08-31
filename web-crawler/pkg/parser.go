package webcrawler

import (
	"log"
	"strings"

	"golang.org/x/net/html"
)

func Parse(body string) []string {
	var urls []string

	doc, err := html.Parse(strings.NewReader(body))
	if err != nil {
		log.Fatal(err)
	}

	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, a := range n.Attr {
				if a.Key == "href" {
					urls = append(urls, a.Val)
					break
				}
			}
		}

		for c := n.FirstChild; c != nil; c = n.NextSibling {
			f(c)
		}
	}

	f(doc)

	return urls
}
