package main

import (
	// "fmt"
	"log"
	"os"
	"runtime"
	"strconv"
	"sync"
	"time"
)

var fileLogger *log.Logger

// https://jusths.tistory.com/128
func main() {
	runtime.GOMAXPROCS(1)
	log.Println("Concurrently logging to File...")

	var wait sync.WaitGroup
	rNum := 100
	wait.Add(rNum)

	fn := "concurrent_log.log"
	f, err := os.OpenFile(fn, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	defer f.Close()
	if err != nil {
		log.Fatal(err)
	}

	fileLogger = log.New(f, "", 0)

	for i := 0; i < rNum; i++ {
		go func(i int, f *os.File) {
			defer wait.Done()
			for j := 0; j < 1000; j++ {
				tLog := time.Now().Format("060102_150405")
				tLog = strconv.Itoa(i) + ": " + tLog //+ "\n"
				log.Println(tLog)
				fileLogger.Output(1, tLog)
				// f.WriteString(tLog)
			}
		}(i, f)
	}
	wait.Wait()
}
