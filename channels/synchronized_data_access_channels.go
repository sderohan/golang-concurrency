/*
	There are n goroutines using the data variable
	achieve the synchronized data access using the channels
*/
package main

import (
	"fmt"
	"sync"
)

func main() {
	n, data := 100, 0
	ch := make(chan int, 1) // need to clear why this is not working on the unbuffered channel
	var wg sync.WaitGroup

	//start n goroutine
	for i := 0; i < n; i++ {
		wg.Add(1)
		go func(ch chan int) {
			defer wg.Done()
			// read the data, increment it and write back on channel
			data := <-ch
			data++
			ch <- data
		}(ch)
	}

	// pass the data on channel
	ch <- data

	// wait for all goroutine to finish the execution
	wg.Wait()

	// get the data from the channel
	data = <-ch

	// print the data value
	fmt.Println(data)
}
