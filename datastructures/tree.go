package main

import "fmt"

type Node struct {
	tag      string
	text     string
	id       string
	children []*Node
}

func main() {
	p := Node{
		tag:  "p",
		text: "This is a simple HTML document.",
		id:   "foo",
	}

	h1 := Node{
		tag:  "h1",
		text: "Hello, World!",
	}

	html := Node{
		tag:      "html",
		children: []*Node{&p, &h1},
	}

	fmt.Println("html ==> ", html)
}
