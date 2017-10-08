package main

import (
	"fmt"
	"net/http"
	"os"

	"golang.org/x/net/html"
)

func visit(links map[string]int, n *html.Node) map[string]int {
	if n == nil {
		return links
	}

	if n.Type == html.ElementNode {
		links[n.Data] = links[n.Data] + 1
	}

	visit(links, n.FirstChild)
	return visit(links, n.NextSibling)
}

func main() {
	for _, url := range os.Args[1:] {
		links, err := findLinks(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "findlinks2: %v\n", err)
			continue
		}
		for element, count := range links {
			fmt.Printf("%s element - %d times\n", element, count)
		}
	}
}

func findLinks(url string) (map[string]int, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	doc, err := html.Parse(resp.Body)

	if err != nil {
		return nil, fmt.Errorf("parsing %s as HTML: %v", url, err)
	}

	links := make(map[string]int)
	return visit(links, doc), nil
}
