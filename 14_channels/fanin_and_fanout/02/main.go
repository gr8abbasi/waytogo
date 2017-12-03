package main

import (
	"fmt"
	"sync"
)

func main() {
	gen := gen(2, 3)

	c1 := sq(gen)
	c2 := sq(gen)

	for n := range merge(c1, c2) {
		fmt.Println(n)
	}
}

func gen(nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
	}()
	return out
}

func sq(nums <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range nums {
			out <- n * n
		}
		close(out)
	}()
	return out
}

func merge(cs ...<-chan int) <-chan int {
	out := make(chan int)
	var wg sync.WaitGroup
	wg.Add(len(cs))

	for _, c := range cs {
		go func(ch <-chan int) { // closure receiving param c
			for n := range ch {
				out <- n
				wg.Done()
			}
		}(c)
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}
