package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlink: %v\n", err)
		os.Exit(1)
	}
	for name, count := range visit(make(map[string]int, 1), doc) {
		fmt.Println(name, count)
	}
}

func visit(names map[string]int, n *html.Node) map[string]int {

	if n == nil {
		return names
	}
	names = visit(names, n.NextSibling)

	if n.Type == html.ElementNode {
		if _, ok := names[n.Data]; !ok {
			names[n.Data] = 0
		}
		names[n.Data]++
	}

	names = visit(names, n.FirstChild)
	return names
}
