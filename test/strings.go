package main

import "fmt"

type _string struct {
	elements *byte
	len      int
}

func main() {
	const World = "world"
	var hello = "hello"

	var helloWorld = hello + " " + World
	helloWorld += "!"
	fmt.Println(helloWorld)

	fmt.Println(hello == "hello")
	fmt.Println(hello > helloWorld)
}
