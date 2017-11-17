package main

import (
	"fmt"
)

func main() {
	numbers := []int{3, 7, 4, 1, 8, 9, 16, 5}
	greater := findGreatest(numbers...)
	fmt.Println(greater)
}

func findGreatest(number ...int) int {
	if len(number) <= 0 {
		return 0
	}

	greater := number[0]
	for _, v := range number {
		if v > greater {
			greater = v
		}
	}
	return greater
}
