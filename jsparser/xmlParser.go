package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/tamerh/xml-stream-parser"
)

func main() {
	f, _ := os.Open("input.xml")
	br := bufio.NewReaderSize(f, 65536)
	parser := xmlparser.NewXMLParser(br, "book")

	for xml := range parser.Stream() {
		fmt.Println(xml)
		//fmt.Println(xml.Childs["title"][0].InnerText)
		//fmt.Println(xml.Childs["comments"][0].Childs["userComment"][0].Attrs["rating"])
		//fmt.Println(xml.Childs["comments"][0].Childs["userComment"][0].InnerText)
	}
}
