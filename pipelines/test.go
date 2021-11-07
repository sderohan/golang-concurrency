package test

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string)
	go func() {
		ch <- "one"
	}()
	// time.Sleep(time.Second)
	go func() {
		time.Sleep(3 * time.Second)
		ch <- "from goroutine two"
	}()
	for i := 0; i < 2; i++ {
		select {
		case data := <-ch:
			fmt.Println(data)
		case <-time.After(time.Second * 3):
			fmt.Println("Quit")
		}
	}
	// main()
	// select {}
}
