package main

import (
	"fmt"
)

func main() {
	c := make(chan int) // unbuffered channel which doesn't specify length

	go func() {
		for i := 0; i < 10; i++ {
			c <- i // adding to channel and passing to next routine to offload it
		}
		close(c) // channel is closed and you can write anything now. only can read from it
	}()

	for n := range c {
		fmt.Println(n)
	}
}
