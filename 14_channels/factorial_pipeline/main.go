package main

import (
	"fmt"
)

func main() {
	gen := generator(13)
	fac := factorial(gen)
	for n := range fac {
		fmt.Println(n)
	}
}

func generator(n int) <-chan int {
	numbers := make(chan int)
	go func() {
		for index := 1; index <= n; index++ {
			numbers <- index
		}
		close(numbers)
	}()
	return numbers
}

func factorial(number <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range number {
			total := 1
			for index := n; index > 0; index-- {
				total *= index
			}
			out <- total
		}
		close(out)
	}()
	return out
}
