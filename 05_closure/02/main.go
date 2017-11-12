package main

import "fmt"

// Closure with anonymus functions
func main() {
	x := 0
	var increment = func() int {
		x++
		return x
	}
	fmt.Println(increment())
	fmt.Println(increment())
}
