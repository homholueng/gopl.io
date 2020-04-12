package main

import (
	"fmt"
	"os"
	"strings"

	"golang.org/x/net/html"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "outline: %v\n", err)
		os.Exit(1)
	}
	forEachNode(doc, startElement, endElement)
}

func forEachNode(n *html.Node, pre, post func(n *html.Node)) {
	if pre != nil {
		pre(n)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}
	if post != nil {
		post(n)
	}
}

var depth int

func startElement(n *html.Node) {
	if n.Type == html.ElementNode {
		attrStrBuilder := strings.Builder{}
		for _, attr := range n.Attr {
			attrStrBuilder.Write([]byte(fmt.Sprintf(" %s='%s'", attr.Key, attr.Val)))
		}
		attrs := attrStrBuilder.String()
		if n.FirstChild != nil {
			fmt.Printf("%*s<%s%s>\n", depth*2, "", n.Data, attrs)
		} else {
			fmt.Printf("%*s<%s%s/>\n", depth*2, "", n.Data, attrs)
		}
		depth++
	} else if n.Type == html.TextNode {
		fmt.Printf("%*s%s\n", depth*2, "", n.Data)
	} else if n.Type == html.CommentNode {
		fmt.Printf("<!--%s-->\n", n.Data)
	}
}

func endElement(n *html.Node) {
	if n.Type == html.ElementNode {
		depth--
		if n.FirstChild != nil {
			fmt.Printf("%*s</%s>\n", depth*2, "", n.Data)
		}
	}
}
