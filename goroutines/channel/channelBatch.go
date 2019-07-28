package main

// https://medium.com/@elliotchance/batch-a-channel-by-size-or-time-in-go-92fa3098f65
import (
	"fmt"
	"time"
)

func BatchStrings(values <-chan string, maxItems int, maxTimeout time.Duration) chan []string {
	batches := make(chan []string)

	go func() {
		defer close(batches)

		for keepGoing := true; keepGoing; {
			var batch []string
			expire := time.After(maxTimeout)
			for {
				select {
				case value, ok := <-values:
					if !ok {
						keepGoing = false
						goto done
					}

					batch = append(batch, value)
					if len(batch) == maxItems {
						goto done
					}

				case <-expire:
					goto done
				}
			}

		done:
			if len(batch) > 0 {
				batches <- batch
			}
		}
	}()

	return batches
}

func main() {
	strings := make(chan string)

	go func() {
		strings <- "hello"
		strings <- "there" // hit limit of 2

		strings <- "how"
		time.Sleep(15 * time.Millisecond) // hit timeout

		strings <- "are"
		strings <- "you" // hit limit of 2

		// A really long time without any new values.
		time.Sleep(500 * time.Millisecond)

		strings <- "doing?" // the last incomplete batch

		close(strings)
	}()

	start := time.Now()
	batches := BatchStrings(strings, 2, 10*time.Millisecond)
	for batch := range batches {
		fmt.Println(time.Now().Sub(start), batch)
	}
}
