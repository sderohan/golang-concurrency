package main

import (
	"fmt"
	"sync"
	"time"
)

/*
	Calcualting the time spent for program execution with and without goroutine
	NOTE : The variables inside function a, b, and temp will overflow. I have used them only for testing.
*/

func main() {
	max := 1000000

	start := time.Now()

	for i := 1; i < max; i++ {
		func(i int) {
			a, b := 0, 1
			for j := 0; j < i; j++ {
				temp := a
				a = b
				b = a + temp
			}
		}(i)
	}
	fmt.Println("Without goroutine time spent ", time.Now().Sub(start).Seconds())

	start = time.Now()

	// using goroutine
	var wg sync.WaitGroup
	for i := 1; i < max; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			a, b := 0, 1
			for j := 0; j < i; j++ {
				temp := a
				a = b
				b = a + temp
			}
		}(i)
	}
	wg.Wait()
	fmt.Println("Using goroutine time spent ", time.Now().Sub(start).Seconds())
}
