package main

import "fmt"

var x = 0

// Closure without anonymus functions
// Closure refers to how scopes are enclosed
func main() {
	fmt.Println(increment())
	fmt.Println(increment())
}

func increment() int {
	x++
	return x
}
