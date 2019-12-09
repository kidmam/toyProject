package main

import "fmt"

func main() {
	var my_int_Channel chan int
	my_string_Channel := make(chan string)

	fmt.Printf("The type of my_int_Channel is '%T' \n", my_int_Channel)
	fmt.Printf("The type of my_string_Channels is '%T' \n", my_string_Channel)
}
