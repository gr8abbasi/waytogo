package main

import (
	"fmt"
)

func main() {
	x := 5
	fmt.Println(halfEven(x))
}

func halfEven(x int) (float64, bool) {
	return float64(x) / 2, x%2 == 0
}
