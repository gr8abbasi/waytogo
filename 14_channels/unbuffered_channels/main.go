package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int) // unbuffered channel which doesn't specify length

	go func() {
		for i := 0; i < 10; i++ {
			c <- i // adding to channel and passing to next routine to offload it
		}
	}()

	go func() {
		for {
			fmt.Println(<-c) // offload the value from the channel
		}
	}()

	time.Sleep(time.Second)
}
