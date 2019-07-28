package main

import "fmt"

func main() {
	data := make([]int, 4)

	loopData := func(hanleData chan<- int) {
		defer close(hanleData)
		for i := range data {
			hanleData <- data[i]
		}
	}

	handleData := make(chan int)
	go loopData(handleData)

	for num := range handleData {
		fmt.Println(num)
	}

}
