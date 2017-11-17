package main

import (
	"fmt"
)

func main() {
	fmt.Println(foo())
	fmt.Println(foo(1, 2))
	fmt.Println(foo(1, 2, 3))
	nSlice := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(foo(nSlice...))
}

func foo(number ...int) []int {
	return number
}
