package main

import (
	"fmt"
	"strconv"
	"sync"
)

// https://medium.com/@codyoss/reflecting-on-worker-pools-in-go-7f91f05a5f01
func main() {
	// Section 1
	parallelism := 5
	inCh1 := make(chan int, parallelism)
	inCh2 := make(chan int, parallelism)
	inCh3 := make(chan int, parallelism)
	outCh := make(chan string, parallelism)
	var wg1 sync.WaitGroup
	var wg2 sync.WaitGroup
	var wg3 sync.WaitGroup

	// Section 2
	for i := 0; i < parallelism; i++ {
		wg1.Add(1)
		wg2.Add(1)
		wg3.Add(1)

		go func() {
			for v := range inCh1 {
				v := add1(v)
				inCh2 <- v
			}
			wg1.Done()
		}()

		go func() {
			for v := range inCh2 {
				v := add2(v)
				inCh3 <- v
			}
			wg2.Done()
		}()

		go func() {
			for v := range inCh3 {
				v := stringify(v)
				outCh <- v
			}
			wg3.Done()
		}()
	}

	// Section 3
	for i := 0; i < 10; i++ {
		inCh1 <- i
		output := <-outCh
		fmt.Println(output)
	}

	// Section 4
	close(inCh1)
	wg1.Wait()
	close(inCh2)
	wg2.Wait()
	close(inCh3)
	wg3.Wait()
	close(outCh)
}

func add1(i int) int {
	return i + 1
}

func add2(i int) int {
	return i + 2
}

func stringify(i int) string {
	return strconv.Itoa(i)
}
