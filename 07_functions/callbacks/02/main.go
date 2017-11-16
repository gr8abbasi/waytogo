package main

import (
	"fmt"
)

func main() {
	xs := filter([]int{1, 2, 3, 4, 5, 6}, func(n int) bool {
		return n > 3
	})
	fmt.Println(xs)
}

func filter(n []int, callback func(n int) bool) []int {
	var xs []int
	for _, v := range n {
		if callback(v) {
			xs = append(xs, v)
		}
	}
	return xs
}
