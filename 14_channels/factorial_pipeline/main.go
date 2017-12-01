package main

import (
	"fmt"
)

func main() {
	gen := generator(100)
	fac := factorial(gen)
	for n := range fac {
		fmt.Println(n)
	}
}

func generator(n int) <-chan int {
	numbers := make(chan int)
	for index := 1; index <= 100; index++ {
		go func() {
			numbers <- index
		}()
	}
	close(numbers)
	return numbers
}

func factorial(number <-chan int) <-chan int {
	out := make(chan int)
	for n := range number {
		go func() {
			total := 1
			for index := n; index > 0; index-- {
				total *= index
			}
			out <- total
		}()
	}
	close(out)
	return out
}
