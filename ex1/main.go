package main

import (
	"fmt"

	"github.com/milbmr/html-parser"
)

var ex = `
<html>
<body>
  <h1>Hello!</h1>
  <a href="/other-page">A link to another page</a>
</body>
</html>
`

func main() {
	links := parser.Parser(ex)
	for _, link := range links {
		fmt.Printf("%+v", link)
	}
}
