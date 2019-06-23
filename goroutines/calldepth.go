package main

import (
	// "fmt"
	"log"
	"os"
	"runtime"
)

var fileLogger *log.Logger

func cd1(f *os.File) {
	cd2(f)
}

func cd2(f *os.File) {
	cd3(f)
}

func cd3(f *os.File) {
	fileLogger.Output(1, "I am cd3")
	fileLogger.Output(2, "I am cd3")
	fileLogger.Output(3, "I am cd3")
}

func main() {
	runtime.GOMAXPROCS(1)
	log.Println("Call depth test...")

	fn := "calldepth.log"
	f, err := os.OpenFile(fn, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	defer f.Close()
	if err != nil {
		log.Fatal(err)
	}

	fileLogger = log.New(f, "", log.Llongfile)

	cd1(f)
}
