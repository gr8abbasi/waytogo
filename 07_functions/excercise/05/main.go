package main

import (
	"fmt"
)

func main() {
	fmt.Println(fibonnaciNumbers())
}

func fibonnaciNumbers() int {
	previous := 1
	var total int
	for i := 0; i < 4000000; i++ {
		previous += i
		if getEvenNumbers(previous) {
			total = sum(total, previous)
		}
	}
	return total
}

func getEvenNumbers(x int) bool {
	return x%2 == 0
}

func sum(sum int, x int) int {
	return sum + x
}
