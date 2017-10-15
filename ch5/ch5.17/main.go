package main

import (
	"fmt"
	"net/http"
	"os"

	"golang.org/x/net/html"
)

func main() {

	url := "https://golang.org"
	for _, link := range os.Args[1:] {
		url = link
		break
	}

	resp, err := http.Get(url)
	if err != nil {
		fmt.Errorf("HTTP request error - %s", err)
	}

	defer resp.Body.Close()

	doc, err := html.Parse(resp.Body)
	if err != nil {
		fmt.Errorf("Parsing %s as HTML: %v", url, err)
	}

	for _, element := range elementsByTagName(doc, "option", "span", "div") {
		fmt.Printf("%v\n", element)
	}
}

func elementsByTagName(doc *html.Node, names ...string) []*html.Node {

	nameMap := make(map[string]bool)
	for _, name := range names {
		nameMap[name] = true
	}

	elements := visit(nil, doc, nameMap)

	return elements
}

func visit(elements []*html.Node, n *html.Node, nameMap map[string]bool) []*html.Node {
	if n == nil {
		return elements
	}

	if n.Type == html.ElementNode {
		if _, ok := nameMap[n.Data]; ok {
			elements = append(elements, n)
		}
	}

	elements = visit(elements, n.FirstChild, nameMap)
	return visit(elements, n.NextSibling, nameMap)
}
