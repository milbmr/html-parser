package parser

import (
	"log"
	"strings"

	"golang.org/x/net/html"
)

type Link struct {
	Link string
	Data string
}

func Parser(document string) []Link {
	doc, err := html.Parse(strings.NewReader(document))
	if err != nil {
		log.Fatal(err)
	}

	links := linkParser(doc)
	urls := linkBuilder(links)

	return urls
}

func linkBuilder(links []*html.Node) []Link {
	var ret []Link
	var value Link
	var text string

	for _, link := range links {
		for _, attr := range link.Attr {
			if attr.Key == "href" {
				text = textParser(link)
				value.Link = attr.Val
				value.Data = text
			}
		}
		ret = append(ret, value)
	}

	return ret
}

func linkParser(n *html.Node) []*html.Node {
	var ret []*html.Node

	if n.Type == html.ElementNode && n.Data == "a" {
		return append(ret, n)
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		ret = append(ret, linkParser(c)...)
	}

	return ret
}

func textParser(link *html.Node) string {
	var ret string

	if link.Type == html.TextNode {
		return link.Data
	}

	if link.Type != html.ElementNode {
		return ""
	}

	for c := link.FirstChild; c != nil; c = c.NextSibling {
		ret += textParser(c)
	}

	return strings.Join(strings.Fields(ret), " ")
}
