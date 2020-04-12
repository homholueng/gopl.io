package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

func ElementByTagName(doc *html.Node, names ...string) []*html.Node {
	nameMap := make(map[string]bool, len(names))

	for _, name := range names {
		nameMap[name] = true
	}

	found := make([]*html.Node, 0)

	preFinder := func(n *html.Node) {
		if n.Type == html.ElementNode && nameMap[n.Data] {
			found = append(found, n)
		}
	}
	forEachNode(doc, preFinder, nil)

	return found
}

func forEachNode(n *html.Node, pre, post func(n *html.Node)) *html.Node {
	if pre != nil {
		pre(n)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		node := forEachNode(c, pre, post)
		if node != nil {
			return node
		}
	}
	if post != nil {
		post(n)
	}
	return nil
}

const usage = `Usage: findnode TAG_NAME [TAG_NAME...]`

func usageDie() {
	fmt.Println(usage)
	os.Exit(1)
}

func main() {
	if len(os.Args) < 2 {
		usageDie()
	}
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "outline: %v\n", err)
		os.Exit(1)
	}
	nodes := ElementByTagName(doc, os.Args[1:]...)
	fmt.Println("Found:")
	for _, node := range nodes {
		fmt.Println(node.Data)
	}
}
