package parser

import (
	"fmt"
	"log"
	"strings"

	"golang.org/x/net/html"
)

var htm = `
  <a href="/animal">
    <div class="container">
      
    </div>
  <a href="/dog">
        <span>Something in a span</span>
        Text not in a span
        <b>Bold text!</b>
      </a>
  </a>`

func Parser() {
	doc, err := html.Parse(strings.NewReader(htm))
	if err != nil {
		log.Fatal(err)
	}

	htmlItirator(doc)
}

func htmlItirator(n *html.Node) {
	if n == nil {
		return
	}

	// for _, a := range n.Attr {
	// 	fmt.Printf("%s: %s\n", a.Key, a.Val)
	// }

	fmt.Println(" ", n.Data)

	// c := n
	//
	// if n.FirstChild != nil {
	// 	htmlItirator(c.FirstChild)
	// }
	//
	// if n.NextSibling != nil {
	// 	htmlItirator(c.NextSibling)
	// }

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		htmlItirator(c)
	}
}
