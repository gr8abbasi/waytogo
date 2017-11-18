package main

import (
	"fmt"
)

func main() {
	x := 8
	funcExpression := func(x int) (float64, bool) {
		return float64(x) / 2, x%2 == 0
	}
	fmt.Println(funcExpression(x))
}
