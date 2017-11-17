package main

import (
	"fmt"
)

func main() {
	x := 5
	fmt.Println(halfEven(x))
}

func halfEven(x int) (float32, bool) {
	var even bool
	if x%2 == 0 {
		even = true
	}
	return float32(x) / 2, even
}
