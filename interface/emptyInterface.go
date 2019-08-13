package main

func main() {
	var i int8 = 1
	read(i)
}

func read(i interface{}) {
	n := i.(int8)
	println(n)
}
