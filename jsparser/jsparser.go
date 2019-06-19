package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/tamerh/jsparser"
)

func main() {
	f, _ := os.Open("input.json")
	br := bufio.NewReaderSize(f,65536)
	parser := jsparser.NewJSONParser(br, "books")

	for json:= range parser.Stream() {
		fmt.Println(json.ObjectVals["title"].StringVal)
		fmt.Println(json.ObjectVals["price"].StringVal)
		fmt.Println(json.ObjectVals["comments"].ArrayVals[0].ObjectVals["rating"].StringVal)
	}
}
