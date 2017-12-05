package main

import (
	"fmt"
	"sync"
)

func main() {
	input := gen()

	// FANOUT, multiple workers working on same channel
	c0 := factorial(input)
	c1 := factorial(input)
	c2 := factorial(input)
	c3 := factorial(input)
	c4 := factorial(input)
	c5 := factorial(input)
	c6 := factorial(input)
	c7 := factorial(input)
	c8 := factorial(input)
	c9 := factorial(input)

	for c := range merge(c0, c1, c2, c3, c4, c5, c6, c7, c8, c9) {
		fmt.Println(c)
	}
}

func gen() <-chan int {
	out := make(chan int)
	go func() {
		for i := 0; i < 100000; i++ {
			for index := 3; index < 13; index++ {
				out <- index
			}
		}
		close(out)
	}()
	return out
}

func factorial(input <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range input {
			out <- fact(n)
		}
		close(out)
	}()
	return out
}

func fact(n int) int {
	total := 1
	for i := n; i > 0; i-- {
		total *= i
	}
	return total
}

func merge(cs ...<-chan int) <-chan int {
	out := make(chan int)
	var wg sync.WaitGroup

	output := func(c <-chan int) {
		for n := range c {
			out <- n
		}
		wg.Done()
	}
	for _, c := range cs {
		go output(c)
	}

	wg.Add(len(cs))
	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}
