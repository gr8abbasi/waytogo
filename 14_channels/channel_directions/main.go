package main

import (
	"fmt"
)

func main() {
	c := incrementor()
	for n := range puller(c) {
		fmt.Println(n)
	}
}

// chan bidirectional channel
// <-chan receive only channel
// chan<- send only channel
func incrementor() <-chan int { // receive only channel
	out := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			out <- i
		}
		close(out)
	}()
	return out
}

func puller(c <-chan int) <-chan int { // receive only channel
	out := make(chan int)
	go func() {
		var sum int
		for n := range c {
			sum += n
		}
		out <- sum
		close(out)
	}()
	return out
}
