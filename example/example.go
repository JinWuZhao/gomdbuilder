package main

import (
	"io/ioutil"

	. "github.com/JinWuZhao/gomdbuilder"
)

func main() {
	doc := Doc(
		BR,
		"Heading",
		HLE,
		H2("Sub-heading"),
		BR,
		P(`Paragraphs are separated
by a blank line.`),
		BR,
		P(`Two spaces at the end of a line  
produces a line break.`),
		BR,
		P("Text attributes ", I("italic"), ","),
		P(B("bold"), ",", M("monospace"), ","),
		P("Horizontal rule:"),
		HR,
		P("Bullet list:"),
		BLi("apples",
			"oranges",
			"pears"),
		BR,
		P("Numbered list:"),
		NLi("wash",
			"rinse",
			"repeat"),
		BR,
		P("A ", Ln("link", "http://example.com"), "."),
		P(Img("Image", "https://avatars2.githubusercontent.com/u/4271558?s=460&v=4")),
		BR,
		Q1("Markdown uses email-style > characters for blockquoting."),
		Q2("Markdown uses email-style > characters for blockquoting."),
		BR,
		Code("go", `
package main

import "fmt"

func main() {
	fmt.Println("Hello World!")
}
`))

	ioutil.WriteFile("example.md", []byte(doc), 0666)
}
