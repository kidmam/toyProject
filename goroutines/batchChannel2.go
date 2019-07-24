package main

import (
	"context"
	"fmt"
	"time"
)

func BatchStringsCtx(ctx context.Context, values <-chan string, maxItems int, maxTimeout time.Duration) chan []string {
	batches := make(chan []string)

	go func() {
		defer close(batches)

		for keepGoing := true; keepGoing; {
			var batch []string
			expire := time.After(maxTimeout)
			for {
				select {
				case <-ctx.Done():
					keepGoing = false
					goto done

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
		// The context was cancelled around 300ms,
		// before this sleep finished.
		time.Sleep(500 * time.Millisecond)

		strings <- "doing?" // never read

		close(strings)
	}()

	ctx, cancel := context.WithTimeout(context.Background(), 300*time.Millisecond)
	defer cancel()

	start := time.Now()
	batches := BatchStringsCtx(ctx, strings, 2, 10*time.Millisecond)
	for batch := range batches {
		fmt.Println(time.Now().Sub(start), batch)
	}
}
