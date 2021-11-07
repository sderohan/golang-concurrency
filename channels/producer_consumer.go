package main

import (
	"fmt"
	"math/rand"
	"time"
)

/*
	Producer Consumer
	Producer writes the data onto the channel, and consumer consumes it
*/

func main() {
	ch := make(chan int)

	// Producer
	go func(ch chan int) {
		i := 0
		for {
			ch <- i
			i++
			time.Sleep(time.Second * (time.Duration(rand.Int()) % 10))
		}
	}(ch)

	// Consumer
	for data := range ch {
		fmt.Println("Consumer -->> ", data)
	}
}
