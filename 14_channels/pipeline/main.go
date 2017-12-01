package main

import (
	"fmt"
)

// Pipeline Using Channels
func main() {
	for n := range square(generate(2, 3)) {
		fmt.Println(n)
	}

}

func generate(nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		for _, n := range nums { // without blank identifier it doesn' works
			out <- n
		}
		close(out)
	}()
	return out
}

func square(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n * n
		}
		close(out)
	}()
	fmt.Println(<-out)
	return out
}
