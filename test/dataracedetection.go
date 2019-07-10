package main

import "sync"

/*func main() {
	m := make(map[string]int, 1)
	m[`foo`] = 1

	var wg sync.WaitGroup

	wg.Add(2)
	go func() {
		defer wg.Done()
		for i := 0; i < 1000; i++  {
			m[`foo`]++
		}
	}()
	go func() {
		defer wg.Done()
		for i := 0; i < 1000; i++  {
			m[`foo`]++
		}
	}()
	wg.Wait()
}*/

var foo = 0

func main() {
	var wg sync.WaitGroup

	wg.Add(2)
	go func() {
		defer wg.Done()
		for i := 0; i < 50000; i++ {
			foo++
		}
	}()
	go func() {
		defer wg.Done()
		for i := 0; i < 50000; i++ {
			foo++
		}
	}()
	wg.Wait()
	println(foo)
}
