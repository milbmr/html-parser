package main

import (
	"fmt"

	"github.com/milbmr/html-parser"
)

func main() {
	links := parser.Parser()
	for _, link := range links {
		fmt.Printf("%+v", link)
	}
}
