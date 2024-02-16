package main

import (
	"fmt"

	"github.com/milbmr/html-parser"
)

var ex = `
  <html>
<body>
  <a href="/dog-cat">dog cat <!-- commented text SHOULD NOT be included! --></a>
</body>
</html>
  `

func main() {
	links := parser.Parser(ex)
	for _, link := range links {
		fmt.Printf("%+v", link)
	}
}
