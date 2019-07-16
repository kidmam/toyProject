package main

import (
	"fmt"
)

func main() {
	stringStream := make(chan string)
	//var stringStream chan<- string
	//stringStream = make(chan<- string)

	go func() {
		/*if 0 != 1 {
			return
		}*/ //deadlock

		stringStream <- "Hello channels!"
	}()

	salutaion, ok := <-stringStream
	fmt.Printf("(%v): %v", ok, salutaion)
}
