package main

import (
	"fmt"
)

func main() {
	x := 8
	funcExpression := func(x int) (float32, bool) {
		var even bool
		if x%2 == 0 {
			even = true
		}
		return float32(x) / 2, even
	}
	fmt.Println(funcExpression(x))
}
