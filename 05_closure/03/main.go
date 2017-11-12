package main

import "fmt"

// Closure with anonymus function Variable Type
func main() {
	increment := wrapper()
	fmt.Println(increment())
	fmt.Println(increment())
	// Returns new function, hence new scope with new copy of x!
	increment = wrapper()
	fmt.Println(increment()) // 1
	fmt.Println(increment()) // 2
}

func wrapper() func() int {
	x := 0
	return func() int {
		x++
		return x
	}
}
