package main

import "fmt"

// Closure with anonymus function Variable Type
func main() {
	var increment = wrapper()
	fmt.Println(increment())
	fmt.Println(increment())
}

func wrapper() func() int {
	x := 0
	return func() int {
		x++
		return x
	}
}
