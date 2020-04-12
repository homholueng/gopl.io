package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

func ElementByID(doc *html.Node, id string) *html.Node {
	preFinder := func(n *html.Node) bool {
		for _, attr := range n.Attr {
			if attr.Key == "id" && attr.Val == id {
				return true
			}
		}
		return false
	}
	return forEachNode(doc, preFinder, nil)
}

func forEachNode(n *html.Node, pre, post func(n *html.Node) bool) *html.Node {
	if pre != nil {
		if pre(n) {
			return n
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		node := forEachNode(c, pre, post)
		if node != nil {
			return node
		}
	}
	if post != nil {
		if post(n) {
			return n
		}
	}
	return nil
}

const usage = `Usage: findnode NODE_ID`

func usageDie() {
	fmt.Println(usage)
	os.Exit(1)
}

func main() {
	if len(os.Args) != 2 {
		usageDie()
	}
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "outline: %v\n", err)
		os.Exit(1)
	}
	node := ElementByID(doc, os.Args[1])
	fmt.Printf("Found: %v\n", node)
}
